package schedular

import (
	"math"
	"sort"
	"time"
	"todo-planner/infrastructure"
	"todo-planner/internal/model"
)

type IService interface {
	ScheduleTasks(tasks []model.Task, developers []model.Developer) error
	GetAllSchedules() ([]model.Schedule, error)
}

type Service struct {
	Repository IRepository
	Logger infrastructure.ILogger
}

func NewService(repository IRepository, logger infrastructure.ILogger) IService {
	return &Service{
		Repository: repository,
		Logger: logger,
	}
}

func (s *Service) ScheduleTasks(tasks []model.Task, developers []model.Developer) error {
	sort.Slice(developers, func(i, j int) bool {
		return developers[i].Seniority > developers[j].Seniority
	})

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].EstimatedDuration.Hours() > tasks[j].EstimatedDuration.Hours()
	})

	schedules := []model.Schedule{}
	sprintWeeks := s.CalculateRequiredWeeks(tasks, developers)

	// o(n^3) :(
	for week := 1; week <= sprintWeeks; week++ {
		for _, developer := range developers {
			developerOverload := 0
			for _, task := range tasks {
				if developerOverload + int(task.EstimatedDuration.Hours()) > developer.WeeklyWorkHours {
					continue
				}
				
				schedules = append(schedules, model.Schedule{
					Task: task,
					Developer: developer,
					SprintWeek: week,
					StartTime: time.Now().AddDate(0, 0, week * 7),	
					EndTime: time.Now().AddDate(0, 0, week * 7).Add(time.Duration(task.EstimatedDuration.Hours()) * time.Hour),
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				})
				developerOverload += int(task.EstimatedDuration.Hours())
				tasks = tasks[1:]
			}
		}
	}

	err := s.Repository.SaveSchedules(schedules)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) CalculateRequiredWeeks(tasks []model.Task, developers []model.Developer) int {
	var totalWorkload float64
    for _, task := range tasks {
        totalWorkload += float64(task.EstimatedDuration.Hours())
    }

    var weeklyTeamCapacity float64
    for _, dev := range developers {
        weeklyTeamCapacity += float64(dev.WeeklyWorkHours)
    }

    requiredWeeks := math.Ceil(totalWorkload / weeklyTeamCapacity)
	return int(requiredWeeks)
}

func (s *Service) GetAllSchedules() ([]model.Schedule, error) {
	return s.Repository.GetAllSchedules()
}
