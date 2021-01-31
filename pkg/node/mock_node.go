// Code generated by MockGen. DO NOT EDIT.
// Source: node.go

// This package is a generated GoMock package.
package node

import (
	bytes "bytes"
	gomock "github.com/golang/mock/gomock"
	wrapper "hypersds-provisioner/pkg/common/wrapper"
	reflect "reflect"
)

// MockNodeInterface is a mock of NodeInterface interface
type MockNodeInterface struct {
	ctrl     *gomock.Controller
	recorder *MockNodeInterfaceMockRecorder
}

// MockNodeInterfaceMockRecorder is the mock recorder for MockNodeInterface
type MockNodeInterfaceMockRecorder struct {
	mock *MockNodeInterface
}

// NewMockNodeInterface creates a new mock instance
func NewMockNodeInterface(ctrl *gomock.Controller) *MockNodeInterface {
	mock := &MockNodeInterface{ctrl: ctrl}
	mock.recorder = &MockNodeInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeInterface) EXPECT() *MockNodeInterfaceMockRecorder {
	return m.recorder
}

// SetUserId mocks base method
func (m *MockNodeInterface) SetUserId(userId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUserId", userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUserId indicates an expected call of SetUserId
func (mr *MockNodeInterfaceMockRecorder) SetUserId(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserId", reflect.TypeOf((*MockNodeInterface)(nil).SetUserId), userId)
}

// SetUserPw mocks base method
func (m *MockNodeInterface) SetUserPw(userPw string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUserPw", userPw)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUserPw indicates an expected call of SetUserPw
func (mr *MockNodeInterfaceMockRecorder) SetUserPw(userPw interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserPw", reflect.TypeOf((*MockNodeInterface)(nil).SetUserPw), userPw)
}

// SetHostSpec mocks base method
func (m *MockNodeInterface) SetHostSpec(hostSpec HostSpecInterface) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHostSpec", hostSpec)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHostSpec indicates an expected call of SetHostSpec
func (mr *MockNodeInterfaceMockRecorder) SetHostSpec(hostSpec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHostSpec", reflect.TypeOf((*MockNodeInterface)(nil).SetHostSpec), hostSpec)
}

// GetUserId mocks base method
func (m *MockNodeInterface) GetUserId() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserId")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserId indicates an expected call of GetUserId
func (mr *MockNodeInterfaceMockRecorder) GetUserId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserId", reflect.TypeOf((*MockNodeInterface)(nil).GetUserId))
}

// GetUserPw mocks base method
func (m *MockNodeInterface) GetUserPw() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserPw")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserPw indicates an expected call of GetUserPw
func (mr *MockNodeInterfaceMockRecorder) GetUserPw() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserPw", reflect.TypeOf((*MockNodeInterface)(nil).GetUserPw))
}

// GetHostSpec mocks base method
func (m *MockNodeInterface) GetHostSpec() (HostSpecInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostSpec")
	ret0, _ := ret[0].(HostSpecInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHostSpec indicates an expected call of GetHostSpec
func (mr *MockNodeInterfaceMockRecorder) GetHostSpec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostSpec", reflect.TypeOf((*MockNodeInterface)(nil).GetHostSpec))
}

// RunSshCmd mocks base method
func (m *MockNodeInterface) RunSshCmd(exec wrapper.ExecInterface, cmdQuery string) (bytes.Buffer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunSshCmd", exec, cmdQuery)
	ret0, _ := ret[0].(bytes.Buffer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunSshCmd indicates an expected call of RunSshCmd
func (mr *MockNodeInterfaceMockRecorder) RunSshCmd(exec, cmdQuery interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunSshCmd", reflect.TypeOf((*MockNodeInterface)(nil).RunSshCmd), exec, cmdQuery)
}