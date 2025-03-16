package user

import (
	"context"

	"task-manager/domain"

	"github.com/google/uuid"
)

type TaskRepository interface {
	ListByUser(ctx context.Context, userID uuid.UUID) ([]domain.Task, error)
}

type Service struct {
	taskRepo TaskRepository
}

func NewService(taskRepo TaskRepository) *Service {
	return &Service{taskRepo: taskRepo}
}

func (s *Service) ListTasks(ctx context.Context, userID uuid.UUID) ([]domain.Task, error) {
	return s.taskRepo.ListByUser(ctx, userID)
}
