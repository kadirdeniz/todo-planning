package schedular_test

import (
	"testing"
	"time"
	"todo-planner/infrastructure"
	"todo-planner/internal/model"
	"todo-planner/internal/schedular"
	"todo-planner/internal/schedular/mocks"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSchedular(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Schedular Suite")
}

var _ = Describe("SchedularRepository", func() {
	var (
		repository schedular.IRepository
		db         infrastructure.IDatabase
	)

	schedules := []model.Schedule{
		{
			StartTime: time.Now(),
			EndTime: time.Now(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Task: model.Task{
				ID:1,
				Value: 2,
				EstimatedDuration: 2,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Developer: model.Developer{
				ID: 1,
				Name: "Developer 1",
				Seniority: 1,
				WeeklyWorkHours: 40,	
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			SprintWeek: 1,
		},
		{
			StartTime: time.Now(),
			EndTime: time.Now(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Task: model.Task{
				ID:2,
				Value: 2,
				EstimatedDuration: 2,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Developer: model.Developer{
				ID: 2,
				Name: "Developer 2",
				Seniority: 1,
				WeeklyWorkHours: 40,	
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			SprintWeek: 2,
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

		Expect(db.Migrate([]interface{}{&model.Developer{},&model.Schedule{},&model.Task{}})).To(Succeed())

		repository = schedular.NewRepository(db)
	})

	Describe("SaveSchedules", func() {
		Context("When schedules are saved", func() {
			It("Should return nil", func() {
				Expect(repository.SaveSchedules(schedules)).To(Succeed())
			})
		})
	})

	Describe("GetAllSchedules", func() {
		Context("When schedules are retrieved", func() {
			It("Should return all schedules", func() {
				Expect(repository.SaveSchedules(schedules)).To(Succeed())

				retrievedSchedules, err := repository.GetAllSchedules()
				Expect(err).To(BeNil())
				Expect(retrievedSchedules).To(HaveLen(2))
			})
		})
	})
})

var _ = Describe("SchedularService", func() {
	var(
		mockRepository *mocks.MockIRepository
		service schedular.IService
	)

	schedules := []model.Schedule{
		{
			TaskID: 1,
			DeveloperID: 1,
			SprintWeek: 1,
			StartTime: time.Now(),
			EndTime: time.Now(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			TaskID: 2,
			DeveloperID: 2,
			SprintWeek: 1,
			StartTime: time.Now(),
			EndTime: time.Now(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	BeforeEach(func() {
		mockCtrl := gomock.NewController(GinkgoT())
		defer mockCtrl.Finish()

		logger, err := infrastructure.NewLogger()
		Expect(err).To(BeNil())

		mockRepository = mocks.NewMockIRepository(mockCtrl)
		service = schedular.NewService(mockRepository, logger)
	})

	Describe("GetAllSchedules", func() {
		Context("When schedules are retrieved", func() {
			It("Should return all schedules", func() {
				mockRepository.EXPECT().GetAllSchedules().Return(schedules, nil)
				result, err := service.GetAllSchedules()
				Expect(err).To(BeNil())
				Expect(result).To(HaveLen(2))
			})
		})
	})
})

var _ = Describe("ScheduleTasks", func() {
	var(
		mockRepository *mocks.MockIRepository
		service schedular.IService
	)

	tasks := []model.Task{
		{
			Value: 2,
			EstimatedDuration: 2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Value: 3,
			EstimatedDuration: 3,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	developers := []model.Developer{
		{
			Name: "Developer 1",
			Seniority: 1,
			WeeklyWorkHours: 40,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name: "Developer 2",
			Seniority: 2,
			WeeklyWorkHours: 40,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	schedules := []model.Schedule{
		{
			TaskID: 1,
			DeveloperID: 1,
			SprintWeek: 1,
			StartTime: time.Now(),
			EndTime: time.Now(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			TaskID: 2,
			DeveloperID: 2,
			SprintWeek: 1,
			StartTime: time.Now(),
			EndTime: time.Now(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	BeforeEach(func() {
		mockCtrl := gomock.NewController(GinkgoT())
		defer mockCtrl.Finish()

		logger, err := infrastructure.NewLogger()
		Expect(err).To(BeNil())

		mockRepository = mocks.NewMockIRepository(mockCtrl)
		service = schedular.NewService(mockRepository, logger)
	})

	Describe("ScheduleTasks", func() {
		Context("When tasks are scheduled", func() {
			It("Should return nil", func() {
				mockRepository.EXPECT().SaveSchedules(schedules).Return(nil)
				Expect(service.ScheduleTasks(tasks, developers)).To(Succeed())
			})
		})
	})
})