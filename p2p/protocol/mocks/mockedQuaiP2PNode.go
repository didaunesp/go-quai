// Code generated by MockGen. DO NOT EDIT.
// Source: p2p/protocol/interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	network "github.com/libp2p/go-libp2p/core/network"
	peer "github.com/libp2p/go-libp2p/core/peer"
	protocol "github.com/libp2p/go-libp2p/core/protocol"
)

// MockQuaiP2PNode is a mock of QuaiP2PNode interface.
type MockQuaiP2PNode struct {
	ctrl     *gomock.Controller
	recorder *MockQuaiP2PNodeMockRecorder
}

// MockQuaiP2PNodeMockRecorder is the mock recorder for MockQuaiP2PNode.
type MockQuaiP2PNodeMockRecorder struct {
	mock *MockQuaiP2PNode
}

// NewMockQuaiP2PNode creates a new mock instance.
func NewMockQuaiP2PNode(ctrl *gomock.Controller) *MockQuaiP2PNode {
	mock := &MockQuaiP2PNode{ctrl: ctrl}
	mock.recorder = &MockQuaiP2PNodeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuaiP2PNode) EXPECT() *MockQuaiP2PNodeMockRecorder {
	return m.recorder
}

// Connect mocks base method.
func (m *MockQuaiP2PNode) Connect(pi peer.AddrInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", pi)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect.
func (mr *MockQuaiP2PNodeMockRecorder) Connect(pi interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockQuaiP2PNode)(nil).Connect), pi)
}

// GetBootPeers mocks base method.
func (m *MockQuaiP2PNode) GetBootPeers() []peer.AddrInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBootPeers")
	ret0, _ := ret[0].([]peer.AddrInfo)
	return ret0
}

// GetBootPeers indicates an expected call of GetBootPeers.
func (mr *MockQuaiP2PNodeMockRecorder) GetBootPeers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBootPeers", reflect.TypeOf((*MockQuaiP2PNode)(nil).GetBootPeers))
}

// Network mocks base method.
func (m *MockQuaiP2PNode) Network() network.Network {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Network")
	ret0, _ := ret[0].(network.Network)
	return ret0
}

// Network indicates an expected call of Network.
func (mr *MockQuaiP2PNodeMockRecorder) Network() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Network", reflect.TypeOf((*MockQuaiP2PNode)(nil).Network))
}

// NewStream mocks base method.
func (m *MockQuaiP2PNode) NewStream(peerID peer.ID, protocolID protocol.ID) (network.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewStream", peerID, protocolID)
	ret0, _ := ret[0].(network.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewStream indicates an expected call of NewStream.
func (mr *MockQuaiP2PNodeMockRecorder) NewStream(peerID, protocolID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewStream", reflect.TypeOf((*MockQuaiP2PNode)(nil).NewStream), peerID, protocolID)
}