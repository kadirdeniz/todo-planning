package task

import (
	"todo-planner/infrastructure"
	"todo-planner/internal/model"
)

type IRepository interface {
	SaveTasks(tasks []model.Task) error
	GetAllTasks() ([]model.Task, error)
}

type Repository struct {
	db infrastructure.IDatabase
	logger infrastructure.ILogger
}

func NewRepository(db infrastructure.IDatabase, logger infrastructure.ILogger) IRepository {
	return &Repository{
		db: db,
		logger: logger,
	}
}

func (r *Repository) SaveTasks(tasks []model.Task) error {
	return r.db.GetDB().Create(tasks).Error
}

func (r *Repository) GetAllTasks() ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.GetDB().Find(&tasks).Error
	return tasks, err
}
