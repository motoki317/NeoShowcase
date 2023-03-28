// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/traPtitech/neoshowcase/pkg/domain"
)

// MockApplicationRepository is a mock of ApplicationRepository interface.
type MockApplicationRepository struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationRepositoryMockRecorder
}

// MockApplicationRepositoryMockRecorder is the mock recorder for MockApplicationRepository.
type MockApplicationRepositoryMockRecorder struct {
	mock *MockApplicationRepository
}

// NewMockApplicationRepository creates a new mock instance.
func NewMockApplicationRepository(ctrl *gomock.Controller) *MockApplicationRepository {
	mock := &MockApplicationRepository{ctrl: ctrl}
	mock.recorder = &MockApplicationRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationRepository) EXPECT() *MockApplicationRepositoryMockRecorder {
	return m.recorder
}

// AddWebsite mocks base method.
func (m *MockApplicationRepository) AddWebsite(ctx context.Context, applicationID string, website *domain.Website) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWebsite", ctx, applicationID, website)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddWebsite indicates an expected call of AddWebsite.
func (mr *MockApplicationRepositoryMockRecorder) AddWebsite(ctx, applicationID, website interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWebsite", reflect.TypeOf((*MockApplicationRepository)(nil).AddWebsite), ctx, applicationID, website)
}

