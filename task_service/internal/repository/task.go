package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/huandu/go-sqlbuilder"
	"task_service/internal/model"
	"task_service/pkg/storage"
)

const (
	taskTable = "\"gateway\""
)

var (
	taskStruct = sqlbuilder.NewStruct(new(model.Task)).For(sqlbuilder.PostgreSQL)
)

type TaskRepository struct {
	db storage.Postgres
}

func NewTaskRepository(db storage.Postgres) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (r *TaskRepository) Create(ctx context.Context, task *model.Task) (*uuid.UUID, error) {
	var id uuid.UUID
	sql, args := taskStruct.WithoutTag("pk").InsertInto(taskTable, task).SQL("RETURNING id").Build()
	fmt.Println(sql)
	err := r.db.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (r *TaskRepository) Update(ctx context.Context, task *model.Task) error {
	u := taskStruct.WithoutTag("pk").Update(taskTable, task)
	u.Where(u.Equal("id", task.Id))
	sql, args := u.Build()
	fmt.Println(sql)
	_, err := r.db.Exec(ctx, sql, args...)
	return err
}

func (r *TaskRepository) GetAll(ctx context.Context, limit int, offset int) ([]*model.Task, error) {
	sb := taskStruct.SelectFrom(taskTable)
	sb.Limit(limit).Offset(offset).OrderBy("id")
	sql, args := sb.Build()
	fmt.Println(sql)
	var data []*model.Task
	err := r.db.Select(ctx, &data, sql, args...)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *TaskRepository) GetById(ctx context.Context, id uuid.UUID) (*model.Task, error) {
	sb := taskStruct.SelectFrom(taskTable)
	sb.Where(sb.Equal("id", id.String()))
	sql, args := sb.Build()
	fmt.Println(sql)
	var data []model.Task
	err := r.db.Select(ctx, &data, sql, args...)
	if err != nil {
		return nil, err
	}
	if len(data) != 1 {
		return nil, fmt.Errorf("task count error")
	}
	return &data[0], nil
}
