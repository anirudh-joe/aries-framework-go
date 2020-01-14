// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hyperledger/aries-framework-go/pkg/didcomm/common/service (interfaces: DIDComm)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	service "github.com/hyperledger/aries-framework-go/pkg/didcomm/common/service"
	reflect "reflect"
)

// MockDIDComm is a mock of DIDComm interface
type MockDIDComm struct {
	ctrl     *gomock.Controller
	recorder *MockDIDCommMockRecorder
}

// MockDIDCommMockRecorder is the mock recorder for MockDIDComm
type MockDIDCommMockRecorder struct {
	mock *MockDIDComm
}

// NewMockDIDComm creates a new mock instance
func NewMockDIDComm(ctrl *gomock.Controller) *MockDIDComm {
	mock := &MockDIDComm{ctrl: ctrl}
	mock.recorder = &MockDIDCommMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDIDComm) EXPECT() *MockDIDCommMockRecorder {
	return m.recorder
}

// HandleInbound mocks base method
func (m *MockDIDComm) HandleInbound(arg0 *service.DIDCommMsg, arg1, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleInbound", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleInbound indicates an expected call of HandleInbound
func (mr *MockDIDCommMockRecorder) HandleInbound(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleInbound", reflect.TypeOf((*MockDIDComm)(nil).HandleInbound), arg0, arg1, arg2)
}

// HandleOutbound mocks base method
func (m *MockDIDComm) HandleOutbound(arg0 *service.DIDCommMsg, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleOutbound", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleOutbound indicates an expected call of HandleOutbound
func (mr *MockDIDCommMockRecorder) HandleOutbound(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleOutbound", reflect.TypeOf((*MockDIDComm)(nil).HandleOutbound), arg0, arg1, arg2)
}

// RegisterActionEvent mocks base method
func (m *MockDIDComm) RegisterActionEvent(arg0 chan<- service.DIDCommAction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterActionEvent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterActionEvent indicates an expected call of RegisterActionEvent
func (mr *MockDIDCommMockRecorder) RegisterActionEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterActionEvent", reflect.TypeOf((*MockDIDComm)(nil).RegisterActionEvent), arg0)
}

// RegisterMsgEvent mocks base method
func (m *MockDIDComm) RegisterMsgEvent(arg0 chan<- service.StateMsg) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterMsgEvent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterMsgEvent indicates an expected call of RegisterMsgEvent
func (mr *MockDIDCommMockRecorder) RegisterMsgEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterMsgEvent", reflect.TypeOf((*MockDIDComm)(nil).RegisterMsgEvent), arg0)
}

// UnregisterActionEvent mocks base method
func (m *MockDIDComm) UnregisterActionEvent(arg0 chan<- service.DIDCommAction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnregisterActionEvent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnregisterActionEvent indicates an expected call of UnregisterActionEvent
func (mr *MockDIDCommMockRecorder) UnregisterActionEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnregisterActionEvent", reflect.TypeOf((*MockDIDComm)(nil).UnregisterActionEvent), arg0)
}

// UnregisterMsgEvent mocks base method
func (m *MockDIDComm) UnregisterMsgEvent(arg0 chan<- service.StateMsg) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnregisterMsgEvent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnregisterMsgEvent indicates an expected call of UnregisterMsgEvent
func (mr *MockDIDCommMockRecorder) UnregisterMsgEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnregisterMsgEvent", reflect.TypeOf((*MockDIDComm)(nil).UnregisterMsgEvent), arg0)
}
