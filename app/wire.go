//go:build wireinject
// +build wireinject

package main

import (
	"task-manager/auth"
	"task-manager/internal/grpc"
	middleware2 "task-manager/internal/grpc/middleware"
	"task-manager/internal/repository/postgres"
	"task-manager/internal/rest"
	"task-manager/internal/rest/middleware"
	"task-manager/pkg/jwtutil"
	kc "task-manager/pkg/keycloak"
	"task-manager/project"
	"task-manager/task"
	"task-manager/user"

	"github.com/google/wire"
)

type App struct {
	RestServer *rest.Server
	GrpcServer *grpc.Server
}

func NewApp(rest *rest.Server, grpc *grpc.Server) *App {
	return &App{
		RestServer: rest,
		GrpcServer: grpc,
	}
}

func InitializeApp() (*App, error) {
	wire.Build(
		jwtutil.FetchRSAPublicKeyFromJWKS,

		postgres.NewDB,
		postgres.NewTaskRepository,
		postgres.NewCommentRepository,

		kc.NewClient,

		wire.Bind(new(task.Repository), new(*postgres.TaskRepository)),
		wire.Bind(new(task.CommentRepository), new(*postgres.CommentRepository)),
		wire.Bind(new(user.TaskRepository), new(*postgres.TaskRepository)),
		wire.Bind(new(project.TaskRepository), new(*postgres.TaskRepository)),

		auth.NewService,
		task.NewService,
		user.NewService,
		project.NewService,

		middleware.JWTAuthMiddleware,

		wire.Bind(new(rest.AuthService), new(*auth.Service)),
		wire.Bind(new(rest.TaskService), new(*task.Service)),
		wire.Bind(new(rest.UserService), new(*user.Service)),
		wire.Bind(new(rest.ProjectService), new(*project.Service)),

		rest.NewServer,

		middleware2.NewJWTUnaryInterceptor,

		wire.Bind(new(grpc.AuthService), new(*auth.Service)),
		wire.Bind(new(grpc.TaskService), new(*task.Service)),
		wire.Bind(new(grpc.UserService), new(*user.Service)),
		wire.Bind(new(grpc.ProjectService), new(*project.Service)),

		grpc.NewServer,

		NewApp,
	)

	return &App{}, nil
}
