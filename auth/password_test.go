package auth

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	HashCost = MinHashCost

	hashed, err := HashPassword("guess")
	if err != nil {
		t.Fatal(err)
	}

	match := CheckPasswordHash("guess", hashed)
	if !match {
		t.Fatal("Password hash does not match")
	}
}
