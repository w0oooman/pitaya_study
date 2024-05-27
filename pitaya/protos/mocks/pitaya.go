// Code generated by MockGen. DO NOT EDIT.
// Source: pitaya_grpc.pb.go
//
// Generated by this command:
//
//	mockgen -source=pitaya_grpc.pb.go -destination=./mocks/pitaya.go -package=mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	protos "github.com/topfreegames/pitaya/v2/protos"
	grpc "google.golang.org/grpc"
)

// MockPitayaClient is a mock of PitayaClient interface.
type MockPitayaClient struct {
	ctrl     *gomock.Controller
	recorder *MockPitayaClientMockRecorder
}

// MockPitayaClientMockRecorder is the mock recorder for MockPitayaClient.
type MockPitayaClientMockRecorder struct {
	mock *MockPitayaClient
}

// NewMockPitayaClient creates a new mock instance.
func NewMockPitayaClient(ctrl *gomock.Controller) *MockPitayaClient {
	mock := &MockPitayaClient{ctrl: ctrl}
	mock.recorder = &MockPitayaClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPitayaClient) EXPECT() *MockPitayaClientMockRecorder {
	return m.recorder
}

// Call mocks base method.
func (m *MockPitayaClient) Call(ctx context.Context, in *protos.Request, opts ...grpc.CallOption) (*protos.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Call", varargs...)
	ret0, _ := ret[0].(*protos.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Call indicates an expected call of Call.
func (mr *MockPitayaClientMockRecorder) Call(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockPitayaClient)(nil).Call), varargs...)
}

