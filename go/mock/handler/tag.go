// Code generated by MockGen. DO NOT EDIT.
// Source: tag.go

// Package mock_handler is a generated GoMock package.
package mock_handler

import (
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockTagHandler is a mock of TagHandler interface
type MockTagHandler struct {
	ctrl     *gomock.Controller
	recorder *MockTagHandlerMockRecorder
}

// MockTagHandlerMockRecorder is the mock recorder for MockTagHandler
type MockTagHandlerMockRecorder struct {
	mock *MockTagHandler
}

// NewMockTagHandler creates a new mock instance
func NewMockTagHandler(ctrl *gomock.Controller) *MockTagHandler {
	mock := &MockTagHandler{ctrl: ctrl}
	mock.recorder = &MockTagHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTagHandler) EXPECT() *MockTagHandlerMockRecorder {
	return m.recorder
}

// AdminCreate mocks base method
func (m *MockTagHandler) AdminCreate(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AdminCreate", w, r)
}

// AdminCreate indicates an expected call of AdminCreate
func (mr *MockTagHandlerMockRecorder) AdminCreate(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminCreate", reflect.TypeOf((*MockTagHandler)(nil).AdminCreate), w, r)
}

// AdminUpdateByID mocks base method
func (m *MockTagHandler) AdminUpdateByID(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AdminUpdateByID", w, r)
}

// AdminUpdateByID indicates an expected call of AdminUpdateByID
func (mr *MockTagHandlerMockRecorder) AdminUpdateByID(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminUpdateByID", reflect.TypeOf((*MockTagHandler)(nil).AdminUpdateByID), w, r)
}

// AdminGetAll mocks base method
func (m *MockTagHandler) AdminGetAll(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AdminGetAll", w, r)
}

// AdminGetAll indicates an expected call of AdminGetAll
func (mr *MockTagHandlerMockRecorder) AdminGetAll(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminGetAll", reflect.TypeOf((*MockTagHandler)(nil).AdminGetAll), w, r)
}

// AdminGetByID mocks base method
func (m *MockTagHandler) AdminGetByID(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AdminGetByID", w, r)
}

// AdminGetByID indicates an expected call of AdminGetByID
func (mr *MockTagHandlerMockRecorder) AdminGetByID(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminGetByID", reflect.TypeOf((*MockTagHandler)(nil).AdminGetByID), w, r)
}

// AdminDeleteByID mocks base method
func (m *MockTagHandler) AdminDeleteByID(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AdminDeleteByID", w, r)
}

// AdminDeleteByID indicates an expected call of AdminDeleteByID
func (mr *MockTagHandlerMockRecorder) AdminDeleteByID(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminDeleteByID", reflect.TypeOf((*MockTagHandler)(nil).AdminDeleteByID), w, r)
}