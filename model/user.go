package model

import (
	"app/auth"
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5"
)

var ErrInvalidLogin = errors.New("invalid login")

// TODO: write a test for this function
func splitName(name string) (firstName string, lastName string) {
	names := strings.Split(name, " ")
	if len(names) > 1 {
		firstName = strings.Join(names[:len(names)-1], " ")
		lastName = names[len(names)-1]
	} else {
		firstName = name
	}
	return firstName, lastName
}

func (s Service) Register(ctx context.Context, name, email, password string) (auth.User, error) {
	firstName, lastName := splitName(name)
	password, err := auth.HashPassword(password)
	if err != nil {
		return auth.User{}, err
	}
	user, err := createAdminUser(ctx, s.DB, email, password, firstName, lastName)
	if err != nil {
		return auth.User{}, err
	}

	return auth.User{
		ID: user.ID,
	}, nil
}

func (s Service) ListAdminUsers(ctx context.Context) ([]AdminUser, error) {
	return listAdminUsers(ctx, s.DB)
}

func (s Service) Login(ctx context.Context, username, password string) (auth.User, error) {

	user, err := getAdminUserByUsername(ctx, s.DB, username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return auth.User{}, ErrInvalidLogin
		}
		return auth.User{}, err
	}

	if !auth.CheckPasswordHash(password, user.Password) {
		return auth.User{}, ErrInvalidLogin
	}

	return auth.User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Type:      auth.Admin,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil
}

func (s Service) GetUser(ctx context.Context, userID uuid.UUID) (auth.User, error) {
	user, err := getAdminUser(ctx, s.DB, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return auth.User{}, ErrInvalidLogin
		}
		return auth.User{}, err
	}
	return auth.User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Type:      auth.Admin,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil
}

type User struct {
	ID       uuid.UUID
	Username string
	Password string
	Email    string
}

type AdminUser struct {
	ID        uuid.UUID
	Username  string
	Password  string
	Email     string
	FirstName string
	LastName  string
}

// SQL
const listAdminUsersSQL = `
select
	u.id, 
	u.username,
	coalesce(u.email, '') as email,
	admin.first_name,
	admin.last_name
from auth.user u
join app.admin admin on (u.id = admin.user_id)
`

// listAdminUsers -
func listAdminUsers(ctx context.Context, db Queryer) ([]AdminUser, error) {
	rows, _ := db.Query(ctx, listAdminUsersSQL)
	user, err := pgx.CollectRows[AdminUser](rows, pgx.RowToStructByNameLax)
	return user, err
}

// SQL
const createUserSQL = `
insert into auth.user (
	username,
	password,
	email
) values (
	$1, $2, $1
) returning id, username, password, email
`

// SQL
const createAdminSQL = `
insert into app.admin (
	user_id,
	first_name,
	last_name
) values (
	$1, $2, $3
) returning user_id, first_name, last_name 
`

// createAdminUser -
func createAdminUser(ctx context.Context, db Queryer, email, password, firstName, lastName string) (AdminUser, error) {
	rows, _ := db.Query(ctx, createUserSQL, email, password)
	user, err := pgx.CollectOneRow[User](rows, pgx.RowToStructByNameLax)
	if err != nil {
		return AdminUser{}, fmt.Errorf("create user error: %w", err)
	}

	type admin struct {
		UserID    uuid.UUID
		FirstName string
		LastName  string
		Email     string
	}
	rows, _ = db.Query(ctx, createAdminSQL, user.ID, firstName, lastName)
	a, err := pgx.CollectOneRow[admin](rows, pgx.RowToStructByNameLax)
	if err != nil {
		return AdminUser{}, fmt.Errorf("create admin error: %w", err)
	}

	return AdminUser{
		ID:        user.ID,
		Username:  user.Username,
		Password:  user.Password,
		Email:     a.Email,
		FirstName: a.FirstName,
		LastName:  a.LastName,
	}, nil
}

// SQL
const getAdminUserByUsernameSQL = `
select
	u.id,
	username,
	password,
	first_name,
	last_name
from auth.user u
join app.admin admin on u.ID = admin.user_id
where username = $1
`

// getAdminUserByUsername get an admin user by username.
func getAdminUserByUsername(ctx context.Context, db Queryer, username string) (AdminUser, error) {
	rows, _ := db.Query(ctx, getAdminUserByUsernameSQL, username)
	user, err := pgx.CollectOneRow[AdminUser](rows, pgx.RowToStructByNameLax)
	return user, err
}

// SQL
const getUserSQL = `
select
	u.id,
	username,
	password,
	first_name,
	last_name
from auth.user u
join app.admin admin on u.ID = admin.user_id
where id = $1
`

// getAdminUser get an admin user by id.
func getAdminUser(ctx context.Context, db Queryer, id uuid.UUID) (AdminUser, error) {
	rows, _ := db.Query(ctx, getUserSQL, id)
	user, err := pgx.CollectOneRow[AdminUser](rows, pgx.RowToStructByNameLax)
	return user, err
}
