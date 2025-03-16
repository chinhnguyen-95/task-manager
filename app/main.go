package main

import (
	"log"
	"net"

	_ "task-manager/docs"
)

// @title						Task Manager API
// @version						1.0
// @description					API for managing tasks, users, projects, comments
// @BasePath					/api/v1
// @host						localhost:8080
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	app, err := InitializeApp()
	if err != nil {
		panic(err)
	}

	// Start gRPC server in a goroutine
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		log.Println("🚀 gRPC server listening on :50051")
		if err := app.GrpcServer.Serve(lis); err != nil {
			log.Fatalf("gRPC server error: %v", err)
		}
	}()

	// Start REST server (blocking)
	log.Println("🌐 REST server listening on :8080")
	app.RestServer.Run(":8080")
}
