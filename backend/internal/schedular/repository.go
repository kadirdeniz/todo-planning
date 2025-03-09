package schedular

import (
	"todo-planner/infrastructure"
	"todo-planner/internal/model"
)

type IRepository interface {
	SaveSchedules(schedules []model.Schedule) error
	GetAllSchedules() ([]model.Schedule, error)	
}

type Repository struct {
	db infrastructure.IDatabase
}

func NewRepository(db infrastructure.IDatabase) IRepository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) SaveSchedules(schedules []model.Schedule) error {
	return r.db.GetDB().Create(&schedules).Error
}

func (r *Repository) GetAllSchedules() ([]model.Schedule, error) {
	var schedules []model.Schedule
	err := r.db.GetDB().Preload("Task").Preload("Developer").Find(&schedules).Error
	if err != nil {
		return nil, err
	}
	return schedules, nil
}
