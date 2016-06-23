// Automatically generated by MockGen. DO NOT EDIT!
// Source: groups_test.go

package tests

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of TestInterface interface
type MockTestInterface struct {
	ctrl     *gomock.Controller
	recorder *_MockTestInterfaceRecorder
}

// Recorder for MockTestInterface (not exported)
type _MockTestInterfaceRecorder struct {
	mock *MockTestInterface
}

func NewMockTestInterface(ctrl *gomock.Controller) *MockTestInterface {
	mock := &MockTestInterface{ctrl: ctrl}
	mock.recorder = &_MockTestInterfaceRecorder{mock}
	return mock
}

func (_m *MockTestInterface) EXPECT() *_MockTestInterfaceRecorder {
	return _m.recorder
}

func (_m *MockTestInterface) Before(_param0 string) {
	_m.ctrl.Call(_m, "Before", _param0)
}

func (_mr *_MockTestInterfaceRecorder) Before(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Before", arg0)
}

func (_m *MockTestInterface) After(_param0 string) {
	_m.ctrl.Call(_m, "After", _param0)
}

func (_mr *_MockTestInterfaceRecorder) After(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "After", arg0)
}

func (_m *MockTestInterface) Body(_param0 string) {
	_m.ctrl.Call(_m, "Body", _param0)
}

func (_mr *_MockTestInterfaceRecorder) Body(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Body", arg0)
}
