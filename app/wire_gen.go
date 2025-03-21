// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"task-manager/auth"
	"task-manager/internal/grpc"
	middleware2 "task-manager/internal/grpc/middleware"
	"task-manager/internal/repository/postgres"
	"task-manager/internal/rest"
	"task-manager/internal/rest/middleware"
	"task-manager/pkg/jwtutil"
	"task-manager/pkg/keycloak"
	"task-manager/project"
	"task-manager/task"
	"task-manager/user"
)

import (
	_ "task-manager/docs"
)

// Injectors from wire.go:

func InitializeApp() (*App, error) {
	publicKey, err := jwtutil.FetchRSAPublicKeyFromJWKS()
	if err != nil {
		return nil, err
	}
	handlerFunc := middleware.JWTAuthMiddleware(publicKey)
	client := keycloak.NewClient()
	service := auth.NewService(client)
	db, err := postgres.NewDB()
	if err != nil {
		return nil, err
	}
	taskRepository := postgres.NewTaskRepository(db)
	commentRepository := postgres.NewCommentRepository(db)
	taskService := task.NewService(taskRepository, commentRepository)
	userService := user.NewService(taskRepository)
	projectService := project.NewService(taskRepository)
	server := rest.NewServer(handlerFunc, service, taskService, userService, projectService)
	unaryServerInterceptor := middleware2.NewJWTUnaryInterceptor(publicKey)
	grpcServer := grpc.NewServer(unaryServerInterceptor, service, taskService, userService, projectService)
	app := NewApp(server, grpcServer)
	return app, nil
}

// wire.go:

type App struct {
	RestServer *rest.Server
	GrpcServer *grpc.Server
}

func NewApp(rest2 *rest.Server, grpc2 *grpc.Server) *App {
	return &App{
		RestServer: rest2,
		GrpcServer: grpc2,
	}
}
