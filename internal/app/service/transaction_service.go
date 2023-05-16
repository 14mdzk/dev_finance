package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/14mdzk/dev_finance/internal/app/model"
	"github.com/14mdzk/dev_finance/internal/app/schema"
	"github.com/14mdzk/dev_finance/internal/pkg/helper"
	"github.com/14mdzk/dev_finance/internal/pkg/reason"
	log "github.com/sirupsen/logrus"
)

type TransactionService struct {
	transactionRepo         ITransactionRepository
	transactionCategoryRepo ITransactionCategoryRepository
	transactionTypeRepo     ITransactionTypeRepository
	userRepo                IUserRepository
}

func NewTransactionService(
	transactionRepo ITransactionRepository,
	transactionCategoryRepo ITransactionCategoryRepository,
	transactionTypeRepo ITransactionTypeRepository,
	userRepo IUserRepository,
) *TransactionService {
	return &TransactionService{
		transactionRepo:         transactionRepo,
		transactionCategoryRepo: transactionCategoryRepo,
		transactionTypeRepo:     transactionTypeRepo,
		userRepo:                userRepo,
	}
}

func (svc *TransactionService) BrowseAll(userID int, transactionTypeID int, pagination schema.PaginationReq) ([]schema.TransactionResp, error) {
	var resp []schema.TransactionResp

	transactions, err := svc.browse(userID, transactionTypeID, helper.GeneratePaginationQuery(pagination.Page, pagination.PageSize))
	if err != nil {
		return resp, errors.New(reason.TransactionFailedBrowse)
	}

	for _, transaction := range transactions {
		var (
			transactionCategoryResp schema.TransactionCategoryResp
			transactionTypeResp     schema.TransactionTypeResp
			currencyResp            schema.CurrencyResp
		)

		err = json.Unmarshal([]byte(transaction.TransactionCategory), &transactionCategoryResp)
		if err != nil {
			log.Error(fmt.Sprint("error at Transaction Service: %w", err))
		}

		err = json.Unmarshal([]byte(transaction.TransactionType), &transactionTypeResp)
		if err != nil {
			log.Error(fmt.Sprint("error at Transaction Service: %w", err))
		}

		err = json.Unmarshal([]byte(transaction.Currency), &currencyResp)
		if err != nil {
			log.Error(fmt.Sprint("error at Transaction Service: %w", err))
		}

		transactionDate, _ := time.Parse("", transaction.TransactionDate)

		item := schema.TransactionResp{
			ID:                  transaction.ID,
			TransactionCategory: transactionCategoryResp,
			TransactionType:     transactionTypeResp,
			Currency:            currencyResp,
			TransactionDate:     transactionDate,
			Note:                transaction.Note,
			Amount:              transaction.Amount,
		}

		resp = append(resp, item)
	}

	return resp, nil
}

func (svc *TransactionService) browse(userID int, transactionTypeID int, pagination string) (transactions []model.Transaction, err error) {
	if transactionTypeID > 0 {
		transactions, err := svc.transactionRepo.BrowseByTransactionType(userID, transactionTypeID, pagination)
		if err != nil {
			return transactions, err
		}

		return transactions, nil
	}

	transactions, err = svc.transactionRepo.Browse(userID, pagination)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (svc *TransactionService) Create(userID int, req schema.TransactionReq) error {
	transaction := model.Transaction{
		TransactionTypeID:     req.TransactionTypeID,
		TransactionCategoryID: req.TransactionCategoryID,
		TransactionDate:       req.TransactionDate.Format("2006-01-02"),
		CurrencyID:            req.CurrencyID,
		Amount:                req.Amount,
		Note:                  req.Note,
	}

	err := svc.transactionRepo.Create(userID, transaction)
	if err != nil {
		return errors.New(reason.TransactionFailedCreate)
	}

	return nil
}

func (svc *TransactionService) Update(userID int, transactionID int, req schema.TransactionReq) error {
	transaction := model.Transaction{
		TransactionTypeID:     req.TransactionTypeID,
		TransactionCategoryID: req.TransactionCategoryID,
		TransactionDate:       req.TransactionDate.Format("2006-01-02"),
		CurrencyID:            req.CurrencyID,
		Amount:                req.Amount,
		Note:                  req.Note,
	}

	err := svc.transactionRepo.Update(userID, transactionID, transaction)
	if err != nil {
		return errors.New(reason.TransactionFailedCreate)
	}

	return nil
}

func (svc *TransactionService) Delete(userID int, transactionID int) error {
	err := svc.transactionRepo.Delete(userID, transactionID)
	if err != nil {
		return errors.New(reason.TransactionFailedDelete)
	}

	return nil
}
