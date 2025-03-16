package model

import (
	"time"

	"task-manager/domain"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Title       string
	Description string
	Status      string
	AssignedTo  uuid.UUID
	ProjectID   uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTaskModel(t domain.Task) Task {
	return Task{
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

func (m Task) ToDomain() domain.Task {
	return domain.Task{
		ID:          m.ID,
		Title:       m.Title,
		Description: m.Description,
		Status:      m.Status,
		AssignedTo:  m.AssignedTo,
		ProjectID:   m.ProjectID,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}
