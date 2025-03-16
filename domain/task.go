package domain

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID
	Title       string
	Description string
	Status      string
	AssignedTo  uuid.UUID
	ProjectID   uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
