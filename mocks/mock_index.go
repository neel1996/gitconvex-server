// Code generated by MockGen. DO NOT EDIT.
// Source: ../git/middleware/index.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	git "github.com/libgit2/git2go/v31"
)

// MockIndex is a mock of Index interface.
type MockIndex struct {
	ctrl     *gomock.Controller
	recorder *MockIndexMockRecorder
}

// MockIndexMockRecorder is the mock recorder for MockIndex.
type MockIndexMockRecorder struct {
	mock *MockIndex
}

// NewMockIndex creates a new mock instance.
func NewMockIndex(ctrl *gomock.Controller) *MockIndex {
	mock := &MockIndex{ctrl: ctrl}
	mock.recorder = &MockIndexMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIndex) EXPECT() *MockIndexMockRecorder {
	return m.recorder
}

// WriteTree mocks base method.
func (m *MockIndex) WriteTree() (*git.Oid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteTree")
	ret0, _ := ret[0].(*git.Oid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteTree indicates an expected call of WriteTree.
func (mr *MockIndexMockRecorder) WriteTree() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteTree", reflect.TypeOf((*MockIndex)(nil).WriteTree))
}