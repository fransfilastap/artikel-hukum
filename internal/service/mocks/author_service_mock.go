// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/author_service.go
//
// Generated by this command:
//
//	mockgen -source=internal/service/author_service.go -destination=internal/service/mocks/author_service_mock.go
//

// Package mock_service is a generated GoMock package.
package mock_service

import (
	v1 "bphn/artikel-hukum/api/v1"
	dto "bphn/artikel-hukum/internal/dto"
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockAuthorService is a mock of AuthorService interface.
type MockAuthorService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorServiceMockRecorder
}

// MockAuthorServiceMockRecorder is the mock recorder for MockAuthorService.
type MockAuthorServiceMockRecorder struct {
	mock *MockAuthorService
}

// NewMockAuthorService creates a new mock instance.
func NewMockAuthorService(ctrl *gomock.Controller) *MockAuthorService {
	mock := &MockAuthorService{ctrl: ctrl}
	mock.recorder = &MockAuthorServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorService) EXPECT() *MockAuthorServiceMockRecorder {
	return m.recorder
}

// ForgotPassword mocks base method.
func (m *MockAuthorService) ForgotPassword(ctx context.Context, request v1.ForgotPasswordRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForgotPassword", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForgotPassword indicates an expected call of ForgotPassword.
func (mr *MockAuthorServiceMockRecorder) ForgotPassword(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForgotPassword", reflect.TypeOf((*MockAuthorService)(nil).ForgotPassword), ctx, request)
}

// List mocks base method.
func (m *MockAuthorService) List(ctx context.Context, query dto.ListQuery) (dto.ListQueryResult[v1.AuthorProfileDataResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, query)
	ret0, _ := ret[0].(dto.ListQueryResult[v1.AuthorProfileDataResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockAuthorServiceMockRecorder) List(ctx, query any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAuthorService)(nil).List), ctx, query)
}

// Profile mocks base method.
func (m *MockAuthorService) Profile(ctx context.Context, Id uint) (v1.AuthorProfileDataResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Profile", ctx, Id)
	ret0, _ := ret[0].(v1.AuthorProfileDataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Profile indicates an expected call of Profile.
func (mr *MockAuthorServiceMockRecorder) Profile(ctx, Id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Profile", reflect.TypeOf((*MockAuthorService)(nil).Profile), ctx, Id)
}

// Register mocks base method.
func (m *MockAuthorService) Register(ctx context.Context, request v1.AuthorRegistrationRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockAuthorServiceMockRecorder) Register(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAuthorService)(nil).Register), ctx, request)
}

// UpdateProfile mocks base method.
func (m *MockAuthorService) UpdateProfile(ctx context.Context, request v1.UpdateAuthorProfileRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProfile indicates an expected call of UpdateProfile.
func (mr *MockAuthorServiceMockRecorder) UpdateProfile(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockAuthorService)(nil).UpdateProfile), ctx, request)
}
