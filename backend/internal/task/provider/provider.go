package provider

import (
	"io"
	"net/http"
	"todo-planner/internal/model"
)

type IProvider interface {
	GetTasks() ([]model.Task, error)
}


type Provider struct {
	URL string
	Mapper IMapper
}

func NewProvider(url string, mapper IMapper) IProvider {
	return &Provider{URL: url, Mapper: mapper}
}

func (p *Provider) GetTasks() ([]model.Task, error) {
	resp, err := http.Get(p.URL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return p.Mapper.Map(data)
}