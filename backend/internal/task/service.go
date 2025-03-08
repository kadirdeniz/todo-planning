package task

import "todo-planner/infrastructure"

type IService interface {
	GetTasks() ([]Model, error)
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

func (s *Service) GetTasks() ([]Model, error) {
	return s.serviceCaller.GetTasks()
}
