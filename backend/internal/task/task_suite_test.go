package task_test

import (
	"testing"
	"time"
	"todo-planner/infrastructure"
	"todo-planner/internal/model"
	"todo-planner/internal/task"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTask(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Task Suite")
}

var _ = Describe("TaskRepository", func() {
	var (
		repository task.IRepository
		db         infrastructure.IDatabase
	)

	tasks := []model.Task{
		{
			Value: 2,
			EstimatedDuration: 2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Value: 2,
			EstimatedDuration: 2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Value: 2,
			EstimatedDuration: 2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	BeforeEach(func() {
		config := infrastructure.Config{
			Database: infrastructure.Database{
				URL: "postgres://test:test@localhost:5433/todo_planning_test?sslmode=disable",
			},
		}
		
		db = infrastructure.NewDatabase(config)
		Expect(db.Connect()).To(Succeed())

		Expect(db.GetDB().Migrator().DropTable(&model.Developer{},&model.Schedule{},&model.Task{},)).To(Succeed())

		Expect(db.Migrate([]interface{}{&model.Task{},})).To(Succeed())
		
		logger, err := infrastructure.NewLogger()
		Expect(err).To(BeNil())

		repository = task.NewRepository(db, logger)
	})


	Describe("SaveTasks", func() {
		Context("When tasks are saved", func() {
			It("Should return nil", func() {
				Expect(repository.SaveTasks(tasks)).To(Succeed())
			})
		})
	})

	Describe("GetAllTasks", func() {
		Context("When tasks are retrieved", func() {
			It("Should return all tasks", func() {
				Expect(repository.SaveTasks(tasks)).To(Succeed())
				retrievedTasks, err := repository.GetAllTasks()
				Expect(err).To(BeNil())
				Expect(retrievedTasks).To(HaveLen(3))
			})
		})
	})
})

var _ = Describe("TaskService", func() {
	var (
		service task.IService
		db         infrastructure.IDatabase
	)

	tasks := []model.Task{
		{
			Value: 2,
			EstimatedDuration: 2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Value: 2,
			EstimatedDuration: 2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Value: 2,
			EstimatedDuration: 2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	BeforeEach(func() {
		config := infrastructure.Config{
			Database: infrastructure.Database{
				URL: "postgres://test:test@localhost:5433/todo_planning_test?sslmode=disable",
			},
		}
		
		db = infrastructure.NewDatabase(config)
		Expect(db.Connect()).To(Succeed())

		Expect(db.GetDB().Migrator().DropTable(&model.Developer{},&model.Schedule{},&model.Task{},)).To(Succeed())

		Expect(db.Migrate([]interface{}{&model.Task{}})).To(Succeed())
		
		logger, err := infrastructure.NewLogger()
		Expect(err).To(BeNil())

		repository := task.NewRepository(db, logger)
		service = task.NewService(repository, task.NewServiceCaller(config, logger), &config.Database, logger)
	})

	Describe("SaveTasks", func() {
		Context("When tasks are saved", func() {
			It("Should return nil", func() {
				Expect(service.SaveTasks(tasks)).To(Succeed())
			})
		})
	})

	Describe("GetAllTasks", func() {
		Context("When tasks are retrieved", func() {
			It("Should return all tasks", func() {
				Expect(service.SaveTasks(tasks)).To(Succeed())
				retrievedTasks, err := service.GetAllTasks()
				Expect(err).To(BeNil())
				Expect(retrievedTasks).To(HaveLen(3))
			})
		})
	})

	Describe("GetTasksFromProviders", func() {
		Context("When tasks are retrieved", func() {
			It("Should return all tasks", func() {
				tasks, err := service.GetTasksFromProviders()
				Expect(err).To(BeNil())
				Expect(tasks).To(HaveLen(16))
			})
		})
	})
})

var _ = Describe("TaskServiceCaller", func() {
	var (
		serviceCaller task.IServiceCaller
	)

	BeforeEach(func() {
		config := infrastructure.Config{
			Providers: []infrastructure.Provider{
				{
					Type: "1",
					URL: "https://raw.githubusercontent.com/WEG-Technology/mock/refs/heads/main/mock-one",
				},
				{
					Type: "2",
					URL: "https://raw.githubusercontent.com/WEG-Technology/mock/refs/heads/main/mock-two",
				},
			},
		}
		
		logger, err := infrastructure.NewLogger()
		Expect(err).To(BeNil())

		serviceCaller = task.NewServiceCaller(config, logger)
	})

	Describe("GetTasks", func() {
		Context("When tasks are retrieved", func() {
			It("Should return all tasks", func() {
				tasks, err := serviceCaller.GetTasks()
				Expect(err).To(BeNil())
				Expect(tasks).To(HaveLen(16))
			})
		})
	})
})