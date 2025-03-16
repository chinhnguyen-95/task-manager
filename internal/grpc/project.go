package grpc

import (
	"context"

	"task-manager/domain"
	taskmanagerpb "task-manager/pkg/pb/taskmanager"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProjectService interface {
	ListTasks(ctx context.Context, projectID uuid.UUID) ([]domain.Task, error)
}

type ProjectServer struct {
	taskmanagerpb.UnimplementedProjectServiceServer
	service ProjectService
}

func NewProjectServer(service ProjectService) *ProjectServer {
	return &ProjectServer{service: service}
}

func (s *ProjectServer) GetTasks(
	ctx context.Context,
	req *taskmanagerpb.GetProjectTasksRequest,
) (*taskmanagerpb.GetProjectTasksReply, error) {
	projectID, err := uuid.Parse(req.GetProjectId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid project_id: %v", err)
	}

	tasks, err := s.service.ListTasks(ctx, projectID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list project tasks: %v", err)
	}

	var protoTasks []*taskmanagerpb.Task
	for _, task := range tasks {
		protoTasks = append(protoTasks, mapTaskToProto(&task))
	}

	return &taskmanagerpb.GetProjectTasksReply{
		Tasks: protoTasks,
	}, nil
}
