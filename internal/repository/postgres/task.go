package postgres

import (
	"context"

	"task-manager/domain"
	"task-manager/internal/repository/postgres/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(ctx context.Context, task *domain.Task) error {
	m := model.NewTaskModel(*task)
	if err := r.db.WithContext(ctx).Create(&m).Error; err != nil {
		return err
	}
	*task = m.ToDomain()
	return nil
}

func (r *TaskRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Task, error) {
	var m model.Task
	if err := r.db.WithContext(ctx).First(&m, "id = ?", id).Error; err != nil {
		return nil, err
	}
	d := m.ToDomain()
	return &d, nil
}

func (r *TaskRepository) Update(ctx context.Context, task *domain.Task) error {
	m := model.NewTaskModel(*task)
	return r.db.WithContext(ctx).Save(&m).Error
}

func (r *TaskRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&model.Task{}, "id = ?", id).Error
}

func (r *TaskRepository) Assign(ctx context.Context, taskID, userID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Model(&model.Task{}).
		Where("id = ?", taskID).
		Update("assigned_to", userID).Error
}

func (r *TaskRepository) ListByUser(ctx context.Context, userID uuid.UUID) ([]domain.Task, error) {
	var models []model.Task
	if err := r.db.WithContext(ctx).
		Where("assigned_to = ?", userID).
		Order("created_at desc").
		Find(&models).Error; err != nil {
		return nil, err
	}

	tasks := make([]domain.Task, 0, len(models))
	for _, m := range models {
		tasks = append(tasks, m.ToDomain())
	}
	return tasks, nil
}

func (r *TaskRepository) ListByProject(ctx context.Context, projectID uuid.UUID) ([]domain.Task, error) {
	var models []model.Task
	if err := r.db.WithContext(ctx).
		Where("project_id = ?", projectID).
		Order("created_at desc").
		Find(&models).Error; err != nil {
		return nil, err
	}

	tasks := make([]domain.Task, 0, len(models))
	for _, m := range models {
		tasks = append(tasks, m.ToDomain())
	}
	return tasks, nil
}
