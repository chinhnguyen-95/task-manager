package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	taskmanagerpb "task-manager/pkg/pb/taskmanager"
)

type Server = grpc.Server

func NewServer(
	jwtInterceptor grpc.UnaryServerInterceptor,
	authSvc AuthService,
	taskSvc TaskService,
	userSvc UserService,
	projectSvc ProjectService,
) *Server {
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			jwtInterceptor,
		),
	)

	reflection.Register(grpcServer)

	taskmanagerpb.RegisterAuthServiceServer(grpcServer, NewAuthServer(authSvc))
	taskmanagerpb.RegisterTaskServiceServer(grpcServer, NewTaskServer(taskSvc))
	taskmanagerpb.RegisterUserServiceServer(grpcServer, NewUserServer(userSvc))
	taskmanagerpb.RegisterProjectServiceServer(grpcServer, NewProjectServer(projectSvc))

	return grpcServer
}
