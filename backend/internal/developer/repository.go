package developer

import (
	"todo-planner/infrastructure"
	"todo-planner/internal/model"
)
type IRepository interface {
	GetAllDevelopers() ([]model.Developer, error)
	SaveDevelopers(developers []model.Developer) error
}

type Repository struct {
	db infrastructure.IDatabase
}

func NewRepository(db infrastructure.IDatabase) IRepository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetAllDevelopers() ([]model.Developer, error) {
	var developers []model.Developer
	err := r.db.GetDB().Find(&developers).Error
	return developers, err
}

func (r *Repository) SaveDevelopers(developers []model.Developer) error {
	return r.db.GetDB().Create(&developers).Error
}
