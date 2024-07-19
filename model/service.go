package model

import (
	"context"
	"io"
	"log/slog"
	"net/url"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/riverdriver/riverpgxv5"
)

type Storage interface {
	GetURL(filename string) string
	Get(ctx context.Context, filename string, wr io.Writer) error
	Upload(ctx context.Context, filename string, rd io.Reader) (string, error)
	Delete(ctx context.Context, filename string) error
}

type Service struct {
	DB   *pgxpool.Pool
	Host url.URL

	Log slog.Logger

	PublicStore  Storage
	PrivateStore Storage
	River        *river.Client[pgx.Tx]
	AppEnv       string
}

func (s *Service) RiverStart(ctx context.Context) error {

	stripeEventWorker := StripeEventWorker{service: s}
	deleteExpiredOrdersWorker := DeleteExiredOrdersWorker{service: s}

	workers := river.NewWorkers()
	river.AddWorker(workers, stripeEventWorker)
	river.AddWorker(workers, deleteExpiredOrdersWorker)

	periodicJobs := []*river.PeriodicJob{
		river.NewPeriodicJob(
			river.PeriodicInterval(1*time.Hour),
			func() (river.JobArgs, *river.InsertOpts) {
				return DeleteExiredOrders{}, nil
			},
			&river.PeriodicJobOpts{RunOnStart: true},
		),
	}
	riverClient, err := river.NewClient(riverpgxv5.New(s.DB), &river.Config{
		Queues: map[string]river.QueueConfig{
			river.QueueDefault: {MaxWorkers: 10},
			QueueStripe:        {MaxWorkers: 50},
			QueueMaintenance:   {MaxWorkers: 5},
		},
		Workers:      workers,
		PeriodicJobs: periodicJobs,
	})
	if err != nil {
		return err
	}
	err = riverClient.Start(ctx)
	if err != nil {
		return err
	}
	s.Log.Info("Service: River started, ready to process jobs")

	s.River = riverClient
	return nil
}
func (s *Service) RiverStop(ctx context.Context) error {
	err := s.River.Stop(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) CreateOrder(ctx context.Context) (Order, error) {
	order, err := createOrder(ctx, s.DB)
	if err != nil {
		return Order{}, err
	}
	return order, err
}

func (s Service) CreateFullOrder(ctx context.Context, items []string) (Order, error) {
	order, err := createFullOrder(ctx, s.DB, items)
	if err != nil {
		return Order{}, err
	}
	return order, err
}

func (s Service) GetOrder(ctx context.Context, id uuid.UUID) (Order, error) {
	order, err := getOrderWithItems(ctx, s.DB, id)
	if err != nil {
		return Order{}, err
	}
	return order, err
}
