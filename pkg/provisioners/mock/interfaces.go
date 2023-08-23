// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/provisioners/interfaces.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	api "k8s.io/client-go/tools/clientcmd/api"
)

// MockRemoteCluster is a mock of RemoteCluster interface.
type MockRemoteCluster struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteClusterMockRecorder
}

// MockRemoteClusterMockRecorder is the mock recorder for MockRemoteCluster.
type MockRemoteClusterMockRecorder struct {
	mock *MockRemoteCluster
}

// NewMockRemoteCluster creates a new mock instance.
func NewMockRemoteCluster(ctrl *gomock.Controller) *MockRemoteCluster {
	mock := &MockRemoteCluster{ctrl: ctrl}
	mock.recorder = &MockRemoteClusterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRemoteCluster) EXPECT() *MockRemoteClusterMockRecorder {
	return m.recorder
}

// Config mocks base method.
func (m *MockRemoteCluster) Config(ctx context.Context) (*api.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config", ctx)
	ret0, _ := ret[0].(*api.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Config indicates an expected call of Config.
func (mr *MockRemoteClusterMockRecorder) Config(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockRemoteCluster)(nil).Config), ctx)
}

// Labels mocks base method.
func (m *MockRemoteCluster) Labels() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Labels")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Labels indicates an expected call of Labels.
func (mr *MockRemoteClusterMockRecorder) Labels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Labels", reflect.TypeOf((*MockRemoteCluster)(nil).Labels))
}

// Name mocks base method.
func (m *MockRemoteCluster) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockRemoteClusterMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockRemoteCluster)(nil).Name))
}

// Server mocks base method.
func (m *MockRemoteCluster) Server(ctx context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Server", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Server indicates an expected call of Server.
func (mr *MockRemoteClusterMockRecorder) Server(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Server", reflect.TypeOf((*MockRemoteCluster)(nil).Server), ctx)
}

// MockProvisioner is a mock of Provisioner interface.
type MockProvisioner struct {
	ctrl     *gomock.Controller
	recorder *MockProvisionerMockRecorder
}

// MockProvisionerMockRecorder is the mock recorder for MockProvisioner.
type MockProvisionerMockRecorder struct {
	mock *MockProvisioner
}

// NewMockProvisioner creates a new mock instance.
func NewMockProvisioner(ctrl *gomock.Controller) *MockProvisioner {
	mock := &MockProvisioner{ctrl: ctrl}
	mock.recorder = &MockProvisionerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProvisioner) EXPECT() *MockProvisionerMockRecorder {
	return m.recorder
}

// Deprovision mocks base method.
func (m *MockProvisioner) Deprovision(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deprovision", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deprovision indicates an expected call of Deprovision.
func (mr *MockProvisionerMockRecorder) Deprovision(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deprovision", reflect.TypeOf((*MockProvisioner)(nil).Deprovision), arg0)
}

// Provision mocks base method.
func (m *MockProvisioner) Provision(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Provision", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Provision indicates an expected call of Provision.
func (mr *MockProvisionerMockRecorder) Provision(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Provision", reflect.TypeOf((*MockProvisioner)(nil).Provision), arg0)
}

// ProvisionerName mocks base method.
func (m *MockProvisioner) ProvisionerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProvisionerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ProvisionerName indicates an expected call of ProvisionerName.
func (mr *MockProvisionerMockRecorder) ProvisionerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProvisionerName", reflect.TypeOf((*MockProvisioner)(nil).ProvisionerName))
}

// MockReadinessCheck is a mock of ReadinessCheck interface.
type MockReadinessCheck struct {
	ctrl     *gomock.Controller
	recorder *MockReadinessCheckMockRecorder
}

// MockReadinessCheckMockRecorder is the mock recorder for MockReadinessCheck.
type MockReadinessCheckMockRecorder struct {
	mock *MockReadinessCheck
}

// NewMockReadinessCheck creates a new mock instance.
func NewMockReadinessCheck(ctrl *gomock.Controller) *MockReadinessCheck {
	mock := &MockReadinessCheck{ctrl: ctrl}
	mock.recorder = &MockReadinessCheckMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReadinessCheck) EXPECT() *MockReadinessCheckMockRecorder {
	return m.recorder
}

// Check mocks base method.
func (m *MockReadinessCheck) Check(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Check indicates an expected call of Check.
func (mr *MockReadinessCheckMockRecorder) Check(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockReadinessCheck)(nil).Check), arg0)
}
