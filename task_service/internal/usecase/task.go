package usecase

import (
	"context"
	"task_service/internal/model"
	"task_service/internal/repository"
)

type TaskUC struct {
	taskRepo  *repository.TaskRepository
	jwtSecret string
}

func NewTaskUC(taskRepo *repository.TaskRepository, jwtSecret string) *TaskUC {
	return &TaskUC{taskRepo: taskRepo, jwtSecret: jwtSecret}
}

func (u *TaskUC) Create(ctx context.Context, in model.CreateTaskInput) (*model.CreateTaskOutput, error) {
	_, err := u.taskRepo.Create(ctx, &model.Task{
		AuthorId: in.GetAuthorId(),
		Title:    in.GetTitle(),
		Content:  in.GetContent(),
		Status:   in.GetStatus(),
	})
	if err != nil {
		return nil, err
	}
	return &model.CreateTaskOutput{Success: true}, nil
}
