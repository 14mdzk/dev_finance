package service

import "github.com/14mdzk/dev_finance/internal/app/model"

type ITransactionCategoryRepository interface {
	Browse(pagination string) ([]model.TransactionCategory, error)
	GetByID(transactionCategoryID string) (model.TransactionCategory, error)
	Create(model.TransactionCategory) error
	Update(transactionCategoryID string, transactionCategory model.TransactionCategory) error
	Delete(transactionCategoryID string) error
}
