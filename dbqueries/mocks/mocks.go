// Code generated by MockGen. DO NOT EDIT.
// Source: testTask/dbqueries (interfaces: CompositeQI,CounterQI,TransactionQI,UserQI)

// Package mock_dbqueries is a generated GoMock package.
package mock_dbqueries

import (
	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
	reflect "reflect"
	db "txCancel/models/db"
)

// MockCompositeQI is a mock of CompositeQI interface
type MockCompositeQI struct {
	ctrl     *gomock.Controller
	recorder *MockCompositeQIMockRecorder
}

// MockCompositeQIMockRecorder is the mock recorder for MockCompositeQI
type MockCompositeQIMockRecorder struct {
	mock *MockCompositeQI
}

// NewMockCompositeQI creates a new mock instance
func NewMockCompositeQI(ctrl *gomock.Controller) *MockCompositeQI {
	mock := &MockCompositeQI{ctrl: ctrl}
	mock.recorder = &MockCompositeQIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCompositeQI) EXPECT() *MockCompositeQIMockRecorder {
	return m.recorder
}

// ApplyTransaction mocks base method
func (m *MockCompositeQI) ApplyTransaction(arg0 *db.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyTransaction", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyTransaction indicates an expected call of ApplyTransaction
func (mr *MockCompositeQIMockRecorder) ApplyTransaction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyTransaction", reflect.TypeOf((*MockCompositeQI)(nil).ApplyTransaction), arg0)
}

// CancelTransactions mocks base method
func (m *MockCompositeQI) CancelTransactions(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelTransactions", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelTransactions indicates an expected call of CancelTransactions
func (mr *MockCompositeQIMockRecorder) CancelTransactions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelTransactions", reflect.TypeOf((*MockCompositeQI)(nil).CancelTransactions), arg0)
}

// MockCounterQI is a mock of CounterQI interface
type MockCounterQI struct {
	ctrl     *gomock.Controller
	recorder *MockCounterQIMockRecorder
}

// MockCounterQIMockRecorder is the mock recorder for MockCounterQI
type MockCounterQIMockRecorder struct {
	mock *MockCounterQI
}

// NewMockCounterQI creates a new mock instance
func NewMockCounterQI(ctrl *gomock.Controller) *MockCounterQI {
	mock := &MockCounterQI{ctrl: ctrl}
	mock.recorder = &MockCounterQIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCounterQI) EXPECT() *MockCounterQIMockRecorder {
	return m.recorder
}

// GetByName mocks base method
func (m *MockCounterQI) GetByName(arg0 *sqlx.Tx, arg1 string) (*db.Counter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", arg0, arg1)
	ret0, _ := ret[0].(*db.Counter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName
func (mr *MockCounterQIMockRecorder) GetByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockCounterQI)(nil).GetByName), arg0, arg1)
}

// UpdateValueTx mocks base method
func (m *MockCounterQI) UpdateValueTx(arg0 *sqlx.Tx, arg1 string, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateValueTx", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateValueTx indicates an expected call of UpdateValueTx
func (mr *MockCounterQIMockRecorder) UpdateValueTx(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateValueTx", reflect.TypeOf((*MockCounterQI)(nil).UpdateValueTx), arg0, arg1, arg2)
}

// MockTransactionQI is a mock of TransactionQI interface
type MockTransactionQI struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionQIMockRecorder
}

// MockTransactionQIMockRecorder is the mock recorder for MockTransactionQI
type MockTransactionQIMockRecorder struct {
	mock *MockTransactionQI
}

// NewMockTransactionQI creates a new mock instance
func NewMockTransactionQI(ctrl *gomock.Controller) *MockTransactionQI {
	mock := &MockTransactionQI{ctrl: ctrl}
	mock.recorder = &MockTransactionQIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTransactionQI) EXPECT() *MockTransactionQIMockRecorder {
	return m.recorder
}

// CancelTx mocks base method
func (m *MockTransactionQI) CancelTx(arg0 *sqlx.Tx, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelTx", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelTx indicates an expected call of CancelTx
func (mr *MockTransactionQIMockRecorder) CancelTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelTx", reflect.TypeOf((*MockTransactionQI)(nil).CancelTx), arg0, arg1)
}

// GetListGreaterThan mocks base method
func (m *MockTransactionQI) GetListGreaterThan(arg0 *sqlx.Tx, arg1 int) ([]db.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListGreaterThan", arg0, arg1)
	ret0, _ := ret[0].([]db.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListGreaterThan indicates an expected call of GetListGreaterThan
func (mr *MockTransactionQIMockRecorder) GetListGreaterThan(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListGreaterThan", reflect.TypeOf((*MockTransactionQI)(nil).GetListGreaterThan), arg0, arg1)
}

// InsertTx mocks base method
func (m *MockTransactionQI) InsertTx(arg0 *sqlx.Tx, arg1 *db.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertTx", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertTx indicates an expected call of InsertTx
func (mr *MockTransactionQIMockRecorder) InsertTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertTx", reflect.TypeOf((*MockTransactionQI)(nil).InsertTx), arg0, arg1)
}

// MockUserQI is a mock of UserQI interface
type MockUserQI struct {
	ctrl     *gomock.Controller
	recorder *MockUserQIMockRecorder
}

// MockUserQIMockRecorder is the mock recorder for MockUserQI
type MockUserQIMockRecorder struct {
	mock *MockUserQI
}

// NewMockUserQI creates a new mock instance
func NewMockUserQI(ctrl *gomock.Controller) *MockUserQI {
	mock := &MockUserQI{ctrl: ctrl}
	mock.recorder = &MockUserQIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserQI) EXPECT() *MockUserQIMockRecorder {
	return m.recorder
}

// GetByID mocks base method
func (m *MockUserQI) GetByID(arg0 int) (*db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(*db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockUserQIMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUserQI)(nil).GetByID), arg0)
}

// UpdateBalanceTx mocks base method
func (m *MockUserQI) UpdateBalanceTx(arg0 *sqlx.Tx, arg1 int, arg2 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBalanceTx", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBalanceTx indicates an expected call of UpdateBalanceTx
func (mr *MockUserQIMockRecorder) UpdateBalanceTx(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBalanceTx", reflect.TypeOf((*MockUserQI)(nil).UpdateBalanceTx), arg0, arg1, arg2)
}
