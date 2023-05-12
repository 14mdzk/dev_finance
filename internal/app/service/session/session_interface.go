package session

import (
	"time"

	"github.com/14mdzk/dev_finance/internal/app/model"
)

type UserRepository interface {
	GetByID(userID string) (model.User, error)
	GetByUsername(username string) (model.User, error)
}
type AuthRepository interface {
	Create(auth model.Auth) error
	DeleteAllByUserID(userID int) error
	Find(userID int, refreshToken string) (model.Auth, error)
}

type TokenGenerator interface {
	GenerateAccessToken(userID int) (string, time.Time, error)
	GenerateRefreshToken(userID int) (string, time.Time, error)
}
