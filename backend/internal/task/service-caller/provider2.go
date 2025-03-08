package servicecaller

import (
	"todo-planner/infrastructure"
	"todo-planner/internal/task"
)

type Provider2 struct {
	URL string
	logger infrastructure.ILogger
}

func NewProvider2(url string, logger infrastructure.ILogger) *Provider2 {
	return &Provider2{
		URL: url,
		logger: logger,
	}
}

func (p *Provider2) GetTasks() ([]task.Model, error) {
	return nil, nil
}
