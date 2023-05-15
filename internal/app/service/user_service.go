package service

import (
	"github.com/14mdzk/dev_finance/internal/app/schema"
	"github.com/14mdzk/dev_finance/internal/pkg/helper"
)

type UserService struct {
	repo IUserRepository
}

func NewUserService(repo IUserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) BrowseAll() ([]schema.UserResp, error) {
	var resp []schema.UserResp

	users, err := svc.repo.Browse()
	if err != nil {
		return resp, err
	}

	for _, item := range users {
		user := schema.UserResp{
			ID:           item.ID,
			Username:     item.Username,
			FullName:     item.FullName,
			RegisteredAt: item.CreatedAt.Format("01 January 2006"),
		}

		resp = append(resp, user)
	}

	return resp, nil
}

func (svc *UserService) Delete(userID string) error {
	err := svc.repo.Delete(userID)
	if err != nil {
		return err
	}

	return nil
}

func (svc *UserService) ChangePassword(userID string, password string) error {
	password, err := helper.HashPassword(password)
	if err != nil {
		return err
	}

	err = svc.repo.UpdatePassword(userID, password)
	if err != nil {
		return err
	}

	return nil
}
