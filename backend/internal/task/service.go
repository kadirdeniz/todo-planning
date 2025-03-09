package task

import (
	"todo-planner/infrastructure"
	"todo-planner/internal/model"
)

type IService interface {
	GetTasksFromProviders() ([]model.Task, error)
	SaveTasks(tasks []model.Task) error
	GetAllTasks() ([]model.Task, error)
}

type Service struct {
	repository IRepository
	serviceCaller IServiceCaller
	config infrastructure.IDatabase
	logger infrastructure.ILogger
}

func NewService(repository IRepository, serviceCaller IServiceCaller, config infrastructure.IDatabase, logger infrastructure.ILogger) IService {
	return &Service{
		repository: repository,
		serviceCaller: serviceCaller,
		config: config,
		logger: logger,
	}
}

func (s *Service) GetTasksFromProviders() ([]model.Task, error) {
	return s.serviceCaller.GetTasks()
}

func (s *Service) SaveTasks(tasks []model.Task) error {
	return s.repository.SaveTasks(tasks)
}

func (s *Service) GetAllTasks() ([]model.Task, error) {
	return s.repository.GetAllTasks()
}
