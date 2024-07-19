package handler

import (
	"app/app"
	"app/internal/htmx"
	"app/model"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"slices"

	"github.com/a-h/templ"
	"github.com/alexedwards/scs/v2"
	"github.com/gofrs/uuid/v5"
	"github.com/gorilla/mux"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
	"github.com/stripe/stripe-go/v76/webhook"
)

func init() {
	stripe.Key = os.Getenv("STRIPE_API_KEY")
}

func createCheckoutSession(host url.URL, s model.Service, sess *scs.SessionManager) http.HandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) {

		items := []string{"apple", "banana", "carrot"}
		var lineItems []*stripe.CheckoutSessionLineItemParams
		for _, item := range items {
			lineItems = append(lineItems, &stripe.CheckoutSessionLineItemParams{
				// Provide the exact Price ID (for example, pr_1234) of the product you want to sell
				// Price:    stripe.String("price_1Nuy3AIf47Y1plSJCpTmapHb"),
				Quantity: stripe.Int64(1),
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name:        stripe.String(item),
						Description: stripe.String(""),
					},
					UnitAmount: stripe.Int64(500), // $5.00
					Currency:   stripe.String("aud"),
					// Product:     stripe.String("prod_OiOqS8eOthcKR3"),
				},
			})
		}

		// Create an order
		order, err := s.CreateFullOrder(req.Context(), items)
		if err != nil {
			slog.Error("App Handler: create checkout session, create order", "error", err)
			http.Error(wr, "Something when wrong", http.StatusInternalServerError)
			return
		}

		params := &stripe.CheckoutSessionParams{
			LineItems:  lineItems,
			Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
			SuccessURL: stripe.String(host.JoinPath("/orders/" + order.ID.String()).String()),
			CancelURL:  stripe.String(host.JoinPath("/cart/").String()),
			Metadata: map[string]string{
				"order_id": order.ID.String(),
				"app_env":  s.AppEnv,
			},
		}

		s, err := session.New(params)
		if err != nil {
			slog.Error("App Handler: create checkout session", "error", err)
			http.Error(wr, "Something when wrong", http.StatusInternalServerError)
			return
		}

		slog.Debug("Stripe checkout session created", "url", s.URL)

		http.Redirect(wr, req, s.URL, http.StatusSeeOther)
	}
}

func createCheckoutSessionBuyNow(host url.URL, s model.Service) http.HandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) {

		// Create an order
		order, err := s.CreateFullOrder(req.Context(), []string{"apple"})
		if err != nil {
			slog.Error("App Handler: create checkout session, create order", "error", err)
			http.Error(wr, "Something when wrong", http.StatusInternalServerError)
			return
		}

		params := &stripe.CheckoutSessionParams{
			LineItems: []*stripe.CheckoutSessionLineItemParams{
				{
					// Provide the exact Price ID (for example, pr_1234) of the product you want to sell
					// Price:    stripe.String("price_1Nuy3AIf47Y1plSJCpTmapHb"),
					Quantity: stripe.Int64(1),
					PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
						ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
							Name: stripe.String("apple"),
						},
						UnitAmount: stripe.Int64(500), // $5.00
						Currency:   stripe.String("aud"),
						// Product:     stripe.String("prod_OiOqS8eOthcKR3"),
					},
				},
			},
			Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
			SuccessURL: stripe.String(host.JoinPath("/orders/" + order.ID.String()).String()),
			CancelURL:  stripe.String(host.JoinPath("/").String()),
			Metadata:   map[string]string{"order_id": order.ID.String()},
		}

		s, err := session.New(params)
		if err != nil {
			slog.Error("App Handler: create checkout session", "error", err)
			http.Error(wr, "Something when wrong", http.StatusInternalServerError)
			return
		}

		slog.Debug("Stripe checkout session created", "url", s.URL)

		http.Redirect(wr, req, s.URL, http.StatusSeeOther)
	}
}

func checkoutSuccess(s model.Service, sess *scs.SessionManager) http.HandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) {
		varOrderID := mux.Vars(req)["order_id"]
		orderID := uuid.FromStringOrNil(varOrderID)

		order, err := s.GetOrder(req.Context(), orderID)
		if err != nil {
			slog.Error("App Handler: index, list events", "error", err)
			http.Error(wr, "Something when wrong", http.StatusInternalServerError)
			return
		}

		var content templ.Component
		switch order.Status {
		case model.OrderPending:
			content = app.OrderPending(order)
		case model.OrderComplete:
			content = app.OrderComplete(order)
		case model.OrderExpired:
			content = app.OrderExpired(order)
		}
		if !htmx.Request(req) {
			content = app.Base("Checkout Success", content)
		}

		err = content.Render(req.Context(), wr)
		if err != nil {
			slog.Error("App Handler: checkoutSuccess, templ.Base", "error", err)
			http.Error(wr, "Something when wrong", http.StatusInternalServerError)
			return
		}
	}
}

func checkoutCancel(wr http.ResponseWriter, req *http.Request) {
	content := app.CheckoutCancel()
	err := app.Base("Checkout Cancelled", content).Render(req.Context(), wr)
	if err != nil {
		slog.Error("App Handler: index, templ.Base", "error", err)
		http.Error(wr, "Something when wrong", http.StatusInternalServerError)
		return
	}
}

func stripeWebhook(secret string, s model.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		const maxBodyBytes = 65536
		req.Body = http.MaxBytesReader(w, req.Body, maxBodyBytes)
		payload, err := io.ReadAll(req.Body)
		if err != nil {
			slog.Error("Stripe webhook: reading request body", "error", err)
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		event, err := webhook.ConstructEvent(payload, req.Header.Get("Stripe-Signature"), secret)
		if err != nil {
			slog.Error("Stripe webhook: verifying webhook signature", "error", err)
			w.WriteHeader(http.StatusBadRequest) // Return a 400 error on a bad signature
			return
		}

		slog.Info("Stripe webhook: event received",
			"stripe_event_id", event.ID,
			"stripe_event_type", event.Type)

		if slices.Contains(model.StripeEvents, event.Type) {

			_, err = s.River.Insert(req.Context(), model.StripeEvent{Event: event}, nil)
			if err != nil {
				slog.Error("Stripe webhook: unable to insert job",
					"stripe_event_id", event.ID,
					"stripe_event_type", event.Type,
					"error", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		w.WriteHeader(http.StatusOK)
	}
}
