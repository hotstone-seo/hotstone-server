// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hotstone-seo/hotstone-seo/server/service (interfaces: ClientKeyService)

// Package service_mock is a generated GoMock package.
package service_mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	repository "github.com/hotstone-seo/hotstone-seo/server/repository"
	dbkit "github.com/typical-go/typical-rest-server/pkg/dbkit"
	reflect "reflect"
)

// MockClientKeyService is a mock of ClientKeyService interface
type MockClientKeyService struct {
	ctrl     *gomock.Controller
	recorder *MockClientKeyServiceMockRecorder
}

// MockClientKeyServiceMockRecorder is the mock recorder for MockClientKeyService
type MockClientKeyServiceMockRecorder struct {
	mock *MockClientKeyService
}

// NewMockClientKeyService creates a new mock instance
func NewMockClientKeyService(ctrl *gomock.Controller) *MockClientKeyService {
	mock := &MockClientKeyService{ctrl: ctrl}
	mock.recorder = &MockClientKeyServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientKeyService) EXPECT() *MockClientKeyServiceMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockClientKeyService) Delete(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockClientKeyServiceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClientKeyService)(nil).Delete), arg0, arg1)
}

// Find mocks base method
func (m *MockClientKeyService) Find(arg0 context.Context, arg1 ...dbkit.FindOption) ([]*repository.ClientKey, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Find", varargs...)
	ret0, _ := ret[0].([]*repository.ClientKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockClientKeyServiceMockRecorder) Find(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockClientKeyService)(nil).Find), varargs...)
}

// FindOne mocks base method
func (m *MockClientKeyService) FindOne(arg0 context.Context, arg1 ...dbkit.FindOption) (*repository.ClientKey, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOne", varargs...)
	ret0, _ := ret[0].(*repository.ClientKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne
func (mr *MockClientKeyServiceMockRecorder) FindOne(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockClientKeyService)(nil).FindOne), varargs...)
}

// Insert mocks base method
func (m *MockClientKeyService) Insert(arg0 context.Context, arg1 repository.ClientKey) (repository.ClientKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(repository.ClientKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert
func (mr *MockClientKeyServiceMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockClientKeyService)(nil).Insert), arg0, arg1)
}

// IsValidClientKey mocks base method
func (m *MockClientKeyService) IsValidClientKey(arg0 context.Context, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidClientKey", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValidClientKey indicates an expected call of IsValidClientKey
func (mr *MockClientKeyServiceMockRecorder) IsValidClientKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidClientKey", reflect.TypeOf((*MockClientKeyService)(nil).IsValidClientKey), arg0, arg1)
}

// Update mocks base method
func (m *MockClientKeyService) Update(arg0 context.Context, arg1 repository.ClientKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockClientKeyServiceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockClientKeyService)(nil).Update), arg0, arg1)
}
