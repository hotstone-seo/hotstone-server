// Code generated by MockGen. DO NOT EDIT.
// Source: app/service/center_service.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	service "github.com/hotstone-seo/hotstone-server/app/service"
	reflect "reflect"
)

// MockCenterService is a mock of CenterService interface
type MockCenterService struct {
	ctrl     *gomock.Controller
	recorder *MockCenterServiceMockRecorder
}

// MockCenterServiceMockRecorder is the mock recorder for MockCenterService
type MockCenterServiceMockRecorder struct {
	mock *MockCenterService
}

// NewMockCenterService creates a new mock instance
func NewMockCenterService(ctrl *gomock.Controller) *MockCenterService {
	mock := &MockCenterService{ctrl: ctrl}
	mock.recorder = &MockCenterServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCenterService) EXPECT() *MockCenterServiceMockRecorder {
	return m.recorder
}

// AddMetaTag mocks base method
func (m *MockCenterService) AddMetaTag(req service.AddMetaTagRequest) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMetaTag", req)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddMetaTag indicates an expected call of AddMetaTag
func (mr *MockCenterServiceMockRecorder) AddMetaTag(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMetaTag", reflect.TypeOf((*MockCenterService)(nil).AddMetaTag), req)
}

// AddTitleTag mocks base method
func (m *MockCenterService) AddTitleTag(req service.AddTitleTagRequest) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTitleTag", req)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTitleTag indicates an expected call of AddTitleTag
func (mr *MockCenterServiceMockRecorder) AddTitleTag(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTitleTag", reflect.TypeOf((*MockCenterService)(nil).AddTitleTag), req)
}

// AddCanonicalTag mocks base method
func (m *MockCenterService) AddCanonicalTag(req service.AddCanonicalTagRequest) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCanonicalTag", req)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCanonicalTag indicates an expected call of AddCanonicalTag
func (mr *MockCenterServiceMockRecorder) AddCanonicalTag(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCanonicalTag", reflect.TypeOf((*MockCenterService)(nil).AddCanonicalTag), req)
}
