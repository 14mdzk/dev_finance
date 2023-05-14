package repository

import (
	"errors"
	"fmt"

	"github.com/14mdzk/dev_finance/internal/app/model"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type TransactionRepository struct {
	DB *sqlx.DB
}

func NewTransactionRepository(DBConn *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{
		DB: DBConn,
	}
}

func (repo *TransactionRepository) Browse(userID int) ([]model.Transaction, error) {
	var (
		transactions []model.Transaction
		statement    = `
			SELECT 
				transactions.id, 
				user_id, 
				ROW_TO_JSON(currencies.*) as currency,
				ROW_TO_JSON(trTypes.*) as transaction_type,
				ROW_TO_JSON(trCategories.*) as transaction_category,
				transaction_date, 
				note, 
				amount
			FROM transactions
			LEFT JOIN transaction_types as trTypes
				ON transaction_type_id = trTypes.id
			LEFT JOIN transaction_categories as trCategories
				ON transaction_category_id = trCategories.id
			LEFT JOIN currencies
				ON currency_id = currencies.id
			WHERE user_id = $1
		`
	)

	rows, err := repo.DB.Queryx(statement, userID)
	if err != nil {
		log.Error(fmt.Errorf("error at Transaction Repository - Browse: %w", err))
		return transactions, err
	}

	transactions = repo.applyRows(rows)

	return transactions, nil
}

func (repo *TransactionRepository) BrowseByTransactionType(userID int, transactionTypeID int) ([]model.Transaction, error) {
	var (
		transactions []model.Transaction
		statement    = `
			SELECT 
				transactions.id, 
				user_id, 
				ROW_TO_JSON(currencies.*) as currency,
				ROW_TO_JSON(trTypes.*) as transaction_type,
				ROW_TO_JSON(trCategories.*) as transaction_category,
				transaction_date, 
				note, 
				amount
			FROM transactions
			LEFT JOIN transaction_types as trTypes
				ON transaction_type_id = trTypes.id
			LEFT JOIN transaction_categories as trCategories
				ON transaction_category_id = trCategories.id
			LEFT JOIN currencies
				ON currency_id = currencies.id
			WHERE 
				user_id = $1 AND transaction_type_id = $2
		`
	)

	rows, err := repo.DB.Queryx(statement, userID, transactionTypeID)
	if err != nil {
		log.Error(fmt.Errorf("error at Transaction Repository - BrowseByTransactionType: %w", err))
		return transactions, err
	}

	transactions = repo.applyRows(rows)

	return transactions, nil
}

func (repo *TransactionRepository) applyRows(data *sqlx.Rows) (transactions []model.Transaction) {
	for data.Next() {
		var transaction model.Transaction
		data.StructScan(&transaction)
		transactions = append(transactions, transaction)
	}

	return transactions
}

func (repo *TransactionRepository) GetByID(userID int, transactionID int) (model.Transaction, error) {
	var (
		transaction model.Transaction
		statement   = `
			SELECT 
				transactions.id, 
				user_id, 
				ROW_TO_JSON(currencies.*) as currency,
				ROW_TO_JSON(trTypes.*) as transaction_type,
				ROW_TO_JSON(trCategories.*) as transaction_category,
				transaction_date, 
				note, 
				amount
			FROM transactions
			LEFT JOIN transaction_types as trTypes
				ON transaction_type_id = trTypes.id
			LEFT JOIN transaction_categories as trCategories
				ON transaction_category_id = trCategories.id
			LEFT JOIN currencies
				ON currency_id = currencies.id
			WHERE 
				id = $1 AND user_id = $2
		`
	)

	err := repo.DB.QueryRowx(statement, transactionID, userID).StructScan(&transaction)
	if err != nil {
		log.Error(fmt.Errorf("error at Transaction Repository - GetByID: %w", err))
		return transaction, err
	}

	return transaction, nil
}

func (repo *TransactionRepository) Create(userID int, transaction model.Transaction) error {
	statement := `
		INSERT INTO transactions(
			user_id, transaction_type_id, transaction_category_id, 
			currency_id, transaction_date, note, amount
		) 
		VALUES($1, $2, $3, $4, $5, $6, $7)`
	log.Error(fmt.Print(userID, transaction))

	_, err := repo.DB.Exec(
		statement,
		userID,
		transaction.TransactionTypeID,
		transaction.TransactionCategoryID,
		transaction.CurrencyID,
		transaction.TransactionDate,
		transaction.Note,
		transaction.Amount,
	)
	if err != nil {
		log.Error(fmt.Errorf("error at Transaction Repository - Create: %w", err))
		return err
	}

	return nil
}

func (repo *TransactionRepository) Update(userID int, transactionID int, transaction model.Transaction) error {
	statement := `
		UPDATE transactions
		SET
			updated_at = NOW(),
			transaction_type_id = $3,
			transaction_category_id = $4,
			currency_id = $5,
			transaction_date = $6,
			note = $7,
			amount = $8
		WHERE id = $1 AND user_id = $2
	`
	result, err := repo.DB.Exec(
		statement,
		transactionID,
		userID,
		transaction.TransactionTypeID,
		transaction.TransactionCategoryID,
		transaction.CurrencyID,
		transaction.TransactionDate,
		transaction.Note,
		transaction.Amount,
	)
	if err != nil {
		log.Error(fmt.Errorf("error at Transaction Repository - Update: %w", err))
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected <= 0 {
		return errors.New("error at Transaction Repository - Update: no rows affected")
	}

	return nil
}
func (repo *TransactionRepository) Delete(userID int, transactionID int) error {
	statement := `DELETE FROM transactions WHERE id = $1 AND user_id $2`

	result, err := repo.DB.Exec(statement, transactionID, userID)
	if err != nil {
		log.Error(fmt.Errorf("error at Transaction Repository - Delete: %w", err))
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected <= 0 {
		return errors.New("error at Transaction Repository - Delete: no rows affected")
	}

	return nil
}
