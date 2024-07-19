package model

import (
	"context"

	"github.com/riverqueue/river"
)

const QueueMaintenance = "maintenance"

type DeleteExiredOrders struct {
}

func (DeleteExiredOrders) Kind() string { return "delete_expired_orders" }

func (DeleteExiredOrders) InsertOpts() river.InsertOpts {
	return river.InsertOpts{
		Queue: QueueMaintenance,
	}
}

type DeleteExiredOrdersWorker struct {
	river.WorkerDefaults[DeleteExiredOrders]
	service *Service
}

func (w DeleteExiredOrdersWorker) Work(ctx context.Context, job *river.Job[DeleteExiredOrders]) error {
	count, err := deleteExpiredOrders(ctx, w.service.DB)
	if err != nil {
		w.service.Log.Error("Delete Expired Orders", "error", err)
	}
	w.service.Log.Info("Deleted Expired Orders", "order_count", count)
	return err
}
