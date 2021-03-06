// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hotstone-seo/hotstone-seo/internal/api/service (interfaces: StructuredDataService)

// Package service_mock is a generated GoMock package.
package service_mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	repository "github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	dbkit "github.com/typical-go/typical-rest-server/pkg/dbkit"
	reflect "reflect"
)

// MockStructuredDataService is a mock of StructuredDataService interface
type MockStructuredDataService struct {
	ctrl     *gomock.Controller
	recorder *MockStructuredDataServiceMockRecorder
}

// MockStructuredDataServiceMockRecorder is the mock recorder for MockStructuredDataService
type MockStructuredDataServiceMockRecorder struct {
	mock *MockStructuredDataService
}

// NewMockStructuredDataService creates a new mock instance
func NewMockStructuredDataService(ctrl *gomock.Controller) *MockStructuredDataService {
	mock := &MockStructuredDataService{ctrl: ctrl}
	mock.recorder = &MockStructuredDataServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStructuredDataService) EXPECT() *MockStructuredDataServiceMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockStructuredDataService) Delete(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockStructuredDataServiceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStructuredDataService)(nil).Delete), arg0, arg1)
}

// Find mocks base method
func (m *MockStructuredDataService) Find(arg0 context.Context, arg1 ...dbkit.SelectOption) ([]*repository.StructuredData, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Find", varargs...)
	ret0, _ := ret[0].([]*repository.StructuredData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockStructuredDataServiceMockRecorder) Find(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockStructuredDataService)(nil).Find), varargs...)
}

// FindByRule mocks base method
func (m *MockStructuredDataService) FindByRule(arg0 context.Context, arg1 int64) ([]*repository.StructuredData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByRule", arg0, arg1)
	ret0, _ := ret[0].([]*repository.StructuredData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByRule indicates an expected call of FindByRule
func (mr *MockStructuredDataServiceMockRecorder) FindByRule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByRule", reflect.TypeOf((*MockStructuredDataService)(nil).FindByRule), arg0, arg1)
}

// FindOne mocks base method
func (m *MockStructuredDataService) FindOne(arg0 context.Context, arg1 int64) (*repository.StructuredData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", arg0, arg1)
	ret0, _ := ret[0].(*repository.StructuredData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne
func (mr *MockStructuredDataServiceMockRecorder) FindOne(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockStructuredDataService)(nil).FindOne), arg0, arg1)
}

// Insert mocks base method
func (m *MockStructuredDataService) Insert(arg0 context.Context, arg1 repository.StructuredData) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert
func (mr *MockStructuredDataServiceMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockStructuredDataService)(nil).Insert), arg0, arg1)
}

// Update mocks base method
func (m *MockStructuredDataService) Update(arg0 context.Context, arg1 repository.StructuredData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockStructuredDataServiceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStructuredDataService)(nil).Update), arg0, arg1)
}
