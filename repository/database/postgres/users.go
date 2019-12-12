package postgres

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/rodrwan/twoppin"
)

type UsersService struct {
	Store SQLExecutor
}

func (us *UsersService) Get(ctx context.Context, id string) (*twoppin.User, error) {
	q := squirrel.Select("*").From("users").Where("deleted_at is null").Where("id = ?", id)

	sqlString, args, err := q.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var user *twoppin.User
	row := us.Store.QueryRowxContext(ctx, sqlString, args...)
	if err := row.StructScan(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UsersService) Select(ctx context.Context) ([]*twoppin.User, error) {
	return nil, errors.New("not implemented")
}

func (us *UsersService) Create(ctx context.Context, u *twoppin.User) error {
	return errors.New("not implemented")
}

func (us *UsersService) Update(ctx context.Context, u *twoppin.User) error {
	return errors.New("not implemented")
}

func (us *UsersService) Delete(ctx context.Context, u *twoppin.User) error {
	return errors.New("not implemented")
}
