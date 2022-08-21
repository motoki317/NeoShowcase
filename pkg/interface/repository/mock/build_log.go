// Code generated by MockGen. DO NOT EDIT.
// Source: build_log.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/traPtitech/neoshowcase/pkg/domain"
	repository "github.com/traPtitech/neoshowcase/pkg/interface/repository"
)

// MockBuildLogRepository is a mock of BuildLogRepository interface.
type MockBuildLogRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBuildLogRepositoryMockRecorder
}

// MockBuildLogRepositoryMockRecorder is the mock recorder for MockBuildLogRepository.
type MockBuildLogRepositoryMockRecorder struct {
	mock *MockBuildLogRepository
}

// NewMockBuildLogRepository creates a new mock instance.
func NewMockBuildLogRepository(ctrl *gomock.Controller) *MockBuildLogRepository {
	mock := &MockBuildLogRepository{ctrl: ctrl}
	mock.recorder = &MockBuildLogRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildLogRepository) EXPECT() *MockBuildLogRepositoryMockRecorder {
	return m.recorder
}

// CreateBuildLog mocks base method.
func (m *MockBuildLogRepository) CreateBuildLog(ctx context.Context, branchID string) (*domain.BuildLog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBuildLog", ctx, branchID)
	ret0, _ := ret[0].(*domain.BuildLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBuildLog indicates an expected call of CreateBuildLog.
func (mr *MockBuildLogRepositoryMockRecorder) CreateBuildLog(ctx, branchID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBuildLog", reflect.TypeOf((*MockBuildLogRepository)(nil).CreateBuildLog), ctx, branchID)
}

// UpdateBuildLog mocks base method.
func (m *MockBuildLogRepository) UpdateBuildLog(ctx context.Context, args repository.UpdateBuildLogArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBuildLog", ctx, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBuildLog indicates an expected call of UpdateBuildLog.
func (mr *MockBuildLogRepositoryMockRecorder) UpdateBuildLog(ctx, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBuildLog", reflect.TypeOf((*MockBuildLogRepository)(nil).UpdateBuildLog), ctx, args)
}