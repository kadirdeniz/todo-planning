package task

import "todo-planner/infrastructure"

type IServiceCaller interface {
	GetTasks() ([]Model, error)
}

type serviceCaller struct {
	providers []IServiceCaller
}

func NewServiceCaller(config infrastructure.Config, providers []IServiceCaller, logger infrastructure.ILogger) IServiceCaller {
	return &serviceCaller{
		providers: providers,
	}
}

func (s *serviceCaller) GetTasks() ([]Model, error) {
	return []Model{}, nil
}