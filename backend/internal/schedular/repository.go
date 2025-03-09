package schedular

import (
	"todo-planner/infrastructure"
	"todo-planner/internal/model"
)

type IRepository interface {
	SaveSchedules(schedules []model.Schedule) error
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
