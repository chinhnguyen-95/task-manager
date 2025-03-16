package rest

import (
	"context"
	"net/http"

	"task-manager/domain"
	"task-manager/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserService interface {
	ListTasks(ctx context.Context, userID uuid.UUID) ([]domain.Task, error)
}

func RegisterUserRoutes(rg *gin.RouterGroup, service UserService) {
	rg.GET("/:user_id/tasks", ListTasksByUserHandler(service))
}

// ListTasksByUserHandler handles GET /users/:user_id/tasks
//
//	@Summary		List tasks by user
//	@Description	Get all tasks assigned to a specific user
//	@Tags			Users
//	@Produce		json
//	@Param			user_id	path		string	true	"User ID"
//	@Success		200		{array}		dto.TaskResponse
//	@Failure		400		{object}	dto.ErrorResponse
//	@Failure		500		{object}	dto.ErrorResponse
//	@Router			/users/{user_id}/tasks [get]
//	@Security		BearerAuth
func ListTasksByUserHandler(service UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := uuid.Parse(c.Param("user_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid user ID"})
			return
		}

		tasks, err := service.ListTasks(c, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, dto.NewTaskResponseList(tasks))
	}
}
