package service

import "github.com/14mdzk/dev_finance/internal/app/model"

type IUserRepository interface {
	Browse() ([]model.User, error)
	Create(user model.User) error
	Delete(userID string) error
	UpdatePassword(userID string, password string) error
}
