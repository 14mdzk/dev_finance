package service

import (
	"errors"

	"github.com/14mdzk/dev_finance/internal/app/model"
	"github.com/14mdzk/dev_finance/internal/app/schema"
	"github.com/14mdzk/dev_finance/internal/pkg/reason"
)

type TransactionCategoryService struct {
	repo ITransactionCategoryRepository
}

func NewTransactionCategoryService(repo ITransactionCategoryRepository) *TransactionCategoryService {
	return &TransactionCategoryService{
		repo: repo,
	}
}

func (svc *TransactionCategoryService) BrowseAll() ([]schema.TransactionCategoryResp, error) {
	var resp []schema.TransactionCategoryResp

	transactionCategories, err := svc.repo.Browse()
	if err != nil {
		return nil, errors.New(reason.TransactionCategoryFailedBrowse)
	}

	for _, transactionCategory := range transactionCategories {
		item := schema.TransactionCategoryResp{
			ID:          transactionCategory.ID,
			Name:        transactionCategory.Name,
			Description: transactionCategory.Description,
		}

		resp = append(resp, item)
	}

	return resp, nil
}

func (svc *TransactionCategoryService) FindById(transactionCategoryID string) (schema.TransactionCategoryResp, error) {
	var resp schema.TransactionCategoryResp

	transactionCategory, err := svc.repo.GetByID(transactionCategoryID)
	if err != nil {
		return resp, errors.New(reason.TransactionCategoryNotFound)
	}

	resp.ID = transactionCategory.ID
	resp.Name = transactionCategory.Name
	resp.Description = transactionCategory.Description

	return resp, nil
}
func (svc *TransactionCategoryService) Create(req schema.TransactionCategoryReq) error {
	transactionCategory := model.TransactionCategory{
		Name:        req.Name,
		Description: req.Description,
	}

	err := svc.repo.Create(transactionCategory)
	if err != nil {
		return errors.New(reason.TransactionCategoryFailedCreate)
	}

	return nil
}
func (svc *TransactionCategoryService) Update(transactionCategoryID string, req schema.TransactionCategoryReq) error {
	transactionCategory := model.TransactionCategory{
		Name:        req.Name,
		Description: req.Description,
	}

	err := svc.repo.Create(transactionCategory)
	if err != nil {
		return errors.New(reason.TransactionCategoryFailedCreate)
	}

	return nil
}

func (svc *TransactionCategoryService) Delete(transactionCategoryID string) error {
	err := svc.repo.Delete(transactionCategoryID)
	if err != nil {
		return errors.New(reason.TransactionCategoryFailedDelete)
	}

	return nil
}
