// cmd/wire.go
//go:build wireinject
// +build wireinject

package main

import (
	"todo-planner/infrastructure"
	"todo-planner/internal/task"
	servicecaller "todo-planner/internal/task/service-caller"

	"github.com/google/wire"
)

func provideTaskProviders(config infrastructure.Config, logger infrastructure.ILogger) []task.IServiceCaller {
	return []task.IServiceCaller{
		servicecaller.NewProvider1(config.Provider1.URL, logger),
		servicecaller.NewProvider2(config.Provider2.URL, logger),
	}
}

func InitializeApplication() (*Application, error) {
	wire.Build(
		infrastructure.NewConfig,
		infrastructure.NewLogger,
		infrastructure.NewDatabase,
		task.NewRepository,
		task.NewService,
		task.NewServiceCaller,
		provideTaskProviders,
		wire.Struct(new(Application), "*"),
	)
	return nil, nil
}
