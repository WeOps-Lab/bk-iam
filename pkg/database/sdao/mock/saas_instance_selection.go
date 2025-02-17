// Code generated by MockGen. DO NOT EDIT.
// Source: saas_instance_selection.go

// Package mock is a generated GoMock package.
package mock

import (
	sdao "iam/pkg/database/sdao"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
)

// MockSaaSInstanceSelectionManager is a mock of SaaSInstanceSelectionManager interface.
type MockSaaSInstanceSelectionManager struct {
	ctrl     *gomock.Controller
	recorder *MockSaaSInstanceSelectionManagerMockRecorder
}

// MockSaaSInstanceSelectionManagerMockRecorder is the mock recorder for MockSaaSInstanceSelectionManager.
type MockSaaSInstanceSelectionManagerMockRecorder struct {
	mock *MockSaaSInstanceSelectionManager
}

// NewMockSaaSInstanceSelectionManager creates a new mock instance.
func NewMockSaaSInstanceSelectionManager(ctrl *gomock.Controller) *MockSaaSInstanceSelectionManager {
	mock := &MockSaaSInstanceSelectionManager{ctrl: ctrl}
	mock.recorder = &MockSaaSInstanceSelectionManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSaaSInstanceSelectionManager) EXPECT() *MockSaaSInstanceSelectionManagerMockRecorder {
	return m.recorder
}

// BulkCreateWithTx mocks base method.
func (m *MockSaaSInstanceSelectionManager) BulkCreateWithTx(tx *sqlx.Tx, saasInstanceSelections []sdao.SaaSInstanceSelection) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreateWithTx", tx, saasInstanceSelections)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkCreateWithTx indicates an expected call of BulkCreateWithTx.
func (mr *MockSaaSInstanceSelectionManagerMockRecorder) BulkCreateWithTx(tx, saasInstanceSelections interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreateWithTx", reflect.TypeOf((*MockSaaSInstanceSelectionManager)(nil).BulkCreateWithTx), tx, saasInstanceSelections)
}

// BulkDeleteWithTx mocks base method.
func (m *MockSaaSInstanceSelectionManager) BulkDeleteWithTx(tx *sqlx.Tx, system string, ids []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteWithTx", tx, system, ids)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkDeleteWithTx indicates an expected call of BulkDeleteWithTx.
func (mr *MockSaaSInstanceSelectionManagerMockRecorder) BulkDeleteWithTx(tx, system, ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteWithTx", reflect.TypeOf((*MockSaaSInstanceSelectionManager)(nil).BulkDeleteWithTx), tx, system, ids)
}

// Get mocks base method.
func (m *MockSaaSInstanceSelectionManager) Get(system, id string) (sdao.SaaSInstanceSelection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", system, id)
	ret0, _ := ret[0].(sdao.SaaSInstanceSelection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSaaSInstanceSelectionManagerMockRecorder) Get(system, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSaaSInstanceSelectionManager)(nil).Get), system, id)
}

// ListBySystem mocks base method.
func (m *MockSaaSInstanceSelectionManager) ListBySystem(system string) ([]sdao.SaaSInstanceSelection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBySystem", system)
	ret0, _ := ret[0].([]sdao.SaaSInstanceSelection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBySystem indicates an expected call of ListBySystem.
func (mr *MockSaaSInstanceSelectionManagerMockRecorder) ListBySystem(system interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySystem", reflect.TypeOf((*MockSaaSInstanceSelectionManager)(nil).ListBySystem), system)
}

// Update mocks base method.
func (m *MockSaaSInstanceSelectionManager) Update(system, instanceSelectionID string, sis sdao.SaaSInstanceSelection) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", system, instanceSelectionID, sis)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockSaaSInstanceSelectionManagerMockRecorder) Update(system, instanceSelectionID, sis interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSaaSInstanceSelectionManager)(nil).Update), system, instanceSelectionID, sis)
}
