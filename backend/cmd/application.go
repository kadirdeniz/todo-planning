package main

import (
	"todo-planner/infrastructure"
	"todo-planner/internal/task"

	"go.uber.org/zap"
)

type Application struct {
    Config      infrastructure.Config
    Logger      infrastructure.ILogger
    Database    infrastructure.IDatabase
    TaskService task.IService
}

func (a *Application) Run() error {
    if err := a.Database.Connect(); err != nil {
        a.Logger.Error("Failed to connect to database", zap.Error(err))
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

    a.Logger.Info("Application started")
    return nil
} 