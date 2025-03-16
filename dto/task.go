package dto

import (
	"time"

	"task-manager/domain"

	"github.com/google/uuid"
)

// CreateTaskRequest is used to create a new task.
// @Description Task creation payload
type CreateTaskRequest struct {
	Title       string    `json:"title" binding:"required" example:"Build demo"`
	Description string    `json:"description" example:"Build a task manager demo with Go"`
	Status      string    `json:"status" binding:"required" example:"open"`
	ProjectID   uuid.UUID `json:"project_id" binding:"required" example:"a3d8d6f3-11de-43a0-8e62-330ac6118c15"`
}

// UpdateTaskRequest is used to update an existing task.
// @Description Task update payload
type UpdateTaskRequest struct {
	Title       *string `json:"title,omitempty" example:"Update title"`
	Description *string `json:"description,omitempty" example:"New desc"`
	Status      *string `json:"status,omitempty" example:"done"`
}

// AssignRequest represents the request body for assigning a task to a user.
// @Description Task assignment request DTO
type AssignRequest struct {
	UserID uuid.UUID `json:"user_id" binding:"required" example:"f2bc33e0-103a-4d61-8a67-5ac5084e9fa1"`
}

type TaskResponse struct {
	ID          uuid.UUID `json:"id" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	Title       string    `json:"title" example:"Fix payment bug"`
	Description string    `json:"description" example:"Fix the bug on the payment screen that causes crashes"`
	Status      string    `json:"status" example:"in_progress"`
	AssignedTo  uuid.UUID `json:"assigned_to" example:"c55c8ee2-5552-4b6c-9f49-bb2e3f0d9d22"`
	ProjectID   uuid.UUID `json:"project_id" example:"a3d8d6f3-11de-43a0-8e62-330ac6118c15"`
	CreatedAt   time.Time `json:"created_at" example:"2025-03-13T10:00:00Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2025-03-13T11:30:00Z"`
}

func NewTaskResponse(t domain.Task) TaskResponse {
	return TaskResponse{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status,
		AssignedTo:  t.AssignedTo,
		ProjectID:   t.ProjectID,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

func NewTaskResponseList(tasks []domain.Task) []TaskResponse {
	res := make([]TaskResponse, 0, len(tasks))
	for _, t := range tasks {
		res = append(res, NewTaskResponse(t))
	}
	return res
}
