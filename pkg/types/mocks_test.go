// Code generated by MockGen. DO NOT EDIT.
// Source: resolver.go

// Package types_test is a generated GoMock package.
package types_test

import (
	reflect "reflect"

	types "github.com/nevalang/neva/pkg/types"
	gomock "github.com/golang/mock/gomock"
)

// MockexprValidator is a mock of exprValidator interface.
type MockexprValidator struct {
	ctrl     *gomock.Controller
	recorder *MockexprValidatorMockRecorder
}

// MockexprValidatorMockRecorder is the mock recorder for MockexprValidator.
type MockexprValidatorMockRecorder struct {
	mock *MockexprValidator
}

// NewMockexprValidator creates a new mock instance.
func NewMockexprValidator(ctrl *gomock.Controller) *MockexprValidator {
	mock := &MockexprValidator{ctrl: ctrl}
	mock.recorder = &MockexprValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockexprValidator) EXPECT() *MockexprValidatorMockRecorder {
	return m.recorder
}

// Validate mocks base method.
func (m *MockexprValidator) Validate(arg0 types.Expr) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockexprValidatorMockRecorder) Validate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockexprValidator)(nil).Validate), arg0)
}

// ValidateDef mocks base method.
func (m *MockexprValidator) ValidateDef(def types.Def) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateDef", def)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateDef indicates an expected call of ValidateDef.
func (mr *MockexprValidatorMockRecorder) ValidateDef(def interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateDef", reflect.TypeOf((*MockexprValidator)(nil).ValidateDef), def)
}

// MockcompatChecker is a mock of compatChecker interface.
type MockcompatChecker struct {
	ctrl     *gomock.Controller
	recorder *MockcompatCheckerMockRecorder
}

// MockcompatCheckerMockRecorder is the mock recorder for MockcompatChecker.
type MockcompatCheckerMockRecorder struct {
	mock *MockcompatChecker
}

// NewMockcompatChecker creates a new mock instance.
func NewMockcompatChecker(ctrl *gomock.Controller) *MockcompatChecker {
	mock := &MockcompatChecker{ctrl: ctrl}
	mock.recorder = &MockcompatCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockcompatChecker) EXPECT() *MockcompatCheckerMockRecorder {
	return m.recorder
}

// Check mocks base method.
func (m *MockcompatChecker) Check(arg0, arg1 types.Expr, arg2 types.TerminatorParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Check indicates an expected call of Check.
func (mr *MockcompatCheckerMockRecorder) Check(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockcompatChecker)(nil).Check), arg0, arg1, arg2)
}

// MockrecursionTerminator is a mock of recursionTerminator interface.
type MockrecursionTerminator struct {
	ctrl     *gomock.Controller
	recorder *MockrecursionTerminatorMockRecorder
}

// MockrecursionTerminatorMockRecorder is the mock recorder for MockrecursionTerminator.
type MockrecursionTerminatorMockRecorder struct {
	mock *MockrecursionTerminator
}

// NewMockrecursionTerminator creates a new mock instance.
func NewMockrecursionTerminator(ctrl *gomock.Controller) *MockrecursionTerminator {
	mock := &MockrecursionTerminator{ctrl: ctrl}
	mock.recorder = &MockrecursionTerminatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockrecursionTerminator) EXPECT() *MockrecursionTerminatorMockRecorder {
	return m.recorder
}

// ShouldTerminate mocks base method.
func (m *MockrecursionTerminator) ShouldTerminate(arg0 types.Trace, arg1 types.Scope) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShouldTerminate", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShouldTerminate indicates an expected call of ShouldTerminate.
func (mr *MockrecursionTerminatorMockRecorder) ShouldTerminate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldTerminate", reflect.TypeOf((*MockrecursionTerminator)(nil).ShouldTerminate), arg0, arg1)
}

// MockScope is a mock of Scope interface.
type MockScope struct {
	ctrl     *gomock.Controller
	recorder *MockScopeMockRecorder
}

// MockScopeMockRecorder is the mock recorder for MockScope.
type MockScopeMockRecorder struct {
	mock *MockScope
}

// NewMockScope creates a new mock instance.
func NewMockScope(ctrl *gomock.Controller) *MockScope {
	mock := &MockScope{ctrl: ctrl}
	mock.recorder = &MockScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScope) EXPECT() *MockScopeMockRecorder {
	return m.recorder
}

// GetType mocks base method.
func (m *MockScope) GetType(ref string) (types.Def, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetType", ref)
	ret0, _ := ret[0].(types.Def)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetType indicates an expected call of GetType.
func (mr *MockScopeMockRecorder) GetType(ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetType", reflect.TypeOf((*MockScope)(nil).GetType), ref)
}

// Update mocks base method.
func (m *MockScope) Update(ref string) (types.Scope, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ref)
	ret0, _ := ret[0].(types.Scope)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockScopeMockRecorder) Update(ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockScope)(nil).Update), ref)
}
