// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/libp2p/go-libp2p/core/network (interfaces: ConnManagementScope)

// Package mocknetwork is a generated GoMock package.
package mocknetwork

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	network "github.com/multiversx/go-libp2p/core/network"
	peer "github.com/multiversx/go-libp2p/core/peer"
)

// MockConnManagementScope is a mock of ConnManagementScope interface.
type MockConnManagementScope struct {
	ctrl     *gomock.Controller
	recorder *MockConnManagementScopeMockRecorder
}

// MockConnManagementScopeMockRecorder is the mock recorder for MockConnManagementScope.
type MockConnManagementScopeMockRecorder struct {
	mock *MockConnManagementScope
}

// NewMockConnManagementScope creates a new mock instance.
func NewMockConnManagementScope(ctrl *gomock.Controller) *MockConnManagementScope {
	mock := &MockConnManagementScope{ctrl: ctrl}
	mock.recorder = &MockConnManagementScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnManagementScope) EXPECT() *MockConnManagementScopeMockRecorder {
	return m.recorder
}

// BeginSpan mocks base method.
func (m *MockConnManagementScope) BeginSpan() (network.ResourceScopeSpan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginSpan")
	ret0, _ := ret[0].(network.ResourceScopeSpan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginSpan indicates an expected call of BeginSpan.
func (mr *MockConnManagementScopeMockRecorder) BeginSpan() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginSpan", reflect.TypeOf((*MockConnManagementScope)(nil).BeginSpan))
}

// Done mocks base method.
func (m *MockConnManagementScope) Done() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Done")
}

// Done indicates an expected call of Done.
func (mr *MockConnManagementScopeMockRecorder) Done() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Done", reflect.TypeOf((*MockConnManagementScope)(nil).Done))
}

// PeerScope mocks base method.
func (m *MockConnManagementScope) PeerScope() network.PeerScope {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeerScope")
	ret0, _ := ret[0].(network.PeerScope)
	return ret0
}

// PeerScope indicates an expected call of PeerScope.
func (mr *MockConnManagementScopeMockRecorder) PeerScope() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeerScope", reflect.TypeOf((*MockConnManagementScope)(nil).PeerScope))
}

// ReleaseMemory mocks base method.
func (m *MockConnManagementScope) ReleaseMemory(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReleaseMemory", arg0)
}

// ReleaseMemory indicates an expected call of ReleaseMemory.
func (mr *MockConnManagementScopeMockRecorder) ReleaseMemory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseMemory", reflect.TypeOf((*MockConnManagementScope)(nil).ReleaseMemory), arg0)
}

// ReserveMemory mocks base method.
func (m *MockConnManagementScope) ReserveMemory(arg0 int, arg1 byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReserveMemory", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReserveMemory indicates an expected call of ReserveMemory.
func (mr *MockConnManagementScopeMockRecorder) ReserveMemory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReserveMemory", reflect.TypeOf((*MockConnManagementScope)(nil).ReserveMemory), arg0, arg1)
}

// SetPeer mocks base method.
func (m *MockConnManagementScope) SetPeer(arg0 peer.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPeer", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPeer indicates an expected call of SetPeer.
func (mr *MockConnManagementScopeMockRecorder) SetPeer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPeer", reflect.TypeOf((*MockConnManagementScope)(nil).SetPeer), arg0)
}

// Stat mocks base method.
func (m *MockConnManagementScope) Stat() network.ScopeStat {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stat")
	ret0, _ := ret[0].(network.ScopeStat)
	return ret0
}

// Stat indicates an expected call of Stat.
func (mr *MockConnManagementScopeMockRecorder) Stat() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockConnManagementScope)(nil).Stat))
}
