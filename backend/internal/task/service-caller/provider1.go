package servicecaller

import (
	"todo-planner/infrastructure"
	"todo-planner/internal/task"
)

type Provider1 struct {
	URL string
	logger infrastructure.ILogger
}

func NewProvider1(url string, logger infrastructure.ILogger) *Provider1 {
	return &Provider1{
		URL: url,
		logger: logger,
	}
}

func (p *Provider1) GetTasks() ([]task.Model, error) {
	return nil, nil
}
