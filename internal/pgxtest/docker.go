package pgxtest

// References:
// GopherCon 2020: Johan Brandhorst-Satzkorn - A Journey to Postgres Productivity with Go
// https://www.youtube.com/watch?v=AgHdVPSty7k

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

type Container struct {
	Repository string
	Tag        string
}

var (
	Postgres14 = Container{
		Repository: "postgres",
		Tag:        "14",
	}
	Postgres15 = Container{
		Repository: "postgres",
		Tag:        "15",
	}
	Postgres16 = Container{
		Repository: "postgres",
		Tag:        "16",
	}
)

func NewConn(t *testing.T, ctx context.Context, container Container) *pgx.Conn {
	t.Helper()

	pool, resource := newDockerContainer(t, container)

	var db *pgx.Conn
	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		port := resource.GetPort("5432/tcp")
		connString := fmt.Sprintf("postgres://postgres:secret@localhost:%s/postgres", port)
		db, err = pgx.Connect(ctx, connString)
		if err != nil {
			return err
		}

		err = db.Ping(ctx)
		if err != nil {
			return err
		}

		t.Logf("Successful connection %s", connString)

		return nil
	}); err != nil {
		t.Fatalf("Could not connect to database: %s", err)
	}

	t.Cleanup(func() {
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	})

	return db
}

// TODO: need to write cleaup funcs
func NewPool(ctx context.Context, t *testing.T, container Container) *pgxpool.Pool {
	t.Helper()

	pool, resource := newDockerContainer(t, container)

	var db *pgxpool.Pool
	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		port := resource.GetPort("5432/tcp")
		connString := fmt.Sprintf("postgres://postgres:secret@localhost:%s/postgres", port)
		db, err = pgxpool.New(ctx, connString)
		if err != nil {
			return err
		}

		err = db.Ping(ctx)
		if err != nil {
			return err
		}

		t.Logf("Successful connection %s", connString)

		return nil
	}); err != nil {
		t.Fatalf("Could not connect to database: %s", err)
	}

	t.Cleanup(func() {
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	})

	return db
}

func newDockerContainer(t *testing.T, container Container) (*dockertest.Pool, *dockertest.Resource) {
	t.Helper()

	pool, err := dockertest.NewPool("")
	if err != nil {
		t.Fatalf("Could not construct pool: %s", err)
	}

	// uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: container.Repository,
		Tag:        container.Tag,
		Env: []string{
			"POSTGRES_PASSWORD=secret",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		t.Fatalf("Could not start resource: %s", err)
	}
	pool.MaxWait = 120 * time.Second
	return pool, resource
}
