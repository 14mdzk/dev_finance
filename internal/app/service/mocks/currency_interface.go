// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/service/currency_interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	model "github.com/14mdzk/dev_finance/internal/app/model"
	gomock "github.com/golang/mock/gomock"
)

// MockICurrencyRepository is a mock of ICurrencyRepository interface.
type MockICurrencyRepository struct {
	ctrl     *gomock.Controller
	recorder *MockICurrencyRepositoryMockRecorder
}

// MockICurrencyRepositoryMockRecorder is the mock recorder for MockICurrencyRepository.
type MockICurrencyRepositoryMockRecorder struct {
	mock *MockICurrencyRepository
}

// NewMockICurrencyRepository creates a new mock instance.
func NewMockICurrencyRepository(ctrl *gomock.Controller) *MockICurrencyRepository {
	mock := &MockICurrencyRepository{ctrl: ctrl}
	mock.recorder = &MockICurrencyRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICurrencyRepository) EXPECT() *MockICurrencyRepositoryMockRecorder {
	return m.recorder
}

// Browse mocks base method.
func (m *MockICurrencyRepository) Browse(pagination string) ([]model.Currency, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Browse", pagination)
	ret0, _ := ret[0].([]model.Currency)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Browse indicates an expected call of Browse.
func (mr *MockICurrencyRepositoryMockRecorder) Browse(pagination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Browse", reflect.TypeOf((*MockICurrencyRepository)(nil).Browse), pagination)
}

// Create mocks base method.
func (m *MockICurrencyRepository) Create(arg0 model.Currency) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockICurrencyRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockICurrencyRepository)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockICurrencyRepository) Delete(currencyID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", currencyID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockICurrencyRepositoryMockRecorder) Delete(currencyID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockICurrencyRepository)(nil).Delete), currencyID)
}

// GetByID mocks base method.
func (m *MockICurrencyRepository) GetByID(currencyID string) (model.Currency, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", currencyID)
	ret0, _ := ret[0].(model.Currency)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockICurrencyRepositoryMockRecorder) GetByID(currencyID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockICurrencyRepository)(nil).GetByID), currencyID)
}

// Update mocks base method.
func (m *MockICurrencyRepository) Update(currencyID string, currency model.Currency) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", currencyID, currency)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockICurrencyRepositoryMockRecorder) Update(currencyID, currency interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockICurrencyRepository)(nil).Update), currencyID, currency)
}
