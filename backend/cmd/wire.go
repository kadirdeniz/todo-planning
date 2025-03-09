// cmd/wire.go
//go:build wireinject
// +build wireinject

package main

import (
	"todo-planner/infrastructure"
	"todo-planner/internal/developer"
	"todo-planner/internal/schedular"
	"todo-planner/internal/task"

	"github.com/google/wire"
)


func InitializeApplication() (*Application, error) {
	wire.Build(
		infrastructure.NewConfig,
		infrastructure.NewLogger,
		infrastructure.NewDatabase,
		task.NewRepository,
		task.NewService,
		task.NewServiceCaller,
		developer.NewRepository,
		developer.NewService,
		schedular.NewService,
		schedular.NewRepository,
		wire.Struct(new(Application), "*"),
	)
	return nil, nil
}
