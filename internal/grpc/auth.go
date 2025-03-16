package grpc

import (
	"context"

	"task-manager/dto"
	taskmanagerpb "task-manager/pkg/pb/taskmanager"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthService interface {
	Register(ctx context.Context, req dto.RegisterRequest) error
	Login(ctx context.Context, req dto.LoginRequest) (string, error)
}

type AuthServer struct {
	taskmanagerpb.UnimplementedAuthServiceServer
	authService AuthService
}

func NewAuthServer(authService AuthService) *AuthServer {
	return &AuthServer{authService: authService}
}

func (s *AuthServer) Register(ctx context.Context, req *taskmanagerpb.RegisterRequest) (
	*taskmanagerpb.SuccessResponse,
	error,
) {
	// Convert gRPC request to DTO
	registerReq := dto.RegisterRequest{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	// Call service
	if err := s.authService.Register(ctx, registerReq); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &taskmanagerpb.SuccessResponse{
		Message: "User registered successfully",
	}, nil
}

func (s *AuthServer) Login(ctx context.Context, req *taskmanagerpb.LoginRequest) (*taskmanagerpb.LoginReply, error) {
	// Convert gRPC request to DTO
	loginReq := dto.LoginRequest{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	// Call service
	token, err := s.authService.Login(ctx, loginReq)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &taskmanagerpb.LoginReply{
		AccessToken: token,
	}, nil
}
