package task

import (
	"context"
	"testing"

	"task-manager/domain"
	"task-manager/task/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_Create(t *testing.T) {
	mockRepo := new(mocks.Repository)
	mockCommentRepo := new(mocks.CommentRepository)
	svc := NewService(mockRepo, mockCommentRepo)

	tk := &domain.Task{
		ID:    uuid.New(),
		Title: "Test Task",
	}

	mockRepo.On("Create", context.Background(), tk).Return(nil)

	err := svc.Create(context.Background(), tk)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestService_Assign(t *testing.T) {
	mockRepo := new(mocks.Repository)
	mockCommentRepo := new(mocks.CommentRepository)
	svc := NewService(mockRepo, mockCommentRepo)

	taskID := uuid.New()
	userID := uuid.New()

	existing := &domain.Task{ID: taskID}

	mockRepo.On("GetByID", context.Background(), taskID).Return(existing, nil)
	mockRepo.On(
		"Update", context.Background(), mock.MatchedBy(
			func(t *domain.Task) bool {
				return t.ID == taskID && t.AssignedTo == userID
			},
		),
	).Return(nil)

	err := svc.Assign(context.Background(), taskID, userID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestService_Comment(t *testing.T) {
	mockRepo := new(mocks.Repository)
	mockCommentRepo := new(mocks.CommentRepository)
	svc := NewService(mockRepo, mockCommentRepo)

	taskID := uuid.New()
	userID := uuid.New()
	content := "This is awesome!"

	mockCommentRepo.On("Create", context.Background(), taskID, userID, content).Return(nil)

	err := svc.Comment(context.Background(), taskID, userID, content)

	assert.NoError(t, err)
	mockCommentRepo.AssertExpectations(t)
}
