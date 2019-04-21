// Code generated by MockGen. DO NOT EDIT.
// Source: os.go

// Package mock_oswrap is a generated GoMock package.
package mock_oswrap

import (
	gomock "github.com/golang/mock/gomock"
	os "os"
	exec "os/exec"
	reflect "reflect"
)

// MockOsWrap is a mock of OsWrap interface
type MockOsWrap struct {
	ctrl     *gomock.Controller
	recorder *MockOsWrapMockRecorder
}

// MockOsWrapMockRecorder is the mock recorder for MockOsWrap
type MockOsWrapMockRecorder struct {
	mock *MockOsWrap
}

// NewMockOsWrap creates a new mock instance
func NewMockOsWrap(ctrl *gomock.Controller) *MockOsWrap {
	mock := &MockOsWrap{ctrl: ctrl}
	mock.recorder = &MockOsWrapMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOsWrap) EXPECT() *MockOsWrapMockRecorder {
	return m.recorder
}

// Command mocks base method
func (m *MockOsWrap) Command(name string, arg ...string) *exec.Cmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{name}
	for _, a := range arg {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Command", varargs...)
	ret0, _ := ret[0].(*exec.Cmd)
	return ret0
}

// Command indicates an expected call of Command
func (mr *MockOsWrapMockRecorder) Command(name interface{}, arg ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name}, arg...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Command", reflect.TypeOf((*MockOsWrap)(nil).Command), varargs...)
}

// Exec mocks base method
func (m *MockOsWrap) Exec(argv0 string, argv, envv []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exec", argv0, argv, envv)
	ret0, _ := ret[0].(error)
	return ret0
}

// Exec indicates an expected call of Exec
func (mr *MockOsWrapMockRecorder) Exec(argv0, argv, envv interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockOsWrap)(nil).Exec), argv0, argv, envv)
}

// Create mocks base method
func (m *MockOsWrap) Create(name string) (*os.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", name)
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockOsWrapMockRecorder) Create(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOsWrap)(nil).Create), name)
}

// LookPath mocks base method
func (m *MockOsWrap) LookPath(name string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LookPath", name)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LookPath indicates an expected call of LookPath
func (mr *MockOsWrapMockRecorder) LookPath(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookPath", reflect.TypeOf((*MockOsWrap)(nil).LookPath), name)
}

// Environ mocks base method
func (m *MockOsWrap) Environ() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Environ")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Environ indicates an expected call of Environ
func (mr *MockOsWrapMockRecorder) Environ() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Environ", reflect.TypeOf((*MockOsWrap)(nil).Environ))
}

// Remove mocks base method
func (m *MockOsWrap) Remove(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockOsWrapMockRecorder) Remove(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockOsWrap)(nil).Remove), name)
}
