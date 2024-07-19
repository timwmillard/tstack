package model

import (
	"app/internal/sqltest"
	"app/model/dbtest"
	"context"
	"log"
	"testing"
)

func TestRegister(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping TestRegister")
	}
	t.Parallel()

	ctx := context.Background()
	db := dbtest.NewPool(ctx, t)

	s := Service{DB: db}

	// SUT
	_, err := s.Register(ctx, "Tim", "tim@test.com", "pass")
	if err != nil {
		log.Fatalf("Could not create user: %v", err)
	}

	sqltest := sqltest.New(t, db)
	sqltest.AssertQuery(ctx, "testdata/register.sql")
}
