package rest

import (
	"context"
	"net/http"

	"task-manager/dto"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Register(ctx context.Context, req dto.RegisterRequest) error
	Login(ctx context.Context, req dto.LoginRequest) (string, error)
}

// RegisterAuthRoutes registers authentication-related routes (public, no middleware).
func RegisterAuthRoutes(rg *gin.RouterGroup, service AuthService) {
	rg.POST("/register", registerHandler(service))
	rg.POST("/login", loginHandler(service))
}

// registerHandler handles user registration via Keycloak Admin API
//
//	@Summary	Register a new user
//	@Tags		Auth
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.RegisterRequest	true	"User registration info"
//	@Success	200		{object}	dto.SuccessResponse
//	@Failure	400		{object}	dto.ErrorResponse
//	@Failure	500		{object}	dto.ErrorResponse
//	@Router		/register [post]
func registerHandler(service AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
			return
		}

		if err := service.Register(c, req); err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, dto.SuccessResponse{Message: "registered successfully"})
	}
}

// loginHandler handles user login and returns a JWT access token
//
//	@Summary	Log in and get JWT token
//	@Tags		Auth
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.LoginRequest	true	"User credentials"
//	@Success	200		{object}	dto.LoginResponse
//	@Failure	400		{object}	dto.ErrorResponse
//	@Failure	401		{object}	dto.ErrorResponse
//	@Router		/login [post]
func loginHandler(service AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
			return
		}

		token, err := service.Login(c, req)
		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, dto.LoginResponse{AccessToken: token})
	}
}
