package provider

import (
	"encoding/json"
	"todo-planner/internal/model"
)

type IMapper interface {
	Map(data []byte) ([]model.Task, error)
}

type Provider1Mapper struct {}

func NewProvider1Mapper() IMapper {
	return &Provider1Mapper{}
}

func (m *Provider1Mapper) Map(data []byte) ([]model.Task, error) {
	var response []Provider1Response

	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	var tasks []model.Task

	for _, task := range response {
		tasks = append(tasks, model.Task{
			Value: task.Value,
			EstimatedDuration: task.EstimatedDuration,
		})
	}

	return tasks, nil
}

type Provider2Mapper struct {}

func NewProvider2Mapper() IMapper {
	return &Provider2Mapper{}
}

func (m *Provider2Mapper) Map(data []byte) ([]model.Task, error) {
	var response []Provider2Response

	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	var tasks []model.Task

	for _, task := range response {
		tasks = append(tasks, model.Task{
			Value: task.Value,
			EstimatedDuration: task.EstimatedDuration,
		})
	}

	return tasks, nil
}
