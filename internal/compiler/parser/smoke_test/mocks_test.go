// Code generated by MockGen. DO NOT EDIT.
// Source: smoke_test.go

// Package smoke_test is a generated GoMock package.
package smoke_test

import (
	reflect "reflect"

	antlr "github.com/antlr4-go/antlr/v4"
	gomock "github.com/golang/mock/gomock"
)

// MockMyErrorListener is a mock of MyErrorListener interface.
type MockMyErrorListener struct {
	ctrl     *gomock.Controller
	recorder *MockMyErrorListenerMockRecorder
}

// MockMyErrorListenerMockRecorder is the mock recorder for MockMyErrorListener.
type MockMyErrorListenerMockRecorder struct {
	mock *MockMyErrorListener
}

// NewMockMyErrorListener creates a new mock instance.
func NewMockMyErrorListener(ctrl *gomock.Controller) *MockMyErrorListener {
	mock := &MockMyErrorListener{ctrl: ctrl}
	mock.recorder = &MockMyErrorListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMyErrorListener) EXPECT() *MockMyErrorListenerMockRecorder {
	return m.recorder
}

// ReportAmbiguity mocks base method.
func (m *MockMyErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReportAmbiguity", recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
}

// ReportAmbiguity indicates an expected call of ReportAmbiguity.
func (mr *MockMyErrorListenerMockRecorder) ReportAmbiguity(recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportAmbiguity", reflect.TypeOf((*MockMyErrorListener)(nil).ReportAmbiguity), recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
}

// ReportAttemptingFullContext mocks base method.
func (m *MockMyErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReportAttemptingFullContext", recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
}

// ReportAttemptingFullContext indicates an expected call of ReportAttemptingFullContext.
func (mr *MockMyErrorListenerMockRecorder) ReportAttemptingFullContext(recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportAttemptingFullContext", reflect.TypeOf((*MockMyErrorListener)(nil).ReportAttemptingFullContext), recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
}

// ReportContextSensitivity mocks base method.
func (m *MockMyErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReportContextSensitivity", recognizer, dfa, startIndex, stopIndex, prediction, configs)
}

// ReportContextSensitivity indicates an expected call of ReportContextSensitivity.
func (mr *MockMyErrorListenerMockRecorder) ReportContextSensitivity(recognizer, dfa, startIndex, stopIndex, prediction, configs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportContextSensitivity", reflect.TypeOf((*MockMyErrorListener)(nil).ReportContextSensitivity), recognizer, dfa, startIndex, stopIndex, prediction, configs)
}

// SyntaxError mocks base method.
func (m *MockMyErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SyntaxError", recognizer, offendingSymbol, line, column, msg, e)
}

// SyntaxError indicates an expected call of SyntaxError.
func (mr *MockMyErrorListenerMockRecorder) SyntaxError(recognizer, offendingSymbol, line, column, msg, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyntaxError", reflect.TypeOf((*MockMyErrorListener)(nil).SyntaxError), recognizer, offendingSymbol, line, column, msg, e)
}
