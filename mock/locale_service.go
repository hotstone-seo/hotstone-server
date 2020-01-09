// Code generated by MockGen. DO NOT EDIT.
// Source: app/service/locale_service.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	repository "github.com/hotstone-seo/hotstone-seo/app/repository"
	reflect "reflect"
)

// MockLocaleService is a mock of LocaleService interface
type MockLocaleService struct {
	ctrl     *gomock.Controller
	recorder *MockLocaleServiceMockRecorder
}

// MockLocaleServiceMockRecorder is the mock recorder for MockLocaleService
type MockLocaleServiceMockRecorder struct {
	mock *MockLocaleService
}

// NewMockLocaleService creates a new mock instance
func NewMockLocaleService(ctrl *gomock.Controller) *MockLocaleService {
	mock := &MockLocaleService{ctrl: ctrl}
	mock.recorder = &MockLocaleServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLocaleService) EXPECT() *MockLocaleServiceMockRecorder {
	return m.recorder
}

// FindOne mocks base method
func (m *MockLocaleService) FindOne(arg0 context.Context, arg1 int64) (*repository.Locale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", arg0, arg1)
	ret0, _ := ret[0].(*repository.Locale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne
func (mr *MockLocaleServiceMockRecorder) FindOne(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockLocaleService)(nil).FindOne), arg0, arg1)
}

// Find mocks base method
func (m *MockLocaleService) Find(arg0 context.Context) ([]*repository.Locale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].([]*repository.Locale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockLocaleServiceMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockLocaleService)(nil).Find), arg0)
}

// Insert mocks base method
func (m *MockLocaleService) Insert(arg0 context.Context, arg1 repository.Locale) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert
func (mr *MockLocaleServiceMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockLocaleService)(nil).Insert), arg0, arg1)
}

// Delete mocks base method
func (m *MockLocaleService) Delete(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockLocaleServiceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockLocaleService)(nil).Delete), arg0, arg1)
}

// Update mocks base method
func (m *MockLocaleService) Update(arg0 context.Context, arg1 repository.Locale) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockLocaleServiceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockLocaleService)(nil).Update), arg0, arg1)
}
