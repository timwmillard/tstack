package auth

import (
	"context"
	"strings"

	"github.com/gofrs/uuid/v5"
)

// Define the context key type.
type contextKey string

// Create a context key for the user
var userContextKey contextKey = "user"

func WithContextUser(ctx context.Context, user User) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}

func UserFromContext(ctx context.Context) User {
	if user, ok := ctx.Value(userContextKey).(User); ok {
		return user
	}
	return User{}
}

const SessionUserID = "user_id"

type UserType byte

const (
	Admin UserType = 1 << iota
	Customer
)

type User struct {
	ID        uuid.UUID
	Username  string
	Email     string
	Type      UserType
	FirstName string
	LastName  string
}

func (u User) Name() string {
	if u.FirstName == "" || u.LastName == "" {
		return u.Username
	}
	return strings.TrimSpace(u.FirstName + " " + u.LastName)
}

func (u User) IsAdmin() bool {
	return u.Type&Admin == Admin
}

func (u User) IsCustomer() bool {
	return u.Type&Customer == Customer
}
