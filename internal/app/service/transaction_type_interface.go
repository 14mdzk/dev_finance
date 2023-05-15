package service

import "github.com/14mdzk/dev_finance/internal/app/model"

type ITransactionTypeRepository interface {
	Browse() ([]model.TransactionType, error)
	GetByID(transactionTypeID string) (model.TransactionType, error)
	Create(model.TransactionType) error
	Update(transactionTypeID string, transactionType model.TransactionType) error
	Delete(transactionTypeID string) error
}
