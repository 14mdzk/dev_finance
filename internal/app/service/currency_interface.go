package service

import "github.com/14mdzk/dev_finance/internal/app/model"

type ICurrencyRepository interface {
	Browse(pagination string) ([]model.Currency, error)
	GetByID(currencyID string) (model.Currency, error)
	Create(model.Currency) error
	Update(currencyID string, currency model.Currency) error
	Delete(currencyID string) error
}
