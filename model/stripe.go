package model

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5"
	"github.com/riverqueue/river"
	"github.com/stripe/stripe-go/v76"
)

var StripeEvents = []stripe.EventType{
	stripe.EventTypeCheckoutSessionCompleted,
	stripe.EventTypeCheckoutSessionExpired,
}

const QueueStripe = "stripe"

type StripeEvent struct {
	Event stripe.Event `json:"event,omitempty"`
}

func (StripeEvent) Kind() string { return "stripe.event" }

func (StripeEvent) InsertOpts() river.InsertOpts {
	return river.InsertOpts{
		Queue: QueueStripe,
	}
}

type StripeEventWorker struct {
	river.WorkerDefaults[StripeEvent]
	service *Service
}

func (w StripeEventWorker) Work(ctx context.Context, job *river.Job[StripeEvent]) error {
	event := job.Args.Event

	if event.Data == nil {
		return river.JobCancel(errors.New("Stripe event missing data field"))
	}

	w.service.Log.Info("Stripe Event processing", "stripe_event_id", event.ID, "stripe_event_type", event.Type)
	switch event.Type {
	case stripe.EventTypeCheckoutSessionCompleted:
		var sess stripe.CheckoutSession
		err := json.Unmarshal(event.Data.Raw, &sess)
		if err != nil {
			return river.JobCancel(err)
		}
		err = w.service.stripeCheckoutSessionEvent(ctx, sess, OrderComplete)
		if err != nil {
			return err
		}
	case stripe.EventTypeCheckoutSessionExpired:
		var sess stripe.CheckoutSession
		err := json.Unmarshal(event.Data.Raw, &sess)
		if err != nil {
			return river.JobCancel(err)
		}
		err = w.service.stripeCheckoutSessionEvent(ctx, sess, OrderExpired)
		if err != nil {
			return err
		}
	default:
		w.service.Log.Info("Stripe Event no handler found", "stripe_event_id", event.ID, "stripe_event_type", event.Type)
		return river.JobCancel(fmt.Errorf("no Stripe event handler for %s", event.Type))
	}

	return nil
}

func (s *Service) stripeCheckoutSessionEvent(ctx context.Context, sess stripe.CheckoutSession, orderStatus OrderStatus) error {

	orderID := uuid.FromStringOrNil(sess.Metadata["order_id"])
	env := sess.Metadata["app_env"]
	if env != s.AppEnv {
		s.Log.Info("Stripe checkout session: wrong env, skipping",
			"stripe_session_id", sess.ID,
			"order_id", orderID,
			"stripe_meta_app_env", env,
			"app_env", s.AppEnv,
		)
		return river.JobCancel(errors.New("wrong env"))
	}

	s.Log.Info("Stripe checkout session event",
		"stripe_session_id", sess.ID,
		"order_id", orderID,
	)
	err := pgx.BeginFunc(ctx, s.DB, func(tx pgx.Tx) error {
		_, err := saveStripeCheckoutSession(ctx, tx, sess, orderID)
		if err != nil {
			return err
		}

		sCustomer := sess.CustomerDetails
		customerID := uuid.Nil
		if sCustomer != nil {
			customer, err := createCustomer(ctx, tx, sCustomer.Name, sCustomer.Email)
			if err != nil {
				return err
			}
			customerID = customer.ID
		}

		_, err = updateOrder(ctx, tx, orderID, orderStatus, customerID)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

const saveStripeCheckoutSessionSQL = `
	insert into stripe.checkout_session (
		id, object, order_id
	) values (
		$1, $2, $3
	)
`

func saveStripeCheckoutSession(ctx context.Context, db DB, sess stripe.CheckoutSession, orderID uuid.UUID) (string, error) {
	object, err := json.Marshal(sess)
	if err != nil {
		return "", err
	}
	_, err = db.Exec(ctx, saveStripeCheckoutSessionSQL,
		sess.ID,
		object,
		orderID,
	)
	return sess.ID, err
}
