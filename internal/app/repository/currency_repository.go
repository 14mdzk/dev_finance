package repository

import (
	"errors"
	"fmt"

	"github.com/14mdzk/dev_finance/internal/app/model"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type CurrencyRepository struct {
	DB *sqlx.DB
}

func NewCurrencyRepository(DBConn *sqlx.DB) *CurrencyRepository {
	return &CurrencyRepository{
		DB: DBConn,
	}
}

func (repo *CurrencyRepository) Browse(pagination string) ([]model.Currency, error) {
	var (
		currencies []model.Currency
		statement  = fmt.Sprintf(`
			SELECT id, name, abbreviation
			FROM currencies
			%s
		`, pagination)
	)

	rows, err := repo.DB.Queryx(statement)
	if err != nil {
		log.Error(fmt.Errorf("error at Currency Repository - Browse: %w", err))
		return currencies, err
	}

	for rows.Next() {
		var currency model.Currency

		err = rows.StructScan(&currency)
		if err != nil {
			continue
		}

		currencies = append(currencies, currency)
	}

	return currencies, nil
}
func (repo *CurrencyRepository) GetByID(id string) (model.Currency, error) {
	var (
		currency  model.Currency
		statement = `
			SELECT id, name, abbreviation
			FROM currencies
			WHERE id = $1
		`
	)

	err := repo.DB.QueryRowx(statement, id).StructScan(&currency)
	if err != nil {
		log.Error(fmt.Errorf("error at Currency Repository - GetByID: %w", err))
		return currency, err
	}

	return currency, nil
}

func (repo *CurrencyRepository) Create(currency model.Currency) error {
	statement := `INSERT INTO currencies(name, abbreviation) VALUES($1, $2)`
	_, err := repo.DB.Exec(statement, currency.Name, currency.Abbreviation)
	if err != nil {
		log.Error(fmt.Errorf("error at Currency Repository - Create: %w", err))
		return err
	}

	return nil
}
func (repo *CurrencyRepository) Update(currencyID string, currency model.Currency) error {
	statement := `
		UPDATE currencies
		SET
			updated_at = NOW(),
			name = $2,
			abbreviation = $3
		WHERE id = $1
	`
	result, err := repo.DB.Exec(statement, currencyID, currency.Name, currency.Abbreviation)
	if err != nil {
		log.Error(fmt.Errorf("error at Currency Repository - Update: %w", err))
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected <= 0 {
		return errors.New("error at Currency Repository - Update: no rows affected")
	}

	return nil
}
func (repo *CurrencyRepository) Delete(currencyID string) error {
	statement := `DELETE FROM currencies WHERE id = $1`

	result, err := repo.DB.Exec(statement, currencyID)
	if err != nil {
		log.Error(fmt.Errorf("error at Currency Repository - Delete: %w", err))
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected <= 0 {
		return errors.New("error at Currency Repository - Delete: no rows affected")
	}

	return nil
}
