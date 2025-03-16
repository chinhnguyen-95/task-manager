package grpc

import (
	"context"

	"task-manager/domain"
	taskmanagerpb "task-manager/pkg/pb/taskmanager"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService interface {
	ListTasks(ctx context.Context, userID uuid.UUID) ([]domain.Task, error)
}

type UserServer struct {
	taskmanagerpb.UnimplementedUserServiceServer
	service UserService
}

func NewUserServer(service UserService) *UserServer {
	return &UserServer{service: service}
}

func (s *UserServer) GetTasks(
	ctx context.Context,
	req *taskmanagerpb.GetUserTasksRequest,
) (*taskmanagerpb.GetUserTasksReply, error) {
	userID, err := uuid.Parse(req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_id: %v", err)
	}

	tasks, err := s.service.ListTasks(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list user tasks: %v", err)
	}

	var protoTasks []*taskmanagerpb.Task
	for _, task := range tasks {
		protoTasks = append(protoTasks, mapTaskToProto(&task))
	}

	return &taskmanagerpb.GetUserTasksReply{
		Tasks: protoTasks,
	}, nil
}
