package rest

import (
	"context"
	"net/http"

	"task-manager/domain"
	"task-manager/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProjectService interface {
	ListTasks(ctx context.Context, projectID uuid.UUID) ([]domain.Task, error)
}

func RegisterProjectRoutes(rg *gin.RouterGroup, service ProjectService) {
	rg.GET("/:project_id/tasks", ListTasksByProjectHandler(service))
}

// ListTasksByProjectHandler handles GET /projects/:project_id/tasks
//
//	@Summary		List tasks by project
//	@Description	Get all tasks belonging to a specific project
//	@Tags			Projects
//	@Produce		json
//	@Param			project_id	path		string	true	"Project ID"
//	@Success		200			{array}		dto.TaskResponse
//	@Failure		400			{object}	dto.ErrorResponse
//	@Failure		500			{object}	dto.ErrorResponse
//	@Router			/projects/{project_id}/tasks [get]
//	@Security		BearerAuth
func ListTasksByProjectHandler(service ProjectService) gin.HandlerFunc {
	return func(c *gin.Context) {
		projectID, err := uuid.Parse(c.Param("project_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid project ID"})
			return
		}

		tasks, err := service.ListTasks(c, projectID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, dto.NewTaskResponseList(tasks))
	}
}
