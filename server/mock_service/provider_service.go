// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hotstone-seo/hotstone-seo/server/service (interfaces: ProviderService)

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	cachekit "github.com/hotstone-seo/hotstone-seo/pkg/cachekit"
	service "github.com/hotstone-seo/hotstone-seo/server/service"
	url "net/url"
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

// FetchTags mocks base method
func (m *MockProviderService) FetchTags(arg0 context.Context, arg1 url.Values) ([]*service.ITag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchTags", arg0, arg1)
	ret0, _ := ret[0].([]*service.ITag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchTags indicates an expected call of FetchTags
func (mr *MockProviderServiceMockRecorder) FetchTags(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchTags", reflect.TypeOf((*MockProviderService)(nil).FetchTags), arg0, arg1)
}

// FetchTagsWithCache mocks base method
func (m *MockProviderService) FetchTagsWithCache(arg0 context.Context, arg1 url.Values, arg2 *cachekit.Pragma) ([]*service.ITag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchTagsWithCache", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*service.ITag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchTagsWithCache indicates an expected call of FetchTagsWithCache
func (mr *MockProviderServiceMockRecorder) FetchTagsWithCache(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchTagsWithCache", reflect.TypeOf((*MockProviderService)(nil).FetchTagsWithCache), arg0, arg1, arg2)
}

// Match mocks base method
func (m *MockProviderService) Match(arg0 context.Context, arg1 url.Values) (*service.MatchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Match", arg0, arg1)
	ret0, _ := ret[0].(*service.MatchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Match indicates an expected call of Match
func (mr *MockProviderServiceMockRecorder) Match(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Match", reflect.TypeOf((*MockProviderService)(nil).Match), arg0, arg1)
}
