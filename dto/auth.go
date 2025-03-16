package dto

// RegisterRequest is the payload for registering a new user.
// @Description User registration request
type RegisterRequest struct {
	Name     string `json:"name" binding:"required" example:"Hucci"`
	Email    string `json:"email" binding:"required,email" example:"hucci@example.com"`
	Password string `json:"password" binding:"required,min=6" example:"demo123"`
}

// LoginRequest is the payload for logging in a user.
// @Description User login request
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"hucci@example.com"`
	Password string `json:"password" binding:"required" example:"demo123"`
}

// LoginResponse contains the JWT token returned upon successful login.
// @Description Login response containing JWT access token
type LoginResponse struct {
	AccessToken string `json:"access_token" example:"eyJhbGciOi..."`
}
