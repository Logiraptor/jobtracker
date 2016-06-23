// Automatically generated by MockGen. DO NOT EDIT!
// Source: app.go

package mocks

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of Logger interface
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *_MockLoggerRecorder
}

// Recorder for MockLogger (not exported)
type _MockLoggerRecorder struct {
	mock *MockLogger
}

func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &_MockLoggerRecorder{mock}
	return mock
}

func (_m *MockLogger) EXPECT() *_MockLoggerRecorder {
	return _m.recorder
}

func (_m *MockLogger) Log(format string, args ...interface{}) {
	_s := []interface{}{format}
	for _, _x := range args {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Log", _s...)
}

func (_mr *_MockLoggerRecorder) Log(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Log", _s...)
}

// Mock of Pather interface
type MockPather struct {
	ctrl     *gomock.Controller
	recorder *_MockPatherRecorder
}

// Recorder for MockPather (not exported)
type _MockPatherRecorder struct {
	mock *MockPather
}

func NewMockPather(ctrl *gomock.Controller) *MockPather {
	mock := &MockPather{ctrl: ctrl}
	mock.recorder = &_MockPatherRecorder{mock}
	return mock
}

func (_m *MockPather) EXPECT() *_MockPatherRecorder {
	return _m.recorder
}

func (_m *MockPather) Path(name string, args ...string) string {
	_s := []interface{}{name}
	for _, _x := range args {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Path", _s...)
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockPatherRecorder) Path(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Path", _s...)
}
