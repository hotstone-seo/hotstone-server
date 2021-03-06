// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hotstone-seo/hotstone-seo/internal/api/service (interfaces: DataSourceService)

// Package service_mock is a generated GoMock package.
package service_mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	repository "github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	reflect "reflect"
)

// MockDataSourceService is a mock of DataSourceService interface
type MockDataSourceService struct {
	ctrl     *gomock.Controller
	recorder *MockDataSourceServiceMockRecorder
}

// MockDataSourceServiceMockRecorder is the mock recorder for MockDataSourceService
type MockDataSourceServiceMockRecorder struct {
	mock *MockDataSourceService
}

// NewMockDataSourceService creates a new mock instance
func NewMockDataSourceService(ctrl *gomock.Controller) *MockDataSourceService {
	mock := &MockDataSourceService{ctrl: ctrl}
	mock.recorder = &MockDataSourceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDataSourceService) EXPECT() *MockDataSourceServiceMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockDataSourceService) Delete(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockDataSourceServiceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDataSourceService)(nil).Delete), arg0, arg1)
}

// Find mocks base method
func (m *MockDataSourceService) Find(arg0 context.Context, arg1 repository.PaginationParam) ([]*repository.DataSource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].([]*repository.DataSource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockDataSourceServiceMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockDataSourceService)(nil).Find), arg0, arg1)
}

// FindOne mocks base method
func (m *MockDataSourceService) FindOne(arg0 context.Context, arg1 int64) (*repository.DataSource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", arg0, arg1)
	ret0, _ := ret[0].(*repository.DataSource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne
func (mr *MockDataSourceServiceMockRecorder) FindOne(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockDataSourceService)(nil).FindOne), arg0, arg1)
}

// Insert mocks base method
func (m *MockDataSourceService) Insert(arg0 context.Context, arg1 repository.DataSource) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert
func (mr *MockDataSourceServiceMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockDataSourceService)(nil).Insert), arg0, arg1)
}

// Update mocks base method
func (m *MockDataSourceService) Update(arg0 context.Context, arg1 repository.DataSource) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockDataSourceServiceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDataSourceService)(nil).Update), arg0, arg1)
}
