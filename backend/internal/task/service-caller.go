package task

import (
	"todo-planner/infrastructure"
	"todo-planner/internal/model"
	"todo-planner/internal/task/provider"
)

type IServiceCaller interface {
	GetTasks() ([]model.Task, error)
}

type serviceCaller struct {
	providers []IServiceCaller
}

func NewServiceCaller(config infrastructure.Config, logger infrastructure.ILogger) IServiceCaller {
	providers := []IServiceCaller{}

	for _, pr := range config.Providers {
		if pr.Type == "1" {
			providers = append(providers, provider.NewProvider(pr.URL, provider.NewProvider1Mapper()))
		} else if pr.Type == "2" {
			providers = append(providers, provider.NewProvider(pr.URL, provider.NewProvider2Mapper()))
		}
	}

	return &serviceCaller{
		providers: providers,
	}
}

func (s *serviceCaller) GetTasks() ([]model.Task, error) {
	// merge tasks from all providers
	tasks := []model.Task{}
	for _, provider := range s.providers {
		providerTasks, err := provider.GetTasks()
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, providerTasks...)
	}

	return tasks, nil
}
