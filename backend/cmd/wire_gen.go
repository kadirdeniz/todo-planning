// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"todo-planner/infrastructure"
	"todo-planner/internal/developer"
	"todo-planner/internal/schedular"
	"todo-planner/internal/task"
)

// Injectors from wire.go:

func InitializeApplication() (*Application, error) {
	config, err := infrastructure.NewConfig()
	if err != nil {
		return nil, err
	}
	iLogger, err := infrastructure.NewLogger()
	if err != nil {
		return nil, err
	}
	iDatabase := infrastructure.NewDatabase(config)
	iRepository := task.NewRepository(iDatabase, iLogger)
	iServiceCaller := task.NewServiceCaller(config, iLogger)
	iService := task.NewService(iRepository, iServiceCaller, iDatabase, iLogger)
	developerIRepository := developer.NewRepository(iDatabase)
	developerIService := developer.NewService(developerIRepository)
	schedularIRepository := schedular.NewRepository(iDatabase)
	schedularIService := schedular.NewService(schedularIRepository, iLogger, iService, developerIService)
	application := &Application{
		Config:           config,
		Logger:           iLogger,
		Database:         iDatabase,
		TaskService:      iService,
		DeveloperService: developerIService,
		SchedularService: schedularIService,
	}
	return application, nil
}
