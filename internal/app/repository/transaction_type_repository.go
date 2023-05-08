package repository

import (
	"errors"
	"fmt"

	"github.com/14mdzk/dev_finance/internal/app/model"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type TransactionTypeRepository struct {
	DB *sqlx.DB
}

func NewTransactionTypeRepository(DBConn *sqlx.DB) *TransactionTypeRepository {
	return &TransactionTypeRepository{
		DB: DBConn,
	}
}

func (repo *TransactionTypeRepository) Browse() ([]model.TransactionType, error) {
	var (
		transactionTypes []model.TransactionType
		statement        = `
			SELECT id, name, description
			FROM transaction_types
		`
	)

	rows, err := repo.DB.Queryx(statement)
	if err != nil {
		log.Error(fmt.Errorf("error at Transaction Type Repository - Browse: %w", err))
		return transactionTypes, err
	}

	for rows.Next() {
		var transactionType model.TransactionType
		rows.StructScan(&transactionType)

		transactionTypes = append(transactionTypes, transactionType)
	}

	return transactionTypes, nil
}
func (repo *TransactionTypeRepository) GetByID(id string) (model.TransactionType, error) {
	var (
		transactionType model.TransactionType
		statement       = `
			SELECT id, name, description
			FROM transaction_types
			WHERE id = $1
		`
	)

	err := repo.DB.QueryRowx(statement, id).StructScan(&transactionType)
	if err != nil {
		log.Error(fmt.Errorf("error at Transaction Type Repository - GetByID: %w", err))
		return transactionType, err
	}

	return transactionType, nil
}

func (repo *TransactionTypeRepository) Create(transactionType model.TransactionType) error {
	statement := `INSERT INTO transactionTypes(name, description) VALUES($1, $2)`
	_, err := repo.DB.Exec(statement, transactionType.Name, transactionType.Description)
	if err != nil {
		log.Error(fmt.Errorf("error at Transaction Type Repository - Create: %w", err))
		return err
	}

	return nil
}
func (repo *TransactionTypeRepository) Update(transactionTypeID string, transactionType model.TransactionType) error {
	statement := `
		UPDATE transactionTypes
		SET
			updated_at = NOW(),
			name = $2,
			description = $3
		WHERE id = $1
	`
	result, err := repo.DB.Exec(statement, transactionTypeID, transactionType.Name, transactionType.Description)
	if err != nil {
		log.Error(fmt.Errorf("error at Transaction Type Repository - Update: %w", err))
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected <= 0 {
		return errors.New("error at Transaction Type Repository - Update: no rows affected")
	}

	return nil
}
func (repo *TransactionTypeRepository) Delete(transactionTypeID string) error {
	statement := `DELETE FROM transactionTypes WHERE id = $1`

	result, err := repo.DB.Exec(statement, transactionTypeID)
	if err != nil {
		log.Error(fmt.Errorf("error at Transaction Type Repository - Delete: %w", err))
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected <= 0 {
		return errors.New("error at Transaction Type Repository - Delete: no rows affected")
	}

	return nil
}
