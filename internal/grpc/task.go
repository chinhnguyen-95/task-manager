package grpc

import (
	"context"
	"time"

	"task-manager/domain"
	"task-manager/internal/grpc/middleware"
	taskmanagerpb "task-manager/pkg/pb/taskmanager"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TaskService interface {
	Create(ctx context.Context, task *domain.Task) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Task, error)
	Update(ctx context.Context, task *domain.Task) error
	Delete(ctx context.Context, id uuid.UUID) error
	Assign(ctx context.Context, taskID uuid.UUID, userID uuid.UUID) error
	Comment(ctx context.Context, taskID, userID uuid.UUID, content string) error
}

type TaskServer struct {
	taskmanagerpb.UnimplementedTaskServiceServer
	service TaskService
}

func NewTaskServer(service TaskService) *TaskServer {
	return &TaskServer{service: service}
}

func (s *TaskServer) CreateTask(
	ctx context.Context,
	req *taskmanagerpb.CreateTaskRequest,
) (*taskmanagerpb.CreateTaskReply, error) {
	projectID, err := uuid.Parse(req.GetProjectId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid project_id: %v", err)
	}

	userIDStr, ok := middleware.GetUserIDFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "user ID not found in context")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid user ID format")
	}

	task := &domain.Task{
		ID:          uuid.New(),
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Status:      req.GetStatus(),
		ProjectID:   projectID,
		AssignedTo:  userID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.service.Create(ctx, task); err != nil {
		return nil, status.Errorf(codes.Internal, "create task failed: %v", err)
	}

	return &taskmanagerpb.CreateTaskReply{
		Task: mapTaskToProto(task),
	}, nil
}

func (s *TaskServer) GetTaskByID(ctx context.Context, req *taskmanagerpb.GetTaskRequest) (
	*taskmanagerpb.GetTaskReply,
	error,
) {
	id, err := uuid.Parse(req.GetTaskId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid task_id: %v", err)
	}

	task, err := s.service.GetByID(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get task failed: %v", err)
	}

	return &taskmanagerpb.GetTaskReply{
		Task: mapTaskToProto(task),
	}, nil
}

func (s *TaskServer) UpdateTaskByID(
	ctx context.Context,
	req *taskmanagerpb.UpdateTaskRequest,
) (*taskmanagerpb.UpdateTaskReply, error) {
	id, err := uuid.Parse(req.GetTaskId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid task_id: %v", err)
	}

	task := &domain.Task{
		ID:          id,
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Status:      req.GetStatus(),
		UpdatedAt:   time.Now(),
	}

	if err := s.service.Update(ctx, task); err != nil {
		return nil, status.Errorf(codes.Internal, "update task failed: %v", err)
	}

	return &taskmanagerpb.UpdateTaskReply{
		Task: mapTaskToProto(task),
	}, nil
}

func (s *TaskServer) DeleteTaskByID(
	ctx context.Context,
	req *taskmanagerpb.DeleteTaskRequest,
) (*taskmanagerpb.SuccessResponse, error) {
	id, err := uuid.Parse(req.GetTaskId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid task_id: %v", err)
	}

	if err := s.service.Delete(ctx, id); err != nil {
		return nil, status.Errorf(codes.Internal, "delete task failed: %v", err)
	}

	return &taskmanagerpb.SuccessResponse{Message: "Task deleted successfully"}, nil
}

func (s *TaskServer) AssignTaskToUser(
	ctx context.Context,
	req *taskmanagerpb.AssignTaskRequest,
) (*taskmanagerpb.SuccessResponse, error) {
	taskID, err := uuid.Parse(req.GetTaskId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid task_id: %v", err)
	}

	userID, err := uuid.Parse(req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_id: %v", err)
	}

	if err := s.service.Assign(ctx, taskID, userID); err != nil {
		return nil, status.Errorf(codes.Internal, "assign failed: %v", err)
	}

	return &taskmanagerpb.SuccessResponse{Message: "Task assigned successfully"}, nil
}

func (s *TaskServer) CommentOnTask(
	ctx context.Context,
	req *taskmanagerpb.CommentTaskRequest,
) (*taskmanagerpb.SuccessResponse, error) {
	taskID, err := uuid.Parse(req.GetTaskId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid task_id: %v", err)
	}

	userIDStr, ok := middleware.GetUserIDFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "user ID not found in context")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid user ID format")
	}

	if err := s.service.Comment(ctx, taskID, userID, req.GetContent()); err != nil {
		return nil, status.Errorf(codes.Internal, "comment failed: %v", err)
	}

	return &taskmanagerpb.SuccessResponse{Message: "Comment added successfully"}, nil
}
