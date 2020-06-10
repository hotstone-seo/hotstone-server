// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hotstone-seo/hotstone-seo/internal/api/repository (interfaces: StructuredDataRepo)

// Package repository_mock is a generated GoMock package.
package repository_mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	repository "github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	dbkit "github.com/typical-go/typical-rest-server/pkg/dbkit"
	reflect "reflect"
)

// MockStructuredDataRepo is a mock of StructuredDataRepo interface
type MockStructuredDataRepo struct {
	ctrl     *gomock.Controller
	recorder *MockStructuredDataRepoMockRecorder
}

// MockStructuredDataRepoMockRecorder is the mock recorder for MockStructuredDataRepo
type MockStructuredDataRepoMockRecorder struct {
	mock *MockStructuredDataRepo
}

// NewMockStructuredDataRepo creates a new mock instance
func NewMockStructuredDataRepo(ctrl *gomock.Controller) *MockStructuredDataRepo {
	mock := &MockStructuredDataRepo{ctrl: ctrl}
	mock.recorder = &MockStructuredDataRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStructuredDataRepo) EXPECT() *MockStructuredDataRepoMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockStructuredDataRepo) Delete(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockStructuredDataRepoMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStructuredDataRepo)(nil).Delete), arg0, arg1)
}

// Find mocks base method
func (m *MockStructuredDataRepo) Find(arg0 context.Context, arg1 ...dbkit.SelectOption) ([]*repository.StructuredData, error) {
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
func (mr *MockStructuredDataRepoMockRecorder) Find(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockStructuredDataRepo)(nil).Find), varargs...)
}

// FindOne mocks base method
func (m *MockStructuredDataRepo) FindOne(arg0 context.Context, arg1 int64) (*repository.StructuredData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", arg0, arg1)
	ret0, _ := ret[0].(*repository.StructuredData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne
func (mr *MockStructuredDataRepoMockRecorder) FindOne(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockStructuredDataRepo)(nil).FindOne), arg0, arg1)
}

// Insert mocks base method
func (m *MockStructuredDataRepo) Insert(arg0 context.Context, arg1 repository.StructuredData) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert
func (mr *MockStructuredDataRepoMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockStructuredDataRepo)(nil).Insert), arg0, arg1)
}

// Update mocks base method
func (m *MockStructuredDataRepo) Update(arg0 context.Context, arg1 repository.StructuredData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockStructuredDataRepoMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStructuredDataRepo)(nil).Update), arg0, arg1)
}
