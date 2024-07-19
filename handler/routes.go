package handler

import (
	"app/model"

	"net/http"

	"app/handler/middleware"

	"github.com/alexedwards/scs/v2"
	"github.com/gorilla/mux"
)

func AppRoutes(path string, r *mux.Router, config Config, service model.Service, sess *scs.SessionManager) {

	appRouter := r.PathPrefix(path).Subrouter()
	appRouter.Use(sess.LoadAndSave)

	appRouter.HandleFunc("/", index(service)).Methods(http.MethodGet)

	// Checkout
	appRouter.HandleFunc("/checkout", createCheckoutSession(config.Host, service, sess)).Methods(http.MethodPost)
	appRouter.HandleFunc("/checkout/buynow", createCheckoutSessionBuyNow(config.Host, service)).Methods(http.MethodPost)
	appRouter.HandleFunc("/checkout/cancel/", checkoutCancel).Methods(http.MethodGet)

	// Order
	appRouter.HandleFunc("/orders/{order_id}", checkoutSuccess(service, sess)).Methods(http.MethodGet)

	// Stripe Webhooks
	appRouter.HandleFunc("/stripe/webhook/", stripeWebhook(config.Stripe.WebhookSecret, service)).Methods(http.MethodPost)
}

// | Verb		| URI					| Action	| Route Name	 |
// |------------|-----------------------|-----------|--------------- |
// | GET		| /photos				| index		| photos.index	 |
// | GET		| /photos/create		| create	| photos.create	 |
// | POST		| /photos				| store		| photos.store	 |
// | GET		| /photos/{photo}		| show		| photos.show	 |
// | GET		| /photos/{photo}/edit	| edit		| photos.edit	 |
// | PUT/PATCH	| /photos/{photo}		| update	| photos.update  |
// | DELETE		| /photos/{photo}		| destroy	| photos.destroy |

func AdminRoutes(path string, r *mux.Router, config Config, srv model.Service, sess *scs.SessionManager) {

	adminR := r.PathPrefix(path).Subrouter()
	adminR.NotFoundHandler = adminNotFound()
	adminR.Use(sess.LoadAndSave)

	// Login
	adminR.HandleFunc("/login", adminLoginForm("")).Methods(http.MethodGet)
	adminR.HandleFunc("/login", adminLogin(srv, sess)).Methods(http.MethodPost)

	authR := adminR.NewRoute().Subrouter()
	authR.Use(middleware.Auth(srv, sess))

	// Logout
	adminR.HandleFunc("/logout", adminLogout(sess)).Methods(http.MethodPost)

	// Event
	authR.HandleFunc("/", adminIndex(srv)).Methods(http.MethodGet)
}
