package service

import (
	"errors"

	"github.com/14mdzk/dev_finance/internal/app/model"
	"github.com/14mdzk/dev_finance/internal/app/schema"
	"github.com/14mdzk/dev_finance/internal/pkg/reason"
)

type TransactionTypeRepository interface {
	Browse() ([]model.TransactionType, error)
	GetByID(transactionTypeID string) (model.TransactionType, error)
	Create(model.TransactionType) error
	Update(transactionTypeID string, transactionType model.TransactionType) error
	Delete(transactionTypeID string) error
}

type TransactionTypeService struct {
	repo TransactionTypeRepository
}

func NewTransactionTypeService(repo TransactionTypeRepository) *TransactionTypeService {
	return &TransactionTypeService{
		repo: repo,
	}
}

func (svc *TransactionTypeService) BrowseAll() ([]schema.TransactionTypeResp, error) {
	var resp []schema.TransactionTypeResp

	transactionTypes, err := svc.repo.Browse()
	if err != nil {
		return nil, errors.New(reason.TransactionTypeFailedBrowse)
	}

	for _, transactionType := range transactionTypes {
		item := schema.TransactionTypeResp{
			ID:          transactionType.ID,
			Name:        transactionType.Name,
			Description: transactionType.Description,
		}

		resp = append(resp, item)
	}

	return resp, nil
}

func (svc *TransactionTypeService) FindById(transactionTypeID string) (schema.TransactionTypeResp, error) {
	var resp schema.TransactionTypeResp

	transactionType, err := svc.repo.GetByID(transactionTypeID)
	if err != nil {
		return resp, errors.New(reason.TransactionTypeNotFound)
	}

	resp.ID = transactionType.ID
	resp.Name = transactionType.Name
	resp.Description = transactionType.Description

	return resp, nil
}
func (svc *TransactionTypeService) Create(req schema.TransactionTypeReq) error {
	transactionType := model.TransactionType{
		Name:        req.Name,
		Description: req.Description,
	}

	err := svc.repo.Create(transactionType)
	if err != nil {
		return errors.New(reason.TransactionTypeFailedCreate)
	}

	return nil
}
func (svc *TransactionTypeService) Update(transactionTypeID string, req schema.TransactionTypeReq) error {
	transactionType := model.TransactionType{
		Name:        req.Name,
		Description: req.Description,
	}

	err := svc.repo.Create(transactionType)
	if err != nil {
		return errors.New(reason.TransactionTypeFailedCreate)
	}

	return nil
}

func (svc *TransactionTypeService) Delete(transactionTypeID string) error {
	err := svc.repo.Delete(transactionTypeID)
	if err != nil {
		return errors.New(reason.TransactionTypeFailedDelete)
	}

	return nil
}
