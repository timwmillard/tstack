package migrations

import (
	"context"
	"embed"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/tern/v2/migrate"
	"github.com/riverqueue/river/riverdriver/riverpgxv5"
	"github.com/riverqueue/river/rivermigrate"
)

//go:embed * code/*
var FS embed.FS

func RunAll(ctx context.Context, pool *pgxpool.Pool) error {

	_, err := pool.Exec(ctx, "create schema if not exists app")
	if err != nil {
		return fmt.Errorf("could create schema app: %w", err)
	}

	err = RunRiver(ctx, pool)
	if err != nil {
		return fmt.Errorf("could not run river migrations: %w", err)
	}

	conn, err := pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("could not acquire connection: %w", err)
	}
	defer conn.Release()

	err = Run(ctx, conn.Conn())
	if err != nil {
		return fmt.Errorf("could not run migration: %w", err)
	}
	return nil
}

func RunRiver(ctx context.Context, db *pgxpool.Pool) error {
	riverMigrator := rivermigrate.New(riverpgxv5.New(db), nil)
	_, err := riverMigrator.Migrate(ctx, rivermigrate.DirectionUp, nil)
	if err != nil {
		return err
	}
	return nil
}

func Run(ctx context.Context, conn *pgx.Conn) error {

	mig, err := migrate.NewMigrator(ctx, conn, "schema_version")
	if err != nil {
		return err
	}
	err = mig.LoadMigrations(FS)
	if err != nil {
		return err
	}
	err = mig.Migrate(ctx)
	if err != nil {
		return err
	}

	return nil
}
