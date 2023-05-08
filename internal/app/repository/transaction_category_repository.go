package repository

import (
	"errors"
	"fmt"

	"github.com/14mdzk/dev_finance/internal/app/model"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type TransactionCategoryRepository struct {
	DB *sqlx.DB
}

func NewTransactionCategoryRepository(DBConn *sqlx.DB) *TransactionCategoryRepository {
	return &TransactionCategoryRepository{
		DB: DBConn,
	}
}

func (repo *TransactionCategoryRepository) Browse() ([]model.TransactionCategory, error) {
	var (
		transactionCategories []model.TransactionCategory
		statement             = `
			SELECT id, name, description
			FROM transaction_categories
		`
	)

	rows, err := repo.DB.Queryx(statement)
	if err != nil {
		log.Error(fmt.Errorf("error at Transaction Category Repository - Browse: %w", err))
		return transactionCategories, err
	}

	for rows.Next() {
		var transactionCategory model.TransactionCategory
		rows.StructScan(&transactionCategory)

		transactionCategories = append(transactionCategories, transactionCategory)
	}

	return transactionCategories, nil
}
func (repo *TransactionCategoryRepository) GetByID(id string) (model.TransactionCategory, error) {
	var (
		transactionCategory model.TransactionCategory
		statement           = `
			SELECT id, name, description
			FROM transaction_categories
			WHERE id = $1
		`
	)

	err := repo.DB.QueryRowx(statement, id).StructScan(&transactionCategory)
	if err != nil {
		log.Error(fmt.Errorf("error at Transaction Category Repository - GetByID: %w", err))
		return transactionCategory, err
	}

	return transactionCategory, nil
}

func (repo *TransactionCategoryRepository) Create(transactionCategory model.TransactionCategory) error {
	statement := `INSERT INTO transaction_categories(name, description) VALUES($1, $2)`
	_, err := repo.DB.Exec(statement, transactionCategory.Name, transactionCategory.Description)
	if err != nil {
		log.Error(fmt.Errorf("error at Transaction Category Repository - Create: %w", err))
		return err
	}

	return nil
}
func (repo *TransactionCategoryRepository) Update(transactionCategoryID string, transactionCategory model.TransactionCategory) error {
	statement := `
		UPDATE transaction_categories
		SET
			updated_at = NOW(),
			name = $2,
			description = $3
		WHERE id = $1
	`
	result, err := repo.DB.Exec(statement, transactionCategoryID, transactionCategory.Name, transactionCategory.Description)
	if err != nil {
		log.Error(fmt.Errorf("error at Transaction Category Repository - Update: %w", err))
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected <= 0 {
		return errors.New("error at Transaction Category Repository - Update: no rows affected")
	}

	return nil
}
func (repo *TransactionCategoryRepository) Delete(transactionCategoryID string) error {
	statement := `DELETE FROM transaction_categories WHERE id = $1`

	result, err := repo.DB.Exec(statement, transactionCategoryID)
	if err != nil {
		log.Error(fmt.Errorf("error at Transaction Category Repository - Delete: %w", err))
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected <= 0 {
		return errors.New("error at Transaction Category Repository - Delete: no rows affected")
	}

	return nil
}
