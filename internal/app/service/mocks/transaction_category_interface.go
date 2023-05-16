// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/service/transaction_category_interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	model "github.com/14mdzk/dev_finance/internal/app/model"
	gomock "github.com/golang/mock/gomock"
)

// MockITransactionCategoryRepository is a mock of ITransactionCategoryRepository interface.
type MockITransactionCategoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockITransactionCategoryRepositoryMockRecorder
}

// MockITransactionCategoryRepositoryMockRecorder is the mock recorder for MockITransactionCategoryRepository.
type MockITransactionCategoryRepositoryMockRecorder struct {
	mock *MockITransactionCategoryRepository
}

// NewMockITransactionCategoryRepository creates a new mock instance.
func NewMockITransactionCategoryRepository(ctrl *gomock.Controller) *MockITransactionCategoryRepository {
	mock := &MockITransactionCategoryRepository{ctrl: ctrl}
	mock.recorder = &MockITransactionCategoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITransactionCategoryRepository) EXPECT() *MockITransactionCategoryRepositoryMockRecorder {
	return m.recorder
}

// Browse mocks base method.
func (m *MockITransactionCategoryRepository) Browse() ([]model.TransactionCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Browse")
	ret0, _ := ret[0].([]model.TransactionCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Browse indicates an expected call of Browse.
func (mr *MockITransactionCategoryRepositoryMockRecorder) Browse() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Browse", reflect.TypeOf((*MockITransactionCategoryRepository)(nil).Browse))
}

// Create mocks base method.
func (m *MockITransactionCategoryRepository) Create(arg0 model.TransactionCategory) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockITransactionCategoryRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockITransactionCategoryRepository)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockITransactionCategoryRepository) Delete(transactionCategoryID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", transactionCategoryID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockITransactionCategoryRepositoryMockRecorder) Delete(transactionCategoryID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockITransactionCategoryRepository)(nil).Delete), transactionCategoryID)
}

// GetByID mocks base method.
func (m *MockITransactionCategoryRepository) GetByID(transactionCategoryID string) (model.TransactionCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", transactionCategoryID)
	ret0, _ := ret[0].(model.TransactionCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockITransactionCategoryRepositoryMockRecorder) GetByID(transactionCategoryID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockITransactionCategoryRepository)(nil).GetByID), transactionCategoryID)
}

// Update mocks base method.
func (m *MockITransactionCategoryRepository) Update(transactionCategoryID string, transactionCategory model.TransactionCategory) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", transactionCategoryID, transactionCategory)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockITransactionCategoryRepositoryMockRecorder) Update(transactionCategoryID, transactionCategory interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockITransactionCategoryRepository)(nil).Update), transactionCategoryID, transactionCategory)
}