// CreateApplication mocks base method.
func (m *MockApplicationRepository) CreateApplication(ctx context.Context, args domain.CreateApplicationArgs) (*domain.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateApplication", ctx, args)
	ret0, _ := ret[0].(*domain.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateApplication indicates an expected call of CreateApplication.
func (mr *MockApplicationRepositoryMockRecorder) CreateApplication(ctx, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateApplication", reflect.TypeOf((*MockApplicationRepository)(nil).CreateApplication), ctx, args)
}

// DeleteWebsite mocks base method.
func (m *MockApplicationRepository) DeleteWebsite(ctx context.Context, applicationID, websiteID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWebsite", ctx, applicationID, websiteID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWebsite indicates an expected call of DeleteWebsite.
func (mr *MockApplicationRepositoryMockRecorder) DeleteWebsite(ctx, applicationID, websiteID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWebsite", reflect.TypeOf((*MockApplicationRepository)(nil).DeleteWebsite), ctx, applicationID, websiteID)
}

// GetApplication mocks base method.
func (m *MockApplicationRepository) GetApplication(ctx context.Context, id string) (*domain.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApplication", ctx, id)
	ret0, _ := ret[0].(*domain.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApplication indicates an expected call of GetApplication.
func (mr *MockApplicationRepositoryMockRecorder) GetApplication(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplication", reflect.TypeOf((*MockApplicationRepository)(nil).GetApplication), ctx, id)
}

// GetApplications mocks base method.
func (m *MockApplicationRepository) GetApplications(ctx context.Context, cond domain.GetApplicationCondition) ([]*domain.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApplications", ctx, cond)
	ret0, _ := ret[0].([]*domain.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApplications indicates an expected call of GetApplications.
func (mr *MockApplicationRepositoryMockRecorder) GetApplications(ctx, cond interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplications", reflect.TypeOf((*MockApplicationRepository)(nil).GetApplications), ctx, cond)
}

// GetWebsites mocks base method.
func (m *MockApplicationRepository) GetWebsites(ctx context.Context, applicationIDs []string) ([]*domain.Website, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWebsites", ctx, applicationIDs)
	ret0, _ := ret[0].([]*domain.Website)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWebsites indicates an expected call of GetWebsites.
func (mr *MockApplicationRepositoryMockRecorder) GetWebsites(ctx, applicationIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWebsites", reflect.TypeOf((*MockApplicationRepository)(nil).GetWebsites), ctx, applicationIDs)
}

// RegisterApplicationOwner mocks base method.
func (m *MockApplicationRepository) RegisterApplicationOwner(ctx context.Context, applicationID, userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterApplicationOwner", ctx, applicationID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterApplicationOwner indicates an expected call of RegisterApplicationOwner.
func (mr *MockApplicationRepositoryMockRecorder) RegisterApplicationOwner(ctx, applicationID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterApplicationOwner", reflect.TypeOf((*MockApplicationRepository)(nil).RegisterApplicationOwner), ctx, applicationID, userID)
}

// UpdateApplication mocks base method.
func (m *MockApplicationRepository) UpdateApplication(ctx context.Context, id string, args domain.UpdateApplicationArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateApplication", ctx, id, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateApplication indicates an expected call of UpdateApplication.
func (mr *MockApplicationRepositoryMockRecorder) UpdateApplication(ctx, id, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateApplication", reflect.TypeOf((*MockApplicationRepository)(nil).UpdateApplication), ctx, id, args)
}

// MockArtifactRepository is a mock of ArtifactRepository interface.
type MockArtifactRepository struct {
	ctrl     *gomock.Controller
	recorder *MockArtifactRepositoryMockRecorder
}

// MockArtifactRepositoryMockRecorder is the mock recorder for MockArtifactRepository.
type MockArtifactRepositoryMockRecorder struct {
	mock *MockArtifactRepository
}

// NewMockArtifactRepository creates a new mock instance.
func NewMockArtifactRepository(ctrl *gomock.Controller) *MockArtifactRepository {
	mock := &MockArtifactRepository{ctrl: ctrl}
	mock.recorder = &MockArtifactRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArtifactRepository) EXPECT() *MockArtifactRepositoryMockRecorder {
	return m.recorder
}

// CreateArtifact mocks base method.
func (m *MockArtifactRepository) CreateArtifact(ctx context.Context, size int64, buildID, sid string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateArtifact", ctx, size, buildID, sid)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateArtifact indicates an expected call of CreateArtifact.
func (mr *MockArtifactRepositoryMockRecorder) CreateArtifact(ctx, size, buildID, sid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateArtifact", reflect.TypeOf((*MockArtifactRepository)(nil).CreateArtifact), ctx, size, buildID, sid)
}

// MockAvailableDomainRepository is a mock of AvailableDomainRepository interface.
type MockAvailableDomainRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAvailableDomainRepositoryMockRecorder
}

// MockAvailableDomainRepositoryMockRecorder is the mock recorder for MockAvailableDomainRepository.
type MockAvailableDomainRepositoryMockRecorder struct {
	mock *MockAvailableDomainRepository
}

// NewMockAvailableDomainRepository creates a new mock instance.
func NewMockAvailableDomainRepository(ctrl *gomock.Controller) *MockAvailableDomainRepository {
	mock := &MockAvailableDomainRepository{ctrl: ctrl}
	mock.recorder = &MockAvailableDomainRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAvailableDomainRepository) EXPECT() *MockAvailableDomainRepositoryMockRecorder {
	return m.recorder
}

// AddAvailableDomain mocks base method.
func (m *MockAvailableDomainRepository) AddAvailableDomain(ctx context.Context, domain string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAvailableDomain", ctx, domain)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAvailableDomain indicates an expected call of AddAvailableDomain.
func (mr *MockAvailableDomainRepositoryMockRecorder) AddAvailableDomain(ctx, domain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAvailableDomain", reflect.TypeOf((*MockAvailableDomainRepository)(nil).AddAvailableDomain), ctx, domain)
}

// DeleteAvailableDomain mocks base method.
func (m *MockAvailableDomainRepository) DeleteAvailableDomain(ctx context.Context, domain string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAvailableDomain", ctx, domain)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAvailableDomain indicates an expected call of DeleteAvailableDomain.
func (mr *MockAvailableDomainRepositoryMockRecorder) DeleteAvailableDomain(ctx, domain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAvailableDomain", reflect.TypeOf((*MockAvailableDomainRepository)(nil).DeleteAvailableDomain), ctx, domain)
}

// GetAvailableDomains mocks base method.
func (m *MockAvailableDomainRepository) GetAvailableDomains(ctx context.Context) (domain.AvailableDomainSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvailableDomains", ctx)
	ret0, _ := ret[0].(domain.AvailableDomainSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvailableDomains indicates an expected call of GetAvailableDomains.
func (mr *MockAvailableDomainRepositoryMockRecorder) GetAvailableDomains(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailableDomains", reflect.TypeOf((*MockAvailableDomainRepository)(nil).GetAvailableDomains), ctx)
}

// MockBuildRepository is a mock of BuildRepository interface.
type MockBuildRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBuildRepositoryMockRecorder
}

// MockBuildRepositoryMockRecorder is the mock recorder for MockBuildRepository.
type MockBuildRepositoryMockRecorder struct {
	mock *MockBuildRepository
}

// NewMockBuildRepository creates a new mock instance.
func NewMockBuildRepository(ctrl *gomock.Controller) *MockBuildRepository {
	mock := &MockBuildRepository{ctrl: ctrl}
	mock.recorder = &MockBuildRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildRepository) EXPECT() *MockBuildRepositoryMockRecorder {
	return m.recorder
}

// CreateBuild mocks base method.
func (m *MockBuildRepository) CreateBuild(ctx context.Context, build *domain.Build) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBuild", ctx, build)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBuild indicates an expected call of CreateBuild.
func (mr *MockBuildRepositoryMockRecorder) CreateBuild(ctx, build interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBuild", reflect.TypeOf((*MockBuildRepository)(nil).CreateBuild), ctx, build)
}

// GetBuild mocks base method.
func (m *MockBuildRepository) GetBuild(ctx context.Context, buildID string) (*domain.Build, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBuild", ctx, buildID)
	ret0, _ := ret[0].(*domain.Build)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBuild indicates an expected call of GetBuild.
func (mr *MockBuildRepositoryMockRecorder) GetBuild(ctx, buildID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBuild", reflect.TypeOf((*MockBuildRepository)(nil).GetBuild), ctx, buildID)
}

// GetBuilds mocks base method.
func (m *MockBuildRepository) GetBuilds(ctx context.Context, condition domain.GetBuildCondition) ([]*domain.Build, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBuilds", ctx, condition)
	ret0, _ := ret[0].([]*domain.Build)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBuilds indicates an expected call of GetBuilds.
func (mr *MockBuildRepositoryMockRecorder) GetBuilds(ctx, condition interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBuilds", reflect.TypeOf((*MockBuildRepository)(nil).GetBuilds), ctx, condition)
}

// MarkCommitAsRetriable mocks base method.
func (m *MockBuildRepository) MarkCommitAsRetriable(ctx context.Context, applicationID, commit string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkCommitAsRetriable", ctx, applicationID, commit)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkCommitAsRetriable indicates an expected call of MarkCommitAsRetriable.
func (mr *MockBuildRepositoryMockRecorder) MarkCommitAsRetriable(ctx, applicationID, commit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkCommitAsRetriable", reflect.TypeOf((*MockBuildRepository)(nil).MarkCommitAsRetriable), ctx, applicationID, commit)
}

// UpdateBuild mocks base method.
func (m *MockBuildRepository) UpdateBuild(ctx context.Context, id string, args domain.UpdateBuildArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBuild", ctx, id, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBuild indicates an expected call of UpdateBuild.
func (mr *MockBuildRepositoryMockRecorder) UpdateBuild(ctx, id, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBuild", reflect.TypeOf((*MockBuildRepository)(nil).UpdateBuild), ctx, id, args)
}

// MockEnvironmentRepository is a mock of EnvironmentRepository interface.
type MockEnvironmentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockEnvironmentRepositoryMockRecorder
}

// MockEnvironmentRepositoryMockRecorder is the mock recorder for MockEnvironmentRepository.
type MockEnvironmentRepositoryMockRecorder struct {
	mock *MockEnvironmentRepository
}

// NewMockEnvironmentRepository creates a new mock instance.
func NewMockEnvironmentRepository(ctrl *gomock.Controller) *MockEnvironmentRepository {
	mock := &MockEnvironmentRepository{ctrl: ctrl}
	mock.recorder = &MockEnvironmentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnvironmentRepository) EXPECT() *MockEnvironmentRepositoryMockRecorder {
	return m.recorder
}

// GetEnv mocks base method.
func (m *MockEnvironmentRepository) GetEnv(ctx context.Context, applicationID string) ([]*domain.Environment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnv", ctx, applicationID)
	ret0, _ := ret[0].([]*domain.Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEnv indicates an expected call of GetEnv.
func (mr *MockEnvironmentRepositoryMockRecorder) GetEnv(ctx, applicationID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnv", reflect.TypeOf((*MockEnvironmentRepository)(nil).GetEnv), ctx, applicationID)
}

// SetEnv mocks base method.
func (m *MockEnvironmentRepository) SetEnv(ctx context.Context, applicationID, key, value string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetEnv", ctx, applicationID, key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetEnv indicates an expected call of SetEnv.
func (mr *MockEnvironmentRepositoryMockRecorder) SetEnv(ctx, applicationID, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEnv", reflect.TypeOf((*MockEnvironmentRepository)(nil).SetEnv), ctx, applicationID, key, value)
}

// MockGitRepositoryRepository is a mock of GitRepositoryRepository interface.
type MockGitRepositoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockGitRepositoryRepositoryMockRecorder
}

// MockGitRepositoryRepositoryMockRecorder is the mock recorder for MockGitRepositoryRepository.
type MockGitRepositoryRepositoryMockRecorder struct {
	mock *MockGitRepositoryRepository
}

// NewMockGitRepositoryRepository creates a new mock instance.
func NewMockGitRepositoryRepository(ctrl *gomock.Controller) *MockGitRepositoryRepository {
	mock := &MockGitRepositoryRepository{ctrl: ctrl}
	mock.recorder = &MockGitRepositoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGitRepositoryRepository) EXPECT() *MockGitRepositoryRepositoryMockRecorder {
	return m.recorder
}

// GetRepository mocks base method.
func (m *MockGitRepositoryRepository) GetRepository(ctx context.Context, rawURL string) (domain.Repository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepository", ctx, rawURL)
	ret0, _ := ret[0].(domain.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepository indicates an expected call of GetRepository.
func (mr *MockGitRepositoryRepositoryMockRecorder) GetRepository(ctx, rawURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepository", reflect.TypeOf((*MockGitRepositoryRepository)(nil).GetRepository), ctx, rawURL)
}

// GetRepositoryByID mocks base method.
func (m *MockGitRepositoryRepository) GetRepositoryByID(ctx context.Context, id string) (domain.Repository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepositoryByID", ctx, id)
	ret0, _ := ret[0].(domain.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepositoryByID indicates an expected call of GetRepositoryByID.
func (mr *MockGitRepositoryRepositoryMockRecorder) GetRepositoryByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepositoryByID", reflect.TypeOf((*MockGitRepositoryRepository)(nil).GetRepositoryByID), ctx, id)
}

// RegisterRepository mocks base method.
func (m *MockGitRepositoryRepository) RegisterRepository(ctx context.Context, args domain.RegisterRepositoryArgs) (domain.Repository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterRepository", ctx, args)
	ret0, _ := ret[0].(domain.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterRepository indicates an expected call of RegisterRepository.
func (mr *MockGitRepositoryRepositoryMockRecorder) RegisterRepository(ctx, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterRepository", reflect.TypeOf((*MockGitRepositoryRepository)(nil).RegisterRepository), ctx, args)
}

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserRepository) CreateUser(ctx context.Context, args domain.CreateUserArgs) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, args)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryMockRecorder) CreateUser(ctx, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepository)(nil).CreateUser), ctx, args)
}

// GetUserByID mocks base method.
func (m *MockUserRepository) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, id)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockUserRepositoryMockRecorder) GetUserByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserRepository)(nil).GetUserByID), ctx, id)
}
