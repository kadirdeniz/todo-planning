package developer

import "todo-planner/internal/model"

type IService interface {
	GetAllDevelopers() ([]model.Developer, error)
}

type Service struct {
	repository IRepository
}

func NewService(repository IRepository) IService {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetAllDevelopers() ([]model.Developer, error) {
	return s.repository.GetAllDevelopers()
}

