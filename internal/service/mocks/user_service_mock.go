// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/user_service.go
//
// Generated by this command:
//
//	mockgen -source=internal/service/user_service.go -destination=internal/service/mocks/user_service_mock.go
//

// Package mock_service is a generated GoMock package.
package mock_service

import (
	v1 "bphn/artikel-hukum/api/v1"
	ito "bphn/artikel-hukum/internal/ito"
	model "bphn/artikel-hukum/internal/model"
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// ChangePasswordByNonAdmin mocks base method.
func (m *MockUserService) ChangePasswordByNonAdmin(ctx context.Context, request ito.ChangePasswordQuery) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePasswordByNonAdmin", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePasswordByNonAdmin indicates an expected call of ChangePasswordByNonAdmin.
func (mr *MockUserServiceMockRecorder) ChangePasswordByNonAdmin(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePasswordByNonAdmin", reflect.TypeOf((*MockUserService)(nil).ChangePasswordByNonAdmin), ctx, request)
}

// Create mocks base method.
func (m *MockUserService) Create(ctx context.Context, request *v1.CreateUserRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserServiceMockRecorder) Create(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserService)(nil).Create), ctx, request)
}

// Delete mocks base method.
func (m *MockUserService) Delete(ctx context.Context, request uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserServiceMockRecorder) Delete(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserService)(nil).Delete), ctx, request)
}

// FindByEmail mocks base method.
func (m *MockUserService) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", ctx, email)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockUserServiceMockRecorder) FindByEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockUserService)(nil).FindByEmail), ctx, email)
}

// FindById mocks base method.
func (m *MockUserService) FindById(ctx context.Context, id uint) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", ctx, id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockUserServiceMockRecorder) FindById(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockUserService)(nil).FindById), ctx, id)
}

// ForgotPassword mocks base method.
func (m *MockUserService) ForgotPassword(ctx context.Context, request v1.ForgotPasswordRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForgotPassword", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForgotPassword indicates an expected call of ForgotPassword.
func (mr *MockUserServiceMockRecorder) ForgotPassword(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForgotPassword", reflect.TypeOf((*MockUserService)(nil).ForgotPassword), ctx, request)
}

// List mocks base method.
func (m *MockUserService) List(ctx context.Context, query ito.ListQuery) (*ito.ListQueryResult[v1.UserDataResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, query)
	ret0, _ := ret[0].(*ito.ListQueryResult[v1.UserDataResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockUserServiceMockRecorder) List(ctx, query any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUserService)(nil).List), ctx, query)
}

// Update mocks base method.
func (m *MockUserService) Update(ctx context.Context, request *v1.UpdateUserRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserServiceMockRecorder) Update(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserService)(nil).Update), ctx, request)
}
