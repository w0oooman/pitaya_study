// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/topfreegames/pitaya/v2/metrics (interfaces: Client)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockClient) Count(arg0 string, arg1 int64, arg2 []string, arg3 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Count indicates an expected call of Count.
func (mr *MockClientMockRecorder) Count(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockClient)(nil).Count), arg0, arg1, arg2, arg3)
}

// Gauge mocks base method.
func (m *MockClient) Gauge(arg0 string, arg1 float64, arg2 []string, arg3 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Gauge", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Gauge indicates an expected call of Gauge.
func (mr *MockClientMockRecorder) Gauge(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Gauge", reflect.TypeOf((*MockClient)(nil).Gauge), arg0, arg1, arg2, arg3)
}

// TimeInMilliseconds mocks base method.
func (m *MockClient) TimeInMilliseconds(arg0 string, arg1 float64, arg2 []string, arg3 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TimeInMilliseconds", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// TimeInMilliseconds indicates an expected call of TimeInMilliseconds.
func (mr *MockClientMockRecorder) TimeInMilliseconds(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TimeInMilliseconds", reflect.TypeOf((*MockClient)(nil).TimeInMilliseconds), arg0, arg1, arg2, arg3)
}