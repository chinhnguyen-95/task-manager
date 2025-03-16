package task

import (
	"context"
	"time"

	"task-manager/domain"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, task *domain.Task) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Task, error)
	Update(ctx context.Context, task *domain.Task) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type CommentRepository interface {
	Create(ctx context.Context, taskID, userID uuid.UUID, content string) error
}

type Service struct {
	repo        Repository
	commentRepo CommentRepository
}

func NewService(repo Repository, commentRepo CommentRepository) *Service {
	return &Service{repo: repo, commentRepo: commentRepo}
}

func (s *Service) Create(ctx context.Context, task *domain.Task) error {
	return s.repo.Create(ctx, task)
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*domain.Task, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) Update(ctx context.Context, task *domain.Task) error {
	return s.repo.Update(ctx, task)
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *Service) Assign(ctx context.Context, taskID, userID uuid.UUID) error {
	task, err := s.repo.GetByID(ctx, taskID)
	if err != nil {
		return err
	}
	task.AssignedTo = userID
	return s.repo.Update(ctx, task)
}

func (s *Service) Comment(ctx context.Context, taskID, userID uuid.UUID, content string) error {
	comment := domain.Comment{
		ID:        uuid.New(),
		TaskID:    taskID,
		UserID:    userID,
		Content:   content,
		CreatedAt: time.Now(),
	}
	return s.commentRepo.Create(ctx, comment.TaskID, comment.UserID, comment.Content)
}
