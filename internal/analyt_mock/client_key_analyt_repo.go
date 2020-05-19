// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hotstone-seo/hotstone-seo/internal/analyt (interfaces: ClientKeyAnalytRepo)

// Package analyt_mock is a generated GoMock package.
package analyt_mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockClientKeyAnalytRepo is a mock of ClientKeyAnalytRepo interface
type MockClientKeyAnalytRepo struct {
	ctrl     *gomock.Controller
	recorder *MockClientKeyAnalytRepoMockRecorder
}

// MockClientKeyAnalytRepoMockRecorder is the mock recorder for MockClientKeyAnalytRepo
type MockClientKeyAnalytRepoMockRecorder struct {
	mock *MockClientKeyAnalytRepo
}

// NewMockClientKeyAnalytRepo creates a new mock instance
func NewMockClientKeyAnalytRepo(ctrl *gomock.Controller) *MockClientKeyAnalytRepo {
	mock := &MockClientKeyAnalytRepo{ctrl: ctrl}
	mock.recorder = &MockClientKeyAnalytRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientKeyAnalytRepo) EXPECT() *MockClientKeyAnalytRepoMockRecorder {
	return m.recorder
}

// ClientKeyLastUsed mocks base method
func (m *MockClientKeyAnalytRepo) ClientKeyLastUsed(arg0 context.Context, arg1 string) (*time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientKeyLastUsed", arg0, arg1)
	ret0, _ := ret[0].(*time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientKeyLastUsed indicates an expected call of ClientKeyLastUsed
func (mr *MockClientKeyAnalytRepoMockRecorder) ClientKeyLastUsed(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientKeyLastUsed", reflect.TypeOf((*MockClientKeyAnalytRepo)(nil).ClientKeyLastUsed), arg0, arg1)
}
