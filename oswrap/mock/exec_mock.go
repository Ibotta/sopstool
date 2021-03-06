// Code generated by MockGen. DO NOT EDIT.
// Source: exec.go

// Package mock_oswrap is a generated GoMock package.
package mock_oswrap

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockExecWrap is a mock of ExecWrap interface
type MockExecWrap struct {
	ctrl     *gomock.Controller
	recorder *MockExecWrapMockRecorder
}

// MockExecWrapMockRecorder is the mock recorder for MockExecWrap
type MockExecWrapMockRecorder struct {
	mock *MockExecWrap
}

// NewMockExecWrap creates a new mock instance
func NewMockExecWrap(ctrl *gomock.Controller) *MockExecWrap {
	mock := &MockExecWrap{ctrl: ctrl}
	mock.recorder = &MockExecWrapMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExecWrap) EXPECT() *MockExecWrapMockRecorder {
	return m.recorder
}

// RunCommandDirect mocks base method
func (m *MockExecWrap) RunCommandDirect(command []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunCommandDirect", command)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunCommandDirect indicates an expected call of RunCommandDirect
func (mr *MockExecWrapMockRecorder) RunCommandDirect(command interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunCommandDirect", reflect.TypeOf((*MockExecWrap)(nil).RunCommandDirect), command)
}

// RunCommandStdoutToFile mocks base method
func (m *MockExecWrap) RunCommandStdoutToFile(outfileName string, command []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunCommandStdoutToFile", outfileName, command)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunCommandStdoutToFile indicates an expected call of RunCommandStdoutToFile
func (mr *MockExecWrapMockRecorder) RunCommandStdoutToFile(outfileName, command interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunCommandStdoutToFile", reflect.TypeOf((*MockExecWrap)(nil).RunCommandStdoutToFile), outfileName, command)
}

// RunSyscallExec mocks base method
func (m *MockExecWrap) RunSyscallExec(args []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunSyscallExec", args)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunSyscallExec indicates an expected call of RunSyscallExec
func (mr *MockExecWrapMockRecorder) RunSyscallExec(args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunSyscallExec", reflect.TypeOf((*MockExecWrap)(nil).RunSyscallExec), args)
}
