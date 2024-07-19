package main

import (
	"app/migrations"
	"app/migrations/code"
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DBPoolConnect connect to a pgx Pool
//
// # Supported enviroment variables
//
// DATABASE_URL
// PGHOST
// PGPORT
// PGDATABASE
// PGUSER
// PGPASSWORD
// PGPASSFILE
// PGSERVICE
// PGSERVICEFILE
// PGSSLMODE
// PGSSLCERT
// PGSSLKEY
// PGSSLROOTCERT
// PGSSLPASSWORD
// PGAPPNAME
// PGCONNECT_TIMEOUT
// PGTARGETSESSIONATTRS
func DBPoolConnect(ctx context.Context, connString string, autoMigrate bool) (*pgxpool.Pool, error) {
	if connString == "" {
		connString = os.Getenv("DATABASE_URL")
	}

	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("database pool connection: %w", err)
	}

	slog.Info("Database Pool: successful connection",
		"host", pool.Config().ConnConfig.Config.Host,
		"port", pool.Config().ConnConfig.Config.Port,
		"user", pool.Config().ConnConfig.Config.User,
		"database", pool.Config().ConnConfig.Config.Database,
	)

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("database ping error: %w", err)
	}
	slog.Info("Database Pool: ping success")

	conn, err := pool.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("database acquire connection: %w", err)
	}
	defer conn.Release()

	if autoMigrate {
		err = migrations.RunAll(ctx, pool)
		if err != nil {
			return nil, fmt.Errorf("migration run all: %w", err)
		}
		slog.Info("Database Pool: all migrations success")
	}

	err = code.Run(ctx, conn.Conn())
	if err != nil {
		return nil, fmt.Errorf("code migration: %w", err)
	}
	slog.Info("Database Pool: code migration success")

	return pool, nil
}

func DBPooleClose(pool *pgxpool.Pool) {

	slog.Info("Database Pool: connection closing")
	pool.Close()
	slog.Info("Database Pool: connection closed")
}
