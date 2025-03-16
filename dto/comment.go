package dto

// CommentRequest is the payload for adding a comment to a task.
// @Description Comment creation request
type CommentRequest struct {
	Content string `json:"content" binding:"required" example:"Great job!"`
}
