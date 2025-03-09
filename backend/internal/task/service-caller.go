package task

import (
	"todo-planner/infrastructure"
	"todo-planner/internal/model"
	"todo-planner/internal/task/provider"

	"go.uber.org/zap"
)

type IServiceCaller interface {
	GetTasks() ([]model.Task, error)
}

type serviceCaller struct {
	providers []IServiceCaller
	logger    infrastructure.ILogger
}

type result struct {
	tasks []model.Task
	err   error
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
		logger:    logger,
	}
}

func (s *serviceCaller) GetTasks() ([]model.Task, error) {
	resultChan := make(chan result, len(s.providers))

	for _, provider := range s.providers {
		go func(p IServiceCaller) {
			tasks, err := p.GetTasks()
			resultChan <- result{tasks: tasks, err: err}
		}(provider)
	}

	var allTasks []model.Task
	for i := 0; i < len(s.providers); i++ {
		r := <-resultChan
		if r.err != nil {
			s.logger.Error("Failed to get tasks from provider", zap.Error(r.err))
		}
		allTasks = append(allTasks, r.tasks...)
	}

	return allTasks, nil
}
