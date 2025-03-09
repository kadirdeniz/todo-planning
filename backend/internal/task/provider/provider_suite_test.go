package provider_test

import (
	"testing"
	"todo-planner/internal/task/provider"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestProvider(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Provider Suite")
}

var _ = Describe("Provider", func() {
	Describe("Provider", func() {
		Context("GetTasks", func() {
			provider := provider.NewProvider("https://raw.githubusercontent.com/WEG-Technology/mock/refs/heads/main/mock-one", provider.NewProvider1Mapper())
			It("should return tasks", func() {
				tasks, err := provider.GetTasks()
				Expect(err).To(BeNil())
				Expect(tasks).To(HaveLen(8))
			})
		})
	})
})

var _ = Describe("Mapper", func() {
	Context("Provider1Mapper", func() {
		mapper := provider.NewProvider1Mapper()
		mockData := `[
			{
				"id": 1,
				"value": 3,
				"estimated_duration": 4
			},
			{
				"id": 2,
				"value": 6,
				"estimated_duration": 12
			},
			{
				"id": 3,
				"value": 5,
				"estimated_duration": 9
			},
			{
				"id": 4,
				"value": 5,
				"estimated_duration": 5
			},
			{
				"id": 5,
				"value": 7,
				"estimated_duration": 7
			},
			{
				"id": 6,
				"value": 3,
				"estimated_duration": 5
			},
			{
				"id": 7,
				"value": 4,
				"estimated_duration": 8
			},
			{
				"id": 8,
				"value": 6,
				"estimated_duration": 3
			}
		]`
		It("should return tasks", func() {
			tasks, err := mapper.Map([]byte(mockData))
			Expect(err).To(BeNil())
			Expect(tasks).To(HaveLen(8))
		})
	})
	Context("Provider2Mapper", func() {
		mapper := provider.NewProvider2Mapper()
		mockData := `[
			{
				"id": 1,
				"zorluk": 3,
				"sure": 5
			},
			{
				"id": 2,
				"zorluk": 2,
				"sure": 3
			},
			{
				"id": 3,
				"zorluk": 1,
				"sure": 2
			},
			{
				"id": 4,
				"zorluk": 4,
				"sure": 7
			},
			{
				"id": 5,
				"zorluk": 5,
				"sure": 8
			},
			{
				"id": 6,
				"zorluk": 2,
				"sure": 4
			},
			{
				"id": 7,
				"zorluk": 3,
				"sure": 6
			},
			{
				"id": 8,
				"zorluk": 1,
				"sure": 3
			}
		]`
		It("should return tasks", func() {
			tasks, err := mapper.Map([]byte(mockData))
			Expect(err).To(BeNil())
			Expect(tasks).To(HaveLen(8))
		})
	})
})