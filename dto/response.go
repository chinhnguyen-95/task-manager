package dto

// SuccessResponse represents a generic success message.
// @Description Generic success response
type SuccessResponse struct {
	Message string `json:"message" example:"action success"`
}

// ErrorResponse represents a standard error response.
// @Description Generic error response
type ErrorResponse struct {
	Error string `json:"error" example:"action error"`
}
