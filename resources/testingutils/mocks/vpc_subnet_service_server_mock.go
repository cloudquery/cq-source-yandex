// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1 (interfaces: SubnetServiceServer)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	vpc "github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"
)

// MockSubnetServiceServer is a mock of SubnetServiceServer interface.
type MockSubnetServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockSubnetServiceServerMockRecorder
}

// MockSubnetServiceServerMockRecorder is the mock recorder for MockSubnetServiceServer.
type MockSubnetServiceServerMockRecorder struct {
	mock *MockSubnetServiceServer
}

// NewMockSubnetServiceServer creates a new mock instance.
func NewMockSubnetServiceServer(ctrl *gomock.Controller) *MockSubnetServiceServer {
	mock := &MockSubnetServiceServer{ctrl: ctrl}
	mock.recorder = &MockSubnetServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubnetServiceServer) EXPECT() *MockSubnetServiceServerMockRecorder {
	return m.recorder
}

// AddCidrBlocks mocks base method.
func (m *MockSubnetServiceServer) AddCidrBlocks(arg0 context.Context, arg1 *vpc.AddSubnetCidrBlocksRequest) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCidrBlocks", arg0, arg1)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCidrBlocks indicates an expected call of AddCidrBlocks.
func (mr *MockSubnetServiceServerMockRecorder) AddCidrBlocks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCidrBlocks", reflect.TypeOf((*MockSubnetServiceServer)(nil).AddCidrBlocks), arg0, arg1)
}

// Create mocks base method.
func (m *MockSubnetServiceServer) Create(arg0 context.Context, arg1 *vpc.CreateSubnetRequest) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockSubnetServiceServerMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSubnetServiceServer)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockSubnetServiceServer) Delete(arg0 context.Context, arg1 *vpc.DeleteSubnetRequest) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockSubnetServiceServerMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSubnetServiceServer)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockSubnetServiceServer) Get(arg0 context.Context, arg1 *vpc.GetSubnetRequest) (*vpc.Subnet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*vpc.Subnet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSubnetServiceServerMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSubnetServiceServer)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockSubnetServiceServer) List(arg0 context.Context, arg1 *vpc.ListSubnetsRequest) (*vpc.ListSubnetsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*vpc.ListSubnetsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockSubnetServiceServerMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSubnetServiceServer)(nil).List), arg0, arg1)
}

// ListOperations mocks base method.
func (m *MockSubnetServiceServer) ListOperations(arg0 context.Context, arg1 *vpc.ListSubnetOperationsRequest) (*vpc.ListSubnetOperationsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOperations", arg0, arg1)
	ret0, _ := ret[0].(*vpc.ListSubnetOperationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOperations indicates an expected call of ListOperations.
func (mr *MockSubnetServiceServerMockRecorder) ListOperations(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOperations", reflect.TypeOf((*MockSubnetServiceServer)(nil).ListOperations), arg0, arg1)
}

// ListUsedAddresses mocks base method.
func (m *MockSubnetServiceServer) ListUsedAddresses(arg0 context.Context, arg1 *vpc.ListUsedAddressesRequest) (*vpc.ListUsedAddressesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsedAddresses", arg0, arg1)
	ret0, _ := ret[0].(*vpc.ListUsedAddressesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsedAddresses indicates an expected call of ListUsedAddresses.
func (mr *MockSubnetServiceServerMockRecorder) ListUsedAddresses(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsedAddresses", reflect.TypeOf((*MockSubnetServiceServer)(nil).ListUsedAddresses), arg0, arg1)
}

// Move mocks base method.
func (m *MockSubnetServiceServer) Move(arg0 context.Context, arg1 *vpc.MoveSubnetRequest) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Move", arg0, arg1)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Move indicates an expected call of Move.
func (mr *MockSubnetServiceServerMockRecorder) Move(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Move", reflect.TypeOf((*MockSubnetServiceServer)(nil).Move), arg0, arg1)
}

// RemoveCidrBlocks mocks base method.
func (m *MockSubnetServiceServer) RemoveCidrBlocks(arg0 context.Context, arg1 *vpc.RemoveSubnetCidrBlocksRequest) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveCidrBlocks", arg0, arg1)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveCidrBlocks indicates an expected call of RemoveCidrBlocks.
func (mr *MockSubnetServiceServerMockRecorder) RemoveCidrBlocks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveCidrBlocks", reflect.TypeOf((*MockSubnetServiceServer)(nil).RemoveCidrBlocks), arg0, arg1)
}

// Update mocks base method.
func (m *MockSubnetServiceServer) Update(arg0 context.Context, arg1 *vpc.UpdateSubnetRequest) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockSubnetServiceServerMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSubnetServiceServer)(nil).Update), arg0, arg1)
}
