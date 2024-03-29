// Code generated by MockGen. DO NOT EDIT.
// Source: validator/validator.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockValidator is a mock of Validator interface.
type MockValidator struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorMockRecorder
}

// MockValidatorMockRecorder is the mock recorder for MockValidator.
type MockValidatorMockRecorder struct {
	mock *MockValidator
}

// NewMockValidator creates a new mock instance.
func NewMockValidator(ctrl *gomock.Controller) *MockValidator {
	mock := &MockValidator{ctrl: ctrl}
	mock.recorder = &MockValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidator) EXPECT() *MockValidatorMockRecorder {
	return m.recorder
}

// Validate mocks base method.
func (m *MockValidator) Validate(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockValidatorMockRecorder) Validate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockValidator)(nil).Validate), arg0)
}

// MockValidatorWithStringFields is a mock of ValidatorWithStringFields interface.
type MockValidatorWithStringFields struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorWithStringFieldsMockRecorder
}

// MockValidatorWithStringFieldsMockRecorder is the mock recorder for MockValidatorWithStringFields.
type MockValidatorWithStringFieldsMockRecorder struct {
	mock *MockValidatorWithStringFields
}

// NewMockValidatorWithStringFields creates a new mock instance.
func NewMockValidatorWithStringFields(ctrl *gomock.Controller) *MockValidatorWithStringFields {
	mock := &MockValidatorWithStringFields{ctrl: ctrl}
	mock.recorder = &MockValidatorWithStringFieldsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidatorWithStringFields) EXPECT() *MockValidatorWithStringFieldsMockRecorder {
	return m.recorder
}

// ValidateWithFields mocks base method.
func (m *MockValidatorWithStringFields) ValidateWithFields(fields ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ValidateWithFields", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateWithFields indicates an expected call of ValidateWithFields.
func (mr *MockValidatorWithStringFieldsMockRecorder) ValidateWithFields(fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateWithFields", reflect.TypeOf((*MockValidatorWithStringFields)(nil).ValidateWithFields), fields...)
}
