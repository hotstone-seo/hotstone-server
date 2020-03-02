// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hotstone-seo/hotstone-seo/server/service (interfaces: URLService)

// Package mock_service is a generated GoMock package.
package mock_service

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockURLService is a mock of URLService interface
type MockURLService struct {
	ctrl     *gomock.Controller
	recorder *MockURLServiceMockRecorder
}

// MockURLServiceMockRecorder is the mock recorder for MockURLService
type MockURLServiceMockRecorder struct {
	mock *MockURLService
}

// NewMockURLService creates a new mock instance
func NewMockURLService(ctrl *gomock.Controller) *MockURLService {
	mock := &MockURLService{ctrl: ctrl}
	mock.recorder = &MockURLServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockURLService) EXPECT() *MockURLServiceMockRecorder {
	return m.recorder
}

// Count mocks base method
func (m *MockURLService) Count() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count")
	ret0, _ := ret[0].(int)
	return ret0
}

// Count indicates an expected call of Count
func (mr *MockURLServiceMockRecorder) Count() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockURLService)(nil).Count))
}

// Delete mocks base method
func (m *MockURLService) Delete(arg0 int64) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockURLServiceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockURLService)(nil).Delete), arg0)
}

// DumpTree mocks base method
func (m *MockURLService) DumpTree() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DumpTree")
	ret0, _ := ret[0].(string)
	return ret0
}

// DumpTree indicates an expected call of DumpTree
func (mr *MockURLServiceMockRecorder) DumpTree() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DumpTree", reflect.TypeOf((*MockURLService)(nil).DumpTree))
}

// FullSync mocks base method
func (m *MockURLService) FullSync() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FullSync")
	ret0, _ := ret[0].(error)
	return ret0
}

// FullSync indicates an expected call of FullSync
func (mr *MockURLServiceMockRecorder) FullSync() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FullSync", reflect.TypeOf((*MockURLService)(nil).FullSync))
}

// Get mocks base method
func (m *MockURLService) Get(arg0 string, arg1 []string) (interface{}, []string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].([]string)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockURLServiceMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockURLService)(nil).Get), arg0, arg1)
}

// Insert mocks base method
func (m *MockURLService) Insert(arg0 int64, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Insert", arg0, arg1)
}

// Insert indicates an expected call of Insert
func (mr *MockURLServiceMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockURLService)(nil).Insert), arg0, arg1)
}

// Match mocks base method
func (m *MockURLService) Match(arg0 string) (int, map[string]string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Match", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(map[string]string)
	return ret0, ret1
}

// Match indicates an expected call of Match
func (mr *MockURLServiceMockRecorder) Match(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Match", reflect.TypeOf((*MockURLService)(nil).Match), arg0)
}

// Sync mocks base method
func (m *MockURLService) Sync() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sync")
	ret0, _ := ret[0].(error)
	return ret0
}

// Sync indicates an expected call of Sync
func (mr *MockURLServiceMockRecorder) Sync() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockURLService)(nil).Sync))
}

// Update mocks base method
func (m *MockURLService) Update(arg0 int64, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Update", arg0, arg1)
}

// Update indicates an expected call of Update
func (mr *MockURLServiceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockURLService)(nil).Update), arg0, arg1)
}