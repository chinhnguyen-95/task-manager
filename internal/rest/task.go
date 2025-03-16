package rest

import (
	"context"
	"net/http"
	"task-manager/internal/rest/middleware"
	"time"

	"task-manager/domain"
	"task-manager/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskService interface {
	Create(ctx context.Context, task *domain.Task) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Task, error)
	Update(ctx context.Context, task *domain.Task) error
	Delete(ctx context.Context, id uuid.UUID) error
	Assign(ctx context.Context, taskID uuid.UUID, userID uuid.UUID) error
	Comment(ctx context.Context, taskID, userID uuid.UUID, content string) error
}

// RegisterTaskRoutes registers task routes to the router group.
func RegisterTaskRoutes(rg *gin.RouterGroup, service TaskService) {
	rg.POST("/", createTaskHandler(service))
	rg.GET("/:id", getTaskHandler(service))
	rg.PUT("/:id", updateTaskHandler(service))
	rg.DELETE("/:id", deleteTaskHandler(service))
	rg.PUT("/:id/assign", assignTaskHandler(service))
	rg.PUT("/:id/comment", commentOnTaskHandler(service))
}

// createTaskHandler handles creating a new task
//
//	@Summary	Create a new task
//	@Tags		Tasks
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.CreateTaskRequest	true	"Task data"
//	@Success	200		{object}	dto.TaskResponse
//	@Failure	400		{object}	dto.ErrorResponse
//	@Failure	500		{object}	dto.ErrorResponse
//	@Router		/tasks [post]
//	@Security	BearerAuth
func createTaskHandler(service TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.CreateTaskRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
			return
		}

		userIDStr, _ := c.Get(middleware.UserIDKey)
		userID, err := uuid.Parse(userIDStr.(string))
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid user ID"})
			return
		}

		t := &domain.Task{
			ID:          uuid.New(),
			Title:       req.Title,
			Description: req.Description,
			Status:      req.Status,
			ProjectID:   req.ProjectID,
			AssignedTo:  userID,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		if err := service.Create(c, t); err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, dto.NewTaskResponse(*t))
	}
}

// getTaskHandler retrieves a task by ID
//
//	@Summary	Get a task by ID
//	@Tags		Tasks
//	@Produce	json
//	@Param		id	path		string	true	"Task ID"
//	@Success	200	{object}	dto.TaskResponse
//	@Failure	400	{object}	dto.ErrorResponse
//	@Failure	404	{object}	dto.ErrorResponse
//	@Router		/tasks/{id} [get]
//	@Security	BearerAuth
func getTaskHandler(service TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid ID"})
			return
		}

		t, err := service.GetByID(c, id)
		if err != nil {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "task not found"})
			return
		}

		c.JSON(http.StatusOK, dto.NewTaskResponse(*t))
	}
}

// updateTaskHandler updates a task
//
//	@Summary	Update a task by ID
//	@Tags		Tasks
//	@Accept		json
//	@Produce	json
//	@Param		id		path		string					true	"Task ID"
//	@Param		request	body		dto.UpdateTaskRequest	true	"Updated task data"
//	@Success	200		{object}	dto.TaskResponse
//	@Failure	400		{object}	dto.ErrorResponse
//	@Failure	500		{object}	dto.ErrorResponse
//	@Router		/tasks/{id} [put]
//	@Security	BearerAuth
func updateTaskHandler(service TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid ID"})
			return
		}

		var req dto.UpdateTaskRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
			return
		}

		existing, err := service.GetByID(c, id)
		if err != nil {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "task not found"})
			return
		}

		if req.Title != nil {
			existing.Title = *req.Title
		}
		if req.Description != nil {
			existing.Description = *req.Description
		}
		if req.Status != nil {
			existing.Status = *req.Status
		}

		if err := service.Update(c, existing); err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, dto.NewTaskResponse(*existing))
	}
}

// deleteTaskHandler deletes a task
//
//	@Summary	Delete a task by ID
//	@Tags		Tasks
//	@Produce	json
//	@Param		id	path		string	true	"Task ID"
//	@Success	200	{object}	dto.SuccessResponse
//	@Failure	400	{object}	dto.ErrorResponse
//	@Failure	500	{object}	dto.ErrorResponse
//	@Router		/tasks/{id} [delete]
//	@Security	BearerAuth
func deleteTaskHandler(service TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid ID"})
			return
		}

		if err := service.Delete(c, id); err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, dto.SuccessResponse{Message: "task deleted"})
	}
}

// assignTaskHandler assigns a task to a user
//
//	@Summary	Assign task to user
//	@Tags		Tasks
//	@Accept		json
//	@Produce	json
//	@Param		id		path		string				true	"Task ID"
//	@Param		request	body		dto.AssignRequest	true	"User to assign"
//	@Success	200		{object}	dto.SuccessResponse
//	@Failure	400		{object}	dto.ErrorResponse
//	@Failure	500		{object}	dto.ErrorResponse
//	@Router		/tasks/{id}/assign [put]
//	@Security	BearerAuth
func assignTaskHandler(service TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		taskID, err := uuid.Parse(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid task ID"})
			return
		}

		var req dto.AssignRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
			return
		}

		if err := service.Assign(c, taskID, req.UserID); err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, dto.SuccessResponse{Message: "task assigned"})
	}
}

// commentOnTaskHandler comments on task
//
// @Summary		Comment on a task
// @Tags		Tasks
// @Accept		json
// @Produce		json
// @Param		id		path		string				true	"Task ID"
// @Param		request	body		dto.CommentRequest	true	"Comment content"
// @Success		200		{object}	dto.SuccessResponse
// @Failure		400		{object}	dto.ErrorResponse
// @Failure		500		{object}	dto.ErrorResponse
// @Router		/tasks/{id}/comment [put]
// @Security	BearerAuth
func commentOnTaskHandler(service TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {
		taskID, _ := uuid.Parse(c.Param("id"))
		userIDStr, _ := c.Get(middleware.UserIDKey)
		userID, _ := uuid.Parse(userIDStr.(string))

		var req dto.CommentRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
			return
		}

		if err := service.Comment(c, taskID, userID, req.Content); err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, dto.SuccessResponse{Message: "comment added"})
	}
}
