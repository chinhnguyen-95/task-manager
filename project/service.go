package project

import (
	"context"

	"task-manager/domain"

	"github.com/google/uuid"
)

type TaskRepository interface {
	ListByProject(ctx context.Context, projectID uuid.UUID) ([]domain.Task, error)
}

type Service struct {
	taskRepository TaskRepository
}

func NewService(taskRepository TaskRepository) *Service {
	return &Service{taskRepository: taskRepository}
}

func (s Service) ListTasks(ctx context.Context, projectID uuid.UUID) ([]domain.Task, error) {
	return s.taskRepository.ListByProject(ctx, projectID)
}
