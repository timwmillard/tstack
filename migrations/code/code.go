package code

import (
	"context"
	"embed"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/tern/v2/migrate"
)

//go:embed *
var FS embed.FS

func Run(ctx context.Context, conn *pgx.Conn) error {
	codePkg, err := migrate.LoadCodePackage(FS)
	if err != nil {
		return err
	}
	err = migrate.InstallCodePackage(ctx, conn, nil, codePkg)
	if err != nil {
		return err
	}
	return nil
}