// KickUser mocks base method.
func (m *MockPitayaClient) KickUser(ctx context.Context, in *protos.KickMsg, opts ...grpc.CallOption) (*protos.KickAnswer, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "KickUser", varargs...)
	ret0, _ := ret[0].(*protos.KickAnswer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KickUser indicates an expected call of KickUser.
func (mr *MockPitayaClientMockRecorder) KickUser(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KickUser", reflect.TypeOf((*MockPitayaClient)(nil).KickUser), varargs...)
}

// NatsCallInSingleRoutine mocks base method.
func (m *MockPitayaClient) NatsCallInSingleRoutine(ctx context.Context, in *protos.Request, opts ...grpc.CallOption) (*protos.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NatsCallInSingleRoutine", varargs...)
	ret0, _ := ret[0].(*protos.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NatsCallInSingleRoutine indicates an expected call of NatsCallInSingleRoutine.
func (mr *MockPitayaClientMockRecorder) NatsCallInSingleRoutine(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NatsCallInSingleRoutine", reflect.TypeOf((*MockPitayaClient)(nil).NatsCallInSingleRoutine), varargs...)
}

// PushToUser mocks base method.
func (m *MockPitayaClient) PushToUser(ctx context.Context, in *protos.Push, opts ...grpc.CallOption) (*protos.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PushToUser", varargs...)
	ret0, _ := ret[0].(*protos.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PushToUser indicates an expected call of PushToUser.
func (mr *MockPitayaClientMockRecorder) PushToUser(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushToUser", reflect.TypeOf((*MockPitayaClient)(nil).PushToUser), varargs...)
}

// SessionBindRemote mocks base method.
func (m *MockPitayaClient) SessionBindRemote(ctx context.Context, in *protos.BindMsg, opts ...grpc.CallOption) (*protos.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SessionBindRemote", varargs...)
	ret0, _ := ret[0].(*protos.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SessionBindRemote indicates an expected call of SessionBindRemote.
func (mr *MockPitayaClientMockRecorder) SessionBindRemote(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SessionBindRemote", reflect.TypeOf((*MockPitayaClient)(nil).SessionBindRemote), varargs...)
}

// MockPitayaServer is a mock of PitayaServer interface.
type MockPitayaServer struct {
	ctrl     *gomock.Controller
	recorder *MockPitayaServerMockRecorder
}

// MockPitayaServerMockRecorder is the mock recorder for MockPitayaServer.
type MockPitayaServerMockRecorder struct {
	mock *MockPitayaServer
}

// NewMockPitayaServer creates a new mock instance.
func NewMockPitayaServer(ctrl *gomock.Controller) *MockPitayaServer {
	mock := &MockPitayaServer{ctrl: ctrl}
	mock.recorder = &MockPitayaServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPitayaServer) EXPECT() *MockPitayaServerMockRecorder {
	return m.recorder
}

// Call mocks base method.
func (m *MockPitayaServer) Call(arg0 context.Context, arg1 *protos.Request) (*protos.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call", arg0, arg1)
	ret0, _ := ret[0].(*protos.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Call indicates an expected call of Call.
func (mr *MockPitayaServerMockRecorder) Call(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockPitayaServer)(nil).Call), arg0, arg1)
}

// KickUser mocks base method.
func (m *MockPitayaServer) KickUser(arg0 context.Context, arg1 *protos.KickMsg) (*protos.KickAnswer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KickUser", arg0, arg1)
	ret0, _ := ret[0].(*protos.KickAnswer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KickUser indicates an expected call of KickUser.
func (mr *MockPitayaServerMockRecorder) KickUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KickUser", reflect.TypeOf((*MockPitayaServer)(nil).KickUser), arg0, arg1)
}

// NatsCallInSingleRoutine mocks base method.
func (m *MockPitayaServer) NatsCallInSingleRoutine(arg0 context.Context, arg1 *protos.Request) (*protos.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NatsCallInSingleRoutine", arg0, arg1)
	ret0, _ := ret[0].(*protos.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NatsCallInSingleRoutine indicates an expected call of NatsCallInSingleRoutine.
func (mr *MockPitayaServerMockRecorder) NatsCallInSingleRoutine(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NatsCallInSingleRoutine", reflect.TypeOf((*MockPitayaServer)(nil).NatsCallInSingleRoutine), arg0, arg1)
}

// PushToUser mocks base method.
func (m *MockPitayaServer) PushToUser(arg0 context.Context, arg1 *protos.Push) (*protos.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushToUser", arg0, arg1)
	ret0, _ := ret[0].(*protos.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PushToUser indicates an expected call of PushToUser.
func (mr *MockPitayaServerMockRecorder) PushToUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushToUser", reflect.TypeOf((*MockPitayaServer)(nil).PushToUser), arg0, arg1)
}

// SessionBindRemote mocks base method.
func (m *MockPitayaServer) SessionBindRemote(arg0 context.Context, arg1 *protos.BindMsg) (*protos.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SessionBindRemote", arg0, arg1)
	ret0, _ := ret[0].(*protos.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SessionBindRemote indicates an expected call of SessionBindRemote.
func (mr *MockPitayaServerMockRecorder) SessionBindRemote(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SessionBindRemote", reflect.TypeOf((*MockPitayaServer)(nil).SessionBindRemote), arg0, arg1)
}

// mustEmbedUnimplementedPitayaServer mocks base method.
func (m *MockPitayaServer) mustEmbedUnimplementedPitayaServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedPitayaServer")
}

// MockUnsafePitayaServer is a mock of UnsafePitayaServer interface.
type MockUnsafePitayaServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafePitayaServerMockRecorder
}

// MockUnsafePitayaServerMockRecorder is the mock recorder for MockUnsafePitayaServer.
type MockUnsafePitayaServerMockRecorder struct {
	mock *MockUnsafePitayaServer
}

// NewMockUnsafePitayaServer creates a new mock instance.
func NewMockUnsafePitayaServer(ctrl *gomock.Controller) *MockUnsafePitayaServer {
	mock := &MockUnsafePitayaServer{ctrl: ctrl}
	mock.recorder = &MockUnsafePitayaServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafePitayaServer) EXPECT() *MockUnsafePitayaServerMockRecorder {
	return m.recorder
}