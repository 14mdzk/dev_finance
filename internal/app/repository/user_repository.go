package repository

import (
	"errors"
	"fmt"

	"github.com/14mdzk/dev_finance/internal/app/model"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(DBConn *sqlx.DB) *UserRepository {
	return &UserRepository{
		DB: DBConn,
	}
}

func (repo *UserRepository) Browse(pagination string) ([]model.User, error) {
	var (
		users     []model.User
		statement = fmt.Sprintf(`
			SELECT id, username, full_name, created_at
			FROM users
			%s
		`, pagination)
	)

	rows, err := repo.DB.Queryx(statement)
	if err != nil {
		log.Error(fmt.Errorf("error at User Repository - Browse: %w", err))
		return users, err
	}

	for rows.Next() {
		var user model.User
		rows.StructScan(&user)

		users = append(users, user)
	}

	return users, nil
}

func (repo *UserRepository) GetByID(userID string) (model.User, error) {
	var (
		user      model.User
		statement = `
			SELECT id, username, full_name, created_at
			FROM users
			WHERE id = $1
		`
	)

	err := repo.DB.QueryRowx(statement, userID).StructScan(&user)
	if err != nil {
		log.Error(fmt.Errorf("error at User Repository - GetByID: %w", err))
		return user, err
	}

	return user, nil
}

func (repo *UserRepository) GetByUsername(username string) (model.User, error) {
	var (
		user      model.User
		statement = `
			SELECT id, username, full_name, password, created_at
			FROM users
			WHERE username = $1
		`
	)

	err := repo.DB.QueryRowx(statement, username).StructScan(&user)
	if err != nil {
		log.Error(fmt.Errorf("error at User Repository - GetByID: %w", err))
		return user, err
	}

	return user, nil
}

func (repo *UserRepository) Create(user model.User) error {
	statement := `
		INSERT INTO users(username, full_name, password) VALUES($1, $2, $3)
	`
	result, err := repo.DB.Exec(statement, user.Username, user.FullName, user.Password)
	if err != nil {
		log.Error(fmt.Errorf("error at User Repository - Create: %w", err))
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected <= 0 {
		return errors.New("error at User Repository - Create: no rows affected")
	}

	return nil
}

func (repo *UserRepository) Update(userID string, user model.User) error {
	statement := `
		UPDATE users
		SET
			updated_at = NOW(),
			full_name = $2
		WHERE id = $1
	`
	result, err := repo.DB.Exec(statement, userID, user.FullName)
	if err != nil {
		log.Error(fmt.Errorf("error at User Repository - Update: %w", err))
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected <= 0 {
		return errors.New("error at User Repository - Update: no rows affected")
	}

	return nil
}

func (repo *UserRepository) UpdatePassword(userID string, password string) error {
	statement := `
		UPDATE users
		SET
			updated_at = NOW(),
			password = $2
		WHERE id = $1
	`
	result, err := repo.DB.Exec(statement, userID, password)
	if err != nil {
		log.Error(fmt.Errorf("error at User Repository - Update: %w", err))
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected <= 0 {
		return errors.New("error at User Repository - Update: no rows affected")
	}

	return nil
}

func (repo *UserRepository) Delete(userID string) error {
	statement := `DELETE FROM users WHERE id = $1`

	result, err := repo.DB.Exec(statement, userID)
	if err != nil {
		log.Error(fmt.Errorf("error at User Repository - Delete: %w", err))
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected <= 0 {
		return errors.New("error at User Repository - Delete: no rows affected")
	}

	return nil
}
