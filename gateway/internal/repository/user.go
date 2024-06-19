package repository

import (
	"context"
	"fmt"
	"gateway/internal/model/user"
	"gateway/pkg/storage"
	"github.com/google/uuid"
	"github.com/huandu/go-sqlbuilder"
)

const (
	userTable = "\"user\""
)

var (
	userStruct = sqlbuilder.NewStruct(new(user.User)).For(sqlbuilder.PostgreSQL)
)

type UserRepository struct {
	db storage.Postgres
}

func NewUserRepository(db storage.Postgres) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *user.User) (*uuid.UUID, error) {
	var id uuid.UUID
	sql, args := userStruct.WithoutTag("pk").InsertInto(userTable, user).SQL("RETURNING id").Build()
	fmt.Println(sql)
	err := r.db.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (r *UserRepository) Update(ctx context.Context, user *user.User) error {
	u := userStruct.WithoutTag("pk").Update(userTable, user)
	u.Where(u.Equal("id", user.ID))
	sql, args := u.Build()
	fmt.Println(sql)
	_, err := r.db.Exec(ctx, sql, args...)
	return err
}

func (r *UserRepository) GetByLogin(ctx context.Context, login string) (*user.User, error) {
	sb := userStruct.SelectFrom(userTable)
	sb.Where(sb.Equal("login", login))
	sql, args := sb.Build()
	fmt.Println(sql)
	var data []user.User
	err := r.db.Select(ctx, &data, sql, args...)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil
	}
	if len(data) > 1 {
		return nil, fmt.Errorf("user count error")
	}
	return &data[0], nil
}

func (r *UserRepository) GetById(ctx context.Context, id uuid.UUID) (*user.User, error) {
	sb := userStruct.SelectFrom(userTable)
	sb.Where(sb.Equal("id", id.String()))
	sql, args := sb.Build()
	fmt.Println(sql)
	var data []user.User
	err := r.db.Select(ctx, &data, sql, args...)
	if err != nil {
		return nil, err
	}
	if len(data) != 1 {
		return nil, fmt.Errorf("user count error")
	}
	return &data[0], nil
}
