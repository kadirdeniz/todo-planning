package developer_test

import (
	"testing"
	"time"
	"todo-planner/infrastructure"
	"todo-planner/internal/developer"
	"todo-planner/internal/developer/mocks"
	"todo-planner/internal/model"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDeveloperRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Developer Repository Suite")
}

var _ = Describe("DeveloperRepository", func() {
	var (
		repository developer.IRepository
		db         infrastructure.IDatabase
	)

	BeforeEach(func() {
		config := infrastructure.Config{
			Database: infrastructure.Database{
				URL: "postgres://test:test@localhost:5433/todo_planning_test?sslmode=disable",
			},
		}
		
		db = infrastructure.NewDatabase(config)
		Expect(db.Connect()).To(Succeed())

		Expect(db.GetDB().Migrator().DropTable(&model.Developer{},&model.Schedule{},&model.Task{},)).To(Succeed())

		Expect(db.Migrate([]interface{}{&model.Developer{}})).To(Succeed())
		
		repository = developer.NewRepository(db)
	})

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
		{
			Name: "Developer 3",
			Seniority: 3,
			WeeklyWorkHours: 40,
			CreatedAt: time.Now(),				
			UpdatedAt: time.Now(),
		},
		{
			Name: "Developer 4",
			Seniority: 4,
			WeeklyWorkHours: 40,
			CreatedAt: time.Now(),				
			UpdatedAt: time.Now(),
		},
		{
			Name: "Developer 5",
			Seniority: 5,
			WeeklyWorkHours: 40,
			CreatedAt: time.Now(),				
			UpdatedAt: time.Now(),
		},
	}

	Describe("SaveDevelopers", func() {
		It("should developers", func() {
			err := repository.SaveDevelopers(developers)
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("GetAllDevelopers", func() {
		It("should return developers", func() {
			err := repository.SaveDevelopers(developers)
			Expect(err).NotTo(HaveOccurred())

			result, err := repository.GetAllDevelopers()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(HaveLen(5))
		})
	})
})

var _ = Describe("DeveloperService", func() {
	var (
		service developer.IService
		repository *mocks.MockIRepository
	)

	BeforeEach(func() {
		mockCtrl := gomock.NewController(GinkgoT())
		defer mockCtrl.Finish()
		repository = mocks.NewMockIRepository(mockCtrl)
		service = developer.NewService(repository)
	})

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
		{
			Name: "Developer 3",
			Seniority: 3,
			WeeklyWorkHours: 40,
			CreatedAt: time.Now(),				
			UpdatedAt: time.Now(),
		},
		{
			Name: "Developer 4",
			Seniority: 4,
			WeeklyWorkHours: 40,
			CreatedAt: time.Now(),				
			UpdatedAt: time.Now(),
		},
		{
			Name: "Developer 5",
			Seniority: 5,
			WeeklyWorkHours: 40,
			CreatedAt: time.Now(),				
			UpdatedAt: time.Now(),
		},
	}

	Describe("SaveDevelopers", func() {
		It("should developers", func() {
			repository.EXPECT().SaveDevelopers(developers).Return(nil)
			err := service.SaveDevelopers(developers)
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("GetAllDevelopers", func() {
		It("should return developers", func() {
			repository.EXPECT().GetAllDevelopers().Return(developers, nil)
			result, err := service.GetAllDevelopers()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(HaveLen(5))
		})
	})
})