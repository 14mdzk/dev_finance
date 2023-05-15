package service

import (
	"github.com/14mdzk/dev_finance/internal/app/model"
)

type ITransactionRepository interface {
	Browse(userID int) ([]model.Transaction, error)
	BrowseByTransactionType(userID int, transactionType int) ([]model.Transaction, error)

	GetByID(userID int, transactionID int) (model.Transaction, error)

	Create(userID int, transaction model.Transaction) error
	Update(userID int, transactionID int, transaction model.Transaction) error
	Delete(userID int, transactionID int) error
}
