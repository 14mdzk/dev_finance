// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/service/transaction_type_interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	model "github.com/14mdzk/dev_finance/internal/app/model"
	gomock "github.com/golang/mock/gomock"
)

// MockITransactionTypeRepository is a mock of ITransactionTypeRepository interface.
type MockITransactionTypeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockITransactionTypeRepositoryMockRecorder
}

// MockITransactionTypeRepositoryMockRecorder is the mock recorder for MockITransactionTypeRepository.
type MockITransactionTypeRepositoryMockRecorder struct {
	mock *MockITransactionTypeRepository
}

// NewMockITransactionTypeRepository creates a new mock instance.
func NewMockITransactionTypeRepository(ctrl *gomock.Controller) *MockITransactionTypeRepository {
	mock := &MockITransactionTypeRepository{ctrl: ctrl}
	mock.recorder = &MockITransactionTypeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITransactionTypeRepository) EXPECT() *MockITransactionTypeRepositoryMockRecorder {
	return m.recorder
}

// Browse mocks base method.
func (m *MockITransactionTypeRepository) Browse(pagination string) ([]model.TransactionType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Browse", pagination)
	ret0, _ := ret[0].([]model.TransactionType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Browse indicates an expected call of Browse.
func (mr *MockITransactionTypeRepositoryMockRecorder) Browse(pagination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Browse", reflect.TypeOf((*MockITransactionTypeRepository)(nil).Browse), pagination)
}

// Create mocks base method.
func (m *MockITransactionTypeRepository) Create(arg0 model.TransactionType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockITransactionTypeRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockITransactionTypeRepository)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockITransactionTypeRepository) Delete(transactionTypeID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", transactionTypeID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockITransactionTypeRepositoryMockRecorder) Delete(transactionTypeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockITransactionTypeRepository)(nil).Delete), transactionTypeID)
}

// GetByID mocks base method.
func (m *MockITransactionTypeRepository) GetByID(transactionTypeID string) (model.TransactionType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", transactionTypeID)
	ret0, _ := ret[0].(model.TransactionType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockITransactionTypeRepositoryMockRecorder) GetByID(transactionTypeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockITransactionTypeRepository)(nil).GetByID), transactionTypeID)
}

// Update mocks base method.
func (m *MockITransactionTypeRepository) Update(transactionTypeID string, transactionType model.TransactionType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", transactionTypeID, transactionType)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockITransactionTypeRepositoryMockRecorder) Update(transactionTypeID, transactionType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockITransactionTypeRepository)(nil).Update), transactionTypeID, transactionType)
}
