// Code generated by MockGen. DO NOT EDIT.
// Source: ssh_wrapper.go

// Package wrapper is a generated GoMock package.
package wrapper

import (
	bytes "bytes"
	gomock "github.com/golang/mock/gomock"
	ssh "golang.org/x/crypto/ssh"
	reflect "reflect"
)

// MockSshInterface is a mock of SshInterface interface
type MockSshInterface struct {
	ctrl     *gomock.Controller
	recorder *MockSshInterfaceMockRecorder
}

// MockSshInterfaceMockRecorder is the mock recorder for MockSshInterface
type MockSshInterfaceMockRecorder struct {
	mock *MockSshInterface
}

// NewMockSshInterface creates a new mock instance
func NewMockSshInterface(ctrl *gomock.Controller) *MockSshInterface {
	mock := &MockSshInterface{ctrl: ctrl}
	mock.recorder = &MockSshInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSshInterface) EXPECT() *MockSshInterfaceMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockSshInterface) Run(addr, command string, resultStdout, resultStderr *bytes.Buffer, config *ssh.ClientConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", addr, command, resultStdout, resultStderr, config)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockSshInterfaceMockRecorder) Run(addr, command, resultStdout, resultStderr, config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockSshInterface)(nil).Run), addr, command, resultStdout, resultStderr, config)
}
