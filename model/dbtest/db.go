package dbtest

import (
	"app/internal/pgxtest"
	"app/migrations"
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(ctx context.Context, t *testing.T) *pgxpool.Pool {
	t.Helper()

	db := pgxtest.NewPool(ctx, t, pgxtest.Postgres16)

	err := migrations.RunAll(ctx, db)
	if err != nil {
		t.Fatalf("Could not run all migrations: %s", err)
	}

	return db
}
