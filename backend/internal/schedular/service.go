package schedular

import (
	"todo-planner/internal/developer"
	"todo-planner/internal/task"
)

type IService interface {
	ScheduleTasks() error
}

type Service struct {
	DeveloperService developer.IService
	TaskService task.IService
}

func NewService(developerService developer.IService, taskService task.IService) IService {
	return &Service{
		DeveloperService: developerService,
		TaskService: taskService,
	}
}

func (s *Service) ScheduleTasks() error {
	developers, err := s.DeveloperService.GetAllDevelopers()
	if err != nil {
		return err
	}

	tasks, err := s.TaskService.GetAllTasks()
	if err != nil {
		return err
	}

	// TODO: Schedule tasks to developers

	return nil
}