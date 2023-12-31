// Code generated by MockGen. DO NOT EDIT.
// Source: /home/kostik/go/pkg/mod/github.com/jackc/pgx/v5@v5.4.3/rows.go
//
// Generated by this command:
//
//	mockgen -source=/home/kostik/go/pkg/mod/github.com/jackc/pgx/v5@v5.4.3/rows.go Rows
//
// Package mock_pgx is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	pgx "github.com/jackc/pgx/v5"
	pgconn "github.com/jackc/pgx/v5/pgconn"
	gomock "go.uber.org/mock/gomock"
)

// MockRows is a mock of Rows interface.
type MockRows struct {
	ctrl     *gomock.Controller
	recorder *MockRowsMockRecorder
}

// MockRowsMockRecorder is the mock recorder for MockRows.
type MockRowsMockRecorder struct {
	mock *MockRows
}

// NewMockRows creates a new mock instance.
func NewMockRows(ctrl *gomock.Controller) *MockRows {
	mock := &MockRows{ctrl: ctrl}
	mock.recorder = &MockRowsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRows) EXPECT() *MockRowsMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRows) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockRowsMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRows)(nil).Close))
}

// CommandTag mocks base method.
func (m *MockRows) CommandTag() pgconn.CommandTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommandTag")
	ret0, _ := ret[0].(pgconn.CommandTag)
	return ret0
}

// CommandTag indicates an expected call of CommandTag.
func (mr *MockRowsMockRecorder) CommandTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommandTag", reflect.TypeOf((*MockRows)(nil).CommandTag))
}

// Conn mocks base method.
func (m *MockRows) Conn() *pgx.Conn {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Conn")
	ret0, _ := ret[0].(*pgx.Conn)
	return ret0
}

// Conn indicates an expected call of Conn.
func (mr *MockRowsMockRecorder) Conn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Conn", reflect.TypeOf((*MockRows)(nil).Conn))
}

// Err mocks base method.
func (m *MockRows) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockRowsMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockRows)(nil).Err))
}

// FieldDescriptions mocks base method.
func (m *MockRows) FieldDescriptions() []pgconn.FieldDescription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FieldDescriptions")
	ret0, _ := ret[0].([]pgconn.FieldDescription)
	return ret0
}

// FieldDescriptions indicates an expected call of FieldDescriptions.
func (mr *MockRowsMockRecorder) FieldDescriptions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FieldDescriptions", reflect.TypeOf((*MockRows)(nil).FieldDescriptions))
}

// Next mocks base method.
func (m *MockRows) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockRowsMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockRows)(nil).Next))
}

// RawValues mocks base method.
func (m *MockRows) RawValues() [][]byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RawValues")
	ret0, _ := ret[0].([][]byte)
	return ret0
}

// RawValues indicates an expected call of RawValues.
func (mr *MockRowsMockRecorder) RawValues() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RawValues", reflect.TypeOf((*MockRows)(nil).RawValues))
}

// Scan mocks base method.
func (m *MockRows) Scan(dest ...any) error {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range dest {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Scan", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Scan indicates an expected call of Scan.
func (mr *MockRowsMockRecorder) Scan(dest ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockRows)(nil).Scan), dest...)
}

// Values mocks base method.
func (m *MockRows) Values() ([]any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Values")
	ret0, _ := ret[0].([]any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Values indicates an expected call of Values.
func (mr *MockRowsMockRecorder) Values() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Values", reflect.TypeOf((*MockRows)(nil).Values))
}

// MockRow is a mock of Row interface.
type MockRow struct {
	ctrl     *gomock.Controller
	recorder *MockRowMockRecorder
}

// MockRowMockRecorder is the mock recorder for MockRow.
type MockRowMockRecorder struct {
	mock *MockRow
}

// NewMockRow creates a new mock instance.
func NewMockRow(ctrl *gomock.Controller) *MockRow {
	mock := &MockRow{ctrl: ctrl}
	mock.recorder = &MockRowMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRow) EXPECT() *MockRowMockRecorder {
	return m.recorder
}

// Scan mocks base method.
func (m *MockRow) Scan(dest ...any) error {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range dest {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Scan", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Scan indicates an expected call of Scan.
func (mr *MockRowMockRecorder) Scan(dest ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockRow)(nil).Scan), dest...)
}

// MockRowScanner is a mock of RowScanner interface.
type MockRowScanner struct {
	ctrl     *gomock.Controller
	recorder *MockRowScannerMockRecorder
}

// MockRowScannerMockRecorder is the mock recorder for MockRowScanner.
type MockRowScannerMockRecorder struct {
	mock *MockRowScanner
}

// NewMockRowScanner creates a new mock instance.
func NewMockRowScanner(ctrl *gomock.Controller) *MockRowScanner {
	mock := &MockRowScanner{ctrl: ctrl}
	mock.recorder = &MockRowScannerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRowScanner) EXPECT() *MockRowScannerMockRecorder {
	return m.recorder
}

// ScanRow mocks base method.
func (m *MockRowScanner) ScanRow(rows pgx.Rows) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScanRow", rows)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScanRow indicates an expected call of ScanRow.
func (mr *MockRowScannerMockRecorder) ScanRow(rows any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanRow", reflect.TypeOf((*MockRowScanner)(nil).ScanRow), rows)
}

// MockCollectableRow is a mock of CollectableRow interface.
type MockCollectableRow struct {
	ctrl     *gomock.Controller
	recorder *MockCollectableRowMockRecorder
}

// MockCollectableRowMockRecorder is the mock recorder for MockCollectableRow.
type MockCollectableRowMockRecorder struct {
	mock *MockCollectableRow
}

// NewMockCollectableRow creates a new mock instance.
func NewMockCollectableRow(ctrl *gomock.Controller) *MockCollectableRow {
	mock := &MockCollectableRow{ctrl: ctrl}
	mock.recorder = &MockCollectableRowMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCollectableRow) EXPECT() *MockCollectableRowMockRecorder {
	return m.recorder
}

// FieldDescriptions mocks base method.
func (m *MockCollectableRow) FieldDescriptions() []pgconn.FieldDescription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FieldDescriptions")
	ret0, _ := ret[0].([]pgconn.FieldDescription)
	return ret0
}

// FieldDescriptions indicates an expected call of FieldDescriptions.
func (mr *MockCollectableRowMockRecorder) FieldDescriptions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FieldDescriptions", reflect.TypeOf((*MockCollectableRow)(nil).FieldDescriptions))
}

// RawValues mocks base method.
func (m *MockCollectableRow) RawValues() [][]byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RawValues")
	ret0, _ := ret[0].([][]byte)
	return ret0
}

// RawValues indicates an expected call of RawValues.
func (mr *MockCollectableRowMockRecorder) RawValues() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RawValues", reflect.TypeOf((*MockCollectableRow)(nil).RawValues))
}

// Scan mocks base method.
func (m *MockCollectableRow) Scan(dest ...any) error {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range dest {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Scan", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Scan indicates an expected call of Scan.
func (mr *MockCollectableRowMockRecorder) Scan(dest ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockCollectableRow)(nil).Scan), dest...)
}

// Values mocks base method.
func (m *MockCollectableRow) Values() ([]any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Values")
	ret0, _ := ret[0].([]any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Values indicates an expected call of Values.
func (mr *MockCollectableRowMockRecorder) Values() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Values", reflect.TypeOf((*MockCollectableRow)(nil).Values))
}
