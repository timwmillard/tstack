package model

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Queryer interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
}

type Execer interface {
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}

// DB is the interface access the database. It is satisfied by *pgx.Conn, pgx.Tx, *pgxpool.Pool, etc.
type DB interface {
	// Begin starts a new pgx.Tx. It may be a true transaction or a pseudo nested transaction implemented by savepoints.
	Begin(ctx context.Context) (pgx.Tx, error)

	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	SendBatch(ctx context.Context, b *pgx.Batch) (br pgx.BatchResults)
}
