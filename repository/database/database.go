package database

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rodrwan/twoppin"
	"github.com/rodrwan/twoppin/repository/database/postgres"
)

// UsersService is the interface that represent basic operation over
// users table in any database.
type UsersService interface {
	Get(context.Context, string) (*twoppin.User, error)
	Select(context.Context) ([]*twoppin.User, error)

	Create(context.Context, *twoppin.User) error
	Update(context.Context, *twoppin.User) error
	Delete(context.Context, *twoppin.User) error
}

// Database hold the database services.
type Database struct {
	UsersService UsersService
}

// NewPostgres creates a new Database with postgres as driver.
func NewPostgres(dsn string) (*Database, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &Database{
		UsersService: &postgres.UsersService{
			Store: db,
		},
	}, nil
}
