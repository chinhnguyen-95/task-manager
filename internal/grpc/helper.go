package grpc

import (
	"time"

	"task-manager/domain"
	taskmanagerpb "task-manager/pkg/pb/taskmanager"
)

func mapTaskToProto(t *domain.Task) *taskmanagerpb.Task {
	return &taskmanagerpb.Task{
		Id:          t.ID.String(),
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status,
		ProjectId:   t.ProjectID.String(),
		AssignedTo:  t.AssignedTo.String(),
		CreatedAt:   t.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   t.UpdatedAt.Format(time.RFC3339),
	}
}
