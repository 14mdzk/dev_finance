package service

import (
	"fmt"

	"github.com/14mdzk/dev_finance/internal/app/model"
	"github.com/14mdzk/dev_finance/internal/app/schema"
	"github.com/14mdzk/dev_finance/internal/pkg/helper"
	log "github.com/sirupsen/logrus"
)

type RegisterRepository interface {
	Create(user model.User) error
}

type RegisterService struct {
	repo RegisterRepository
}

func NewRegisterService(repo RegisterRepository) *RegisterService {
	return &RegisterService{
		repo: repo,
	}
}

func (svc *RegisterService) Register(req schema.UserRegisterReq) error {
	password, err := helper.HashPassword(req.Password)
	if err != nil {
		log.Error(fmt.Errorf("error at Register Service Create: %w", err))
		return err
	}

	user := model.User{
		Username: req.Username,
		FullName: req.FullName,
		Password: password,
	}

	err = svc.repo.Create(user)
	if err != nil {
		return err
	}

	return nil
}
