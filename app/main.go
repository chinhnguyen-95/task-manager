package main

import _ "task-manager/docs"

// @title						Task Manager API
// @version					1.0
// @description				API for managing tasks, users, projects, comments
// @BasePath					/api/v1
// @host						localhost:8080
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	server, err := InitializeServer()

	if err != nil {
		panic(err)
	}

	server.Run(":8080")
}
