// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hotstone-seo/hotstone-seo/server/service (interfaces: ProviderService)

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	repository "github.com/hotstone-seo/hotstone-seo/server/repository"
	service "github.com/hotstone-seo/hotstone-seo/server/service"
	reflect "reflect"
)

// MockProviderService is a mock of ProviderService interface
type MockProviderService struct {
	ctrl     *gomock.Controller
	recorder *MockProviderServiceMockRecorder
}

// MockProviderServiceMockRecorder is the mock recorder for MockProviderService
type MockProviderServiceMockRecorder struct {
	mock *MockProviderService
}

// NewMockProviderService creates a new mock instance
func NewMockProviderService(ctrl *gomock.Controller) *MockProviderService {
	mock := &MockProviderService{ctrl: ctrl}
	mock.recorder = &MockProviderServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProviderService) EXPECT() *MockProviderServiceMockRecorder {
	return m.recorder
}

// DumpRuleTree mocks base method
func (m *MockProviderService) DumpRuleTree(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DumpRuleTree", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DumpRuleTree indicates an expected call of DumpRuleTree
func (mr *MockProviderServiceMockRecorder) DumpRuleTree(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DumpRuleTree", reflect.TypeOf((*MockProviderService)(nil).DumpRuleTree), arg0)
}

// FetchTags mocks base method
func (m *MockProviderService) FetchTags(arg0 context.Context, arg1 int64, arg2 string) ([]*repository.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchTags", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*repository.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchTags indicates an expected call of FetchTags
func (mr *MockProviderServiceMockRecorder) FetchTags(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchTags", reflect.TypeOf((*MockProviderService)(nil).FetchTags), arg0, arg1, arg2)
}

// MatchRule mocks base method
func (m *MockProviderService) MatchRule(arg0 context.Context, arg1 service.MatchRuleRequest) (*service.MatchRuleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MatchRule", arg0, arg1)
	ret0, _ := ret[0].(*service.MatchRuleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MatchRule indicates an expected call of MatchRule
func (mr *MockProviderServiceMockRecorder) MatchRule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MatchRule", reflect.TypeOf((*MockProviderService)(nil).MatchRule), arg0, arg1)
}
