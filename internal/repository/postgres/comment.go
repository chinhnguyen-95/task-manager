package postgres

import (
	"context"
	"time"

	"task-manager/internal/repository/postgres/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) Create(ctx context.Context, taskID, userID uuid.UUID, content string) error {
	comment := model.Comment{
		ID:        uuid.New(),
		TaskID:    taskID,
		UserID:    userID,
		Content:   content,
		CreatedAt: time.Now(),
	}
	return r.db.WithContext(ctx).Create(&comment).Error
}
