// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/disaster37/go-centreon-rest/v21/api (interfaces: ServiceGroupAPI)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/disaster37/go-centreon-rest/v21/models"
	gomock "github.com/golang/mock/gomock"
)

// MockServiceGroupAPI is a mock of ServiceGroupAPI interface.
type MockServiceGroupAPI struct {
	ctrl     *gomock.Controller
	recorder *MockServiceGroupAPIMockRecorder
}

// MockServiceGroupAPIMockRecorder is the mock recorder for MockServiceGroupAPI.
type MockServiceGroupAPIMockRecorder struct {
	mock *MockServiceGroupAPI
}

// NewMockServiceGroupAPI creates a new mock instance.
func NewMockServiceGroupAPI(ctrl *gomock.Controller) *MockServiceGroupAPI {
	mock := &MockServiceGroupAPI{ctrl: ctrl}
	mock.recorder = &MockServiceGroupAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceGroupAPI) EXPECT() *MockServiceGroupAPIMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockServiceGroupAPI) Add(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockServiceGroupAPIMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockServiceGroupAPI)(nil).Add), arg0, arg1)
}

// Delete mocks base method.
func (m *MockServiceGroupAPI) Delete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceGroupAPIMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServiceGroupAPI)(nil).Delete), arg0)
}

// Get mocks base method.
func (m *MockServiceGroupAPI) Get(arg0 string) (*models.ServiceGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*models.ServiceGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockServiceGroupAPIMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockServiceGroupAPI)(nil).Get), arg0)
}

// GetParam mocks base method.
func (m *MockServiceGroupAPI) GetParam(arg0 string, arg1 []string) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParam", arg0, arg1)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParam indicates an expected call of GetParam.
func (mr *MockServiceGroupAPIMockRecorder) GetParam(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParam", reflect.TypeOf((*MockServiceGroupAPI)(nil).GetParam), arg0, arg1)
}

// List mocks base method.
func (m *MockServiceGroupAPI) List() ([]*models.ServiceGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*models.ServiceGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockServiceGroupAPIMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockServiceGroupAPI)(nil).List))
}

// SetParam mocks base method.
func (m *MockServiceGroupAPI) SetParam(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetParam", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetParam indicates an expected call of SetParam.
func (mr *MockServiceGroupAPIMockRecorder) SetParam(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetParam", reflect.TypeOf((*MockServiceGroupAPI)(nil).SetParam), arg0, arg1, arg2)
}
