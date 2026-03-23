package service

import (
	"context"
	"github.com/bestuzheva153/async-task-service/internal/model"
	"github.com/bestuzheva153/async-task-service/internal/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(ctx context.Context, t *model.Task) error {
	t.Status = model.StatusPending
	return s.repo.Create(ctx, t)
}

func (s *TaskService) GetTask(ctx context.Context, id int64) (*model.Task, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *TaskService) GetAllTasks(ctx context.Context) ([]model.Task, error) {
	return s.repo.GetAll(ctx)
}
