package rest

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(
	jwtMiddleware gin.HandlerFunc,
	authSvc AuthService,
	taskSvc TaskService,
	userSvc UserService,
	projectSvc ProjectService,
) *Server {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")

	RegisterAuthRoutes(api, authSvc)

	taskGroup := api.Group("/tasks", jwtMiddleware)
	RegisterTaskRoutes(taskGroup, taskSvc)

	userGroup := api.Group("/users", jwtMiddleware)
	RegisterUserRoutes(userGroup, userSvc)

	projectGroup := api.Group("/projects", jwtMiddleware)
	RegisterProjectRoutes(projectGroup, projectSvc)

	return &Server{engine: r}
}

func (s *Server) Run(addr string) {
	log.Printf("Starting server at %s", addr)
	_ = s.engine.Run(addr)
}
