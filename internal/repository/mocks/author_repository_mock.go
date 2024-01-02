// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/author_repository.go
//
// Generated by this command:
//
//	mockgen -source=internal/repository/author_repository.go -destination=internal/repository/mocks/author_repository_mock.go
//

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	dto "bphn/artikel-hukum/internal/dto"
	model "bphn/artikel-hukum/internal/model"
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockAuthorRepository is a mock of AuthorRepository interface.
type MockAuthorRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorRepositoryMockRecorder
}

// MockAuthorRepositoryMockRecorder is the mock recorder for MockAuthorRepository.
type MockAuthorRepositoryMockRecorder struct {
	mock *MockAuthorRepository
}

// NewMockAuthorRepository creates a new mock instance.
func NewMockAuthorRepository(ctrl *gomock.Controller) *MockAuthorRepository {
	mock := &MockAuthorRepository{ctrl: ctrl}
	mock.recorder = &MockAuthorRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorRepository) EXPECT() *MockAuthorRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAuthorRepository) Create(ctx context.Context, detail model.AuthorDetail) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, detail)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockAuthorRepositoryMockRecorder) Create(ctx, detail any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAuthorRepository)(nil).Create), ctx, detail)
}

// Delete mocks base method.
func (m *MockAuthorRepository) Delete(ctx context.Context, id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAuthorRepositoryMockRecorder) Delete(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAuthorRepository)(nil).Delete), ctx, id)
}

// FindAll mocks base method.
func (m *MockAuthorRepository) FindAll(ctx context.Context, query dto.ListQuery) (dto.ListQueryResult[model.AuthorDetail], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx, query)
	ret0, _ := ret[0].(dto.ListQueryResult[model.AuthorDetail])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockAuthorRepositoryMockRecorder) FindAll(ctx, query any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockAuthorRepository)(nil).FindAll), ctx, query)
}

// FindById mocks base method.
func (m *MockAuthorRepository) FindById(ctx context.Context, id uint) (*model.AuthorDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", ctx, id)
	ret0, _ := ret[0].(*model.AuthorDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockAuthorRepositoryMockRecorder) FindById(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockAuthorRepository)(nil).FindById), ctx, id)
}

// Update mocks base method.
func (m *MockAuthorRepository) Update(ctx context.Context, detail model.AuthorDetail) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, detail)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockAuthorRepositoryMockRecorder) Update(ctx, detail any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAuthorRepository)(nil).Update), ctx, detail)
}
