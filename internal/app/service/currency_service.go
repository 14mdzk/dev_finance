package service

import (
	"errors"

	"github.com/14mdzk/dev_finance/internal/app/model"
	"github.com/14mdzk/dev_finance/internal/app/schema"
	"github.com/14mdzk/dev_finance/internal/pkg/reason"
)

type CurrencyRepository interface {
	Browse() ([]model.Currency, error)
	GetByID(currencyID string) (model.Currency, error)
	Create(model.Currency) error
	Update(currencyID string, currency model.Currency) error
	Delete(currencyID string) error
}

type CurrencyService struct {
	repo CurrencyRepository
}

func NewCurrencyService(repo CurrencyRepository) *CurrencyService {
	return &CurrencyService{
		repo: repo,
	}
}

func (svc *CurrencyService) BrowseAll() ([]schema.CurrencyResp, error) {
	var resp []schema.CurrencyResp

	currencies, err := svc.repo.Browse()
	if err != nil {
		return nil, errors.New(reason.CurrencyFailedBrowse)
	}

	for _, currency := range currencies {
		item := schema.CurrencyResp{
			ID:           currency.ID,
			Name:         currency.Name,
			Abbreviation: currency.Abbreviation,
		}

		resp = append(resp, item)
	}

	return resp, nil
}

func (svc *CurrencyService) FindById(currencyID string) (schema.CurrencyResp, error) {
	var resp schema.CurrencyResp

	currency, err := svc.repo.GetByID(currencyID)
	if err != nil {
		return resp, errors.New(reason.CurrencyNotFound)
	}

	resp.ID = currency.ID
	resp.Name = currency.Name
	resp.Abbreviation = currency.Abbreviation

	return resp, nil
}
func (svc *CurrencyService) Create(req schema.CurrencyReq) error {
	currency := model.Currency{
		Name:         req.Name,
		Abbreviation: req.Abbreviation,
	}

	err := svc.repo.Create(currency)
	if err != nil {
		return errors.New(reason.CurrencyFailedCreate)
	}

	return nil
}
func (svc *CurrencyService) Update(currencyID string, req schema.CurrencyReq) error {
	currency := model.Currency{
		Name:         req.Name,
		Abbreviation: req.Abbreviation,
	}

	err := svc.repo.Create(currency)
	if err != nil {
		return errors.New(reason.CurrencyFailedCreate)
	}

	return nil
}

func (svc *CurrencyService) Delete(currencyID string) error {
	err := svc.repo.Delete(currencyID)
	if err != nil {
		return errors.New(reason.CurrencyFailedDelete)
	}

	return nil
}
