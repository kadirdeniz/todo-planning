// Code generated by MockGen. DO NOT EDIT.
// Source: internal/schedular/repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	model "todo-planner/internal/model"

	gomock "github.com/golang/mock/gomock"
)

// MockIRepository is a mock of IRepository interface.
type MockIRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRepositoryMockRecorder
}

// MockIRepositoryMockRecorder is the mock recorder for MockIRepository.
type MockIRepositoryMockRecorder struct {
	mock *MockIRepository
}

// NewMockIRepository creates a new mock instance.
func NewMockIRepository(ctrl *gomock.Controller) *MockIRepository {
	mock := &MockIRepository{ctrl: ctrl}
	mock.recorder = &MockIRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRepository) EXPECT() *MockIRepositoryMockRecorder {
	return m.recorder
}

// GetAllSchedules mocks base method.
func (m *MockIRepository) GetAllSchedules() ([]model.Schedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSchedules")
	ret0, _ := ret[0].([]model.Schedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSchedules indicates an expected call of GetAllSchedules.
func (mr *MockIRepositoryMockRecorder) GetAllSchedules() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSchedules", reflect.TypeOf((*MockIRepository)(nil).GetAllSchedules))
}

// SaveSchedules mocks base method.
func (m *MockIRepository) SaveSchedules(schedules []model.Schedule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveSchedules", schedules)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveSchedules indicates an expected call of SaveSchedules.
func (mr *MockIRepositoryMockRecorder) SaveSchedules(schedules interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveSchedules", reflect.TypeOf((*MockIRepository)(nil).SaveSchedules), schedules)
}
