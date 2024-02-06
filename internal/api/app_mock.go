// Code generated by MockGen. DO NOT EDIT.
// Source: api.go
//
// Generated by this command:
//
//	mockgen -source=api.go -destination=app_mock.go -package=api
//

// Package api is a generated GoMock package.
package api

import (
	reflect "reflect"

	fi "github.com/ggrrrr/fibonacci-svc/internal/fi"
	gomock "go.uber.org/mock/gomock"
)

// MockApp is a mock of App interface.
type MockApp struct {
	ctrl     *gomock.Controller
	recorder *MockAppMockRecorder
}

// MockAppMockRecorder is the mock recorder for MockApp.
type MockAppMockRecorder struct {
	mock *MockApp
}

// NewMockApp creates a new mock instance.
func NewMockApp(ctrl *gomock.Controller) *MockApp {
	mock := &MockApp{ctrl: ctrl}
	mock.recorder = &MockAppMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApp) EXPECT() *MockAppMockRecorder {
	return m.recorder
}

// Current mocks base method.
func (m *MockApp) Current() fi.Number {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Current")
	ret0, _ := ret[0].(fi.Number)
	return ret0
}

// Current indicates an expected call of Current.
func (mr *MockAppMockRecorder) Current() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Current", reflect.TypeOf((*MockApp)(nil).Current))
}

// Next mocks base method.
func (m *MockApp) Next() (fi.Number, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(fi.Number)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Next indicates an expected call of Next.
func (mr *MockAppMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockApp)(nil).Next))
}

// Previous mocks base method.
func (m *MockApp) Previous() (fi.Number, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Previous")
	ret0, _ := ret[0].(fi.Number)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Previous indicates an expected call of Previous.
func (mr *MockAppMockRecorder) Previous() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Previous", reflect.TypeOf((*MockApp)(nil).Previous))
}