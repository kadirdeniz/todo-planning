package main

import (
	"todo-planner/infrastructure"
	"todo-planner/internal/developer"
	"todo-planner/internal/model"
	"todo-planner/internal/schedular"
	"todo-planner/internal/task"

	"go.uber.org/zap"
)

type Application struct {
    Config      infrastructure.Config
    Logger      infrastructure.ILogger
    Database    infrastructure.IDatabase
    TaskService task.IService
    DeveloperService developer.IService
    SchedularService schedular.IService
}

func (a *Application) Run() error {
    if err := a.Database.Connect(); err != nil {
        a.Logger.Error("Failed to connect to database", zap.Error(err))
        return err
    }

    if err := a.Database.Migrate([]interface{}{
        &model.Developer{},
        &model.Task{},
    }); err != nil {
        a.Logger.Error("Failed to migrate database", zap.Error(err))
        return err
    }

    tasks, err := a.TaskService.GetTasksFromProviders()
    if err != nil {
        a.Logger.Error("Failed to get tasks from providers", zap.Error(err))
        return err
    }

    err = a.TaskService.SaveTasks(tasks)
    if err != nil {
        a.Logger.Error("Failed to save tasks", zap.Error(err))
        return err
    }

    err = a.DeveloperService.SaveDevelopers([]model.Developer{
        {
            Name: "John Doe",
            Seniority: 1,
        },
        {
            Name: "Jane Doe",
            Seniority: 2,
        },
        {
            Name: "Jim Doe",
            Seniority: 3,
        },
        {
            Name: "Jill Doe",
            Seniority: 4,
        },
        {
            Name: "Jack Doe",
            Seniority: 5,
        },
    })
    if err != nil {
        a.Logger.Error("Failed to save developers", zap.Error(err))
        return err
    }

    a.Logger.Info("Application started")
    return nil
} 