// The MIT License (MIT)
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: timerSequence.go

// Package history is a generated GoMock package.
package history

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MocktimerSequence is a mock of timerSequence interface
type MocktimerSequence struct {
	ctrl     *gomock.Controller
	recorder *MocktimerSequenceMockRecorder
}

// MocktimerSequenceMockRecorder is the mock recorder for MocktimerSequence
type MocktimerSequenceMockRecorder struct {
	mock *MocktimerSequence
}

// NewMocktimerSequence creates a new mock instance
func NewMocktimerSequence(ctrl *gomock.Controller) *MocktimerSequence {
	mock := &MocktimerSequence{ctrl: ctrl}
	mock.recorder = &MocktimerSequenceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MocktimerSequence) EXPECT() *MocktimerSequenceMockRecorder {
	return m.recorder
}

// isExpired mocks base method
func (m *MocktimerSequence) isExpired(referenceTime time.Time, timerSequenceID timerSequenceID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "isExpired", referenceTime, timerSequenceID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// isExpired indicates an expected call of isExpired
func (mr *MocktimerSequenceMockRecorder) isExpired(referenceTime, timerSequenceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isExpired", reflect.TypeOf((*MocktimerSequence)(nil).isExpired), referenceTime, timerSequenceID)
}

// createNextUserTimer mocks base method
func (m *MocktimerSequence) createNextUserTimer() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "createNextUserTimer")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// createNextUserTimer indicates an expected call of createNextUserTimer
func (mr *MocktimerSequenceMockRecorder) createNextUserTimer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "createNextUserTimer", reflect.TypeOf((*MocktimerSequence)(nil).createNextUserTimer))
}

// createNextActivityTimer mocks base method
func (m *MocktimerSequence) createNextActivityTimer() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "createNextActivityTimer")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// createNextActivityTimer indicates an expected call of createNextActivityTimer
func (mr *MocktimerSequenceMockRecorder) createNextActivityTimer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "createNextActivityTimer", reflect.TypeOf((*MocktimerSequence)(nil).createNextActivityTimer))
}

// loadAndSortUserTimers mocks base method
func (m *MocktimerSequence) loadAndSortUserTimers() []timerSequenceID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "loadAndSortUserTimers")
	ret0, _ := ret[0].([]timerSequenceID)
	return ret0
}

// loadAndSortUserTimers indicates an expected call of loadAndSortUserTimers
func (mr *MocktimerSequenceMockRecorder) loadAndSortUserTimers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "loadAndSortUserTimers", reflect.TypeOf((*MocktimerSequence)(nil).loadAndSortUserTimers))
}

// loadAndSortActivityTimers mocks base method
func (m *MocktimerSequence) loadAndSortActivityTimers() []timerSequenceID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "loadAndSortActivityTimers")
	ret0, _ := ret[0].([]timerSequenceID)
	return ret0
}

// loadAndSortActivityTimers indicates an expected call of loadAndSortActivityTimers
func (mr *MocktimerSequenceMockRecorder) loadAndSortActivityTimers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "loadAndSortActivityTimers", reflect.TypeOf((*MocktimerSequence)(nil).loadAndSortActivityTimers))
}
