package repository

import (
	"fmt"

	"github.com/14mdzk/dev_finance/internal/app/model"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type AuthRepository struct {
	DB *sqlx.DB
}

func NewAuthRepository(DBConn *sqlx.DB) *AuthRepository {
	return &AuthRepository{
		DB: DBConn,
	}
}

func (ar *AuthRepository) Find(userID int, refreshToken string) (model.Auth, error) {
	var (
		statement = `
			SELECT
				id, token, auth_type, user_id, expired_at
			FROM auths
			WHERE
				user_id = $1 AND token = $2
		`
		auth model.Auth
	)

	err := ar.DB.QueryRowx(statement, userID, refreshToken).StructScan(&auth)
	if err != nil {
		log.Error(fmt.Errorf("error: AuthRepository - Find: %w", err))
		return auth, err
	}

	return auth, nil
}

func (ar *AuthRepository) Create(auth model.Auth) error {
	var (
		statement = `
			INSERT INTO
				auths(user_id, auth_type, token, expired_at)
			VALUES ($1, $2, $3, $4)
		`
	)

	_, err := ar.DB.Exec(statement, auth.UserID, auth.AuthType, auth.Token, auth.ExpiredAt)
	if err != nil {
		if err != nil {
			log.Error(fmt.Errorf("error: AuthRepository - Create: %w", err))
			return err
		}
	}

	return nil
}

func (ar *AuthRepository) DeleteAllByUserID(userID int) error {
	var (
		statement = `
			DELETE FROM auths
			WHERE
				user_id = $1
		`
	)

	_, err := ar.DB.Exec(statement, userID)
	if err != nil {
		log.Error(fmt.Errorf("error: AuthRepository - DeleteAllByUserID: %w", err))
		return err
	}

	return nil
}
