// Code generated by MockGen. DO NOT EDIT.
// Source: tag.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	gomock "github.com/golang/mock/gomock"
	controller "homepage/pkg/interface/controller"
	reflect "reflect"
)

// MockTagController is a mock of TagController interface
type MockTagController struct {
	ctrl     *gomock.Controller
	recorder *MockTagControllerMockRecorder
}

// MockTagControllerMockRecorder is the mock recorder for MockTagController
type MockTagControllerMockRecorder struct {
	mock *MockTagController
}

// NewMockTagController creates a new mock instance
func NewMockTagController(ctrl *gomock.Controller) *MockTagController {
	mock := &MockTagController{ctrl: ctrl}
	mock.recorder = &MockTagControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTagController) EXPECT() *MockTagControllerMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockTagController) GetAll() (*controller.TagsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].(*controller.TagsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockTagControllerMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockTagController)(nil).GetAll))
}

// GetByID mocks base method
func (m *MockTagController) GetByID(id int) (*controller.TagResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*controller.TagResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockTagControllerMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockTagController)(nil).GetByID), id)
}

// Create mocks base method
func (m *MockTagController) Create(name string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", name)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockTagControllerMockRecorder) Create(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTagController)(nil).Create), name)
}

// UpdateByID mocks base method
func (m *MockTagController) UpdateByID(id int, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByID", id, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateByID indicates an expected call of UpdateByID
func (mr *MockTagControllerMockRecorder) UpdateByID(id, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByID", reflect.TypeOf((*MockTagController)(nil).UpdateByID), id, name)
}

// DeleteByID mocks base method
func (m *MockTagController) DeleteByID(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID
func (mr *MockTagControllerMockRecorder) DeleteByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockTagController)(nil).DeleteByID), id)
}

// AdminGetAll mocks base method
func (m *MockTagController) AdminGetAll() ([]map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdminGetAll")
	ret0, _ := ret[0].([]map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AdminGetAll indicates an expected call of AdminGetAll
func (mr *MockTagControllerMockRecorder) AdminGetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminGetAll", reflect.TypeOf((*MockTagController)(nil).AdminGetAll))
}

// AdminGetByID mocks base method
func (m *MockTagController) AdminGetByID(id int) (*controller.FieldsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdminGetByID", id)
	ret0, _ := ret[0].(*controller.FieldsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AdminGetByID indicates an expected call of AdminGetByID
func (mr *MockTagControllerMockRecorder) AdminGetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminGetByID", reflect.TypeOf((*MockTagController)(nil).AdminGetByID), id)
}