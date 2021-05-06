// Code generated by MockGen. DO NOT EDIT.
// Source: research.go

// Package mock_interactor is a generated GoMock package.
package mock_interactor

import (
	gomock "github.com/golang/mock/gomock"
	entity "homepage/pkg/domain/entity"
	reflect "reflect"
)

// MockResearchInteractor is a mock of ResearchInteractor interface
type MockResearchInteractor struct {
	ctrl     *gomock.Controller
	recorder *MockResearchInteractorMockRecorder
}

// MockResearchInteractorMockRecorder is the mock recorder for MockResearchInteractor
type MockResearchInteractorMockRecorder struct {
	mock *MockResearchInteractor
}

// NewMockResearchInteractor creates a new mock instance
func NewMockResearchInteractor(ctrl *gomock.Controller) *MockResearchInteractor {
	mock := &MockResearchInteractor{ctrl: ctrl}
	mock.recorder = &MockResearchInteractorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResearchInteractor) EXPECT() *MockResearchInteractorMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockResearchInteractor) GetAll() ([]*entity.Research, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*entity.Research)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockResearchInteractorMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockResearchInteractor)(nil).GetAll))
}

// GetByID mocks base method
func (m *MockResearchInteractor) GetByID(id int) (*entity.Research, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*entity.Research)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockResearchInteractorMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockResearchInteractor)(nil).GetByID), id)
}

// Create mocks base method
func (m *MockResearchInteractor) Create(title, author, file, comment string, activation int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", title, author, file, comment, activation)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockResearchInteractorMockRecorder) Create(title, author, file, comment, activation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockResearchInteractor)(nil).Create), title, author, file, comment, activation)
}

// UpdateByID mocks base method
func (m *MockResearchInteractor) UpdateByID(id int, title, author, file, comment string, activation int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByID", id, title, author, file, comment, activation)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateByID indicates an expected call of UpdateByID
func (mr *MockResearchInteractorMockRecorder) UpdateByID(id, title, author, file, comment, activation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByID", reflect.TypeOf((*MockResearchInteractor)(nil).UpdateByID), id, title, author, file, comment, activation)
}

// DeleteByID mocks base method
func (m *MockResearchInteractor) DeleteByID(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID
func (mr *MockResearchInteractorMockRecorder) DeleteByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockResearchInteractor)(nil).DeleteByID), id)
}