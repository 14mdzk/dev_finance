package session

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/14mdzk/dev_finance/internal/app/model"
	"github.com/14mdzk/dev_finance/internal/app/schema"
	"github.com/14mdzk/dev_finance/internal/pkg/helper"
	"github.com/14mdzk/dev_finance/internal/pkg/reason"
	log "github.com/sirupsen/logrus"
)

type SessionService struct {
	userRepo   UserRepository
	authRepo   AuthRepository
	tokenMaker TokenGenerator
}

func NewSessionService(
	userRepo UserRepository,
	authRepo AuthRepository,
	tokenMaker TokenGenerator,
) *SessionService {
	return &SessionService{
		userRepo:   userRepo,
		authRepo:   authRepo,
		tokenMaker: tokenMaker,
	}
}

func (svc *SessionService) Authenticate(req schema.LoginReq) (schema.LoginResp, error) {
	var resp schema.LoginResp

	existingUser, _ := svc.userRepo.GetByUsername(req.Username)
	if existingUser.ID <= 0 {
		return resp, errors.New(reason.LoginFailed)
	}

	passwordMatch := helper.VerifyPassword(existingUser.Password, req.Password)
	if !passwordMatch {
		return resp, errors.New(reason.LoginFailed)
	}

	accessToken, _, err := svc.tokenMaker.GenerateAccessToken(existingUser.ID)
	if err != nil {
		log.Error(fmt.Errorf("error at Session Service - Access Token: %w", err))
		return resp, errors.New(reason.LoginFailed)
	}

	refreshToken, expiredAt, err := svc.tokenMaker.GenerateRefreshToken(existingUser.ID)
	if err != nil {
		log.Error(fmt.Errorf("error at Session Service - Refresh Token: %w", err))
		return resp, errors.New(reason.LoginFailed)
	}

	authPayload := model.Auth{
		UserID:    existingUser.ID,
		Token:     refreshToken,
		AuthType:  "refresh_token",
		ExpiredAt: expiredAt,
	}

	err = svc.authRepo.Create(authPayload)
	if err != nil {
		return resp, errors.New(reason.LoginFailed)
	}

	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken

	return resp, nil
}

func (svc *SessionService) RefreshToken(req schema.RefreshTokenReq) (schema.RefreshTokenResp, error) {
	var resp schema.RefreshTokenResp

	existingUser, _ := svc.userRepo.GetByID(strconv.Itoa(req.UserID))
	if existingUser.ID <= 0 {
		return resp, errors.New(reason.RefreshTokenInvalid)
	}

	auth, err := svc.authRepo.Find(existingUser.ID, req.RefreshToken)
	if err != nil || auth.UserID <= 0 {
		log.Error(fmt.Errorf("error at SessionService - Refresh: %w", err))
		return resp, err
	}

	accesstoken, _, _ := svc.tokenMaker.GenerateAccessToken(existingUser.ID)
	resp.AccessToken = accesstoken

	return resp, nil
}

func (svc *SessionService) Destroy(userID int) error {
	err := svc.authRepo.DeleteAllByUserID(userID)
	if err != nil {
		log.Error(fmt.Errorf("error at Session Service - Logout: %w", err))
		return err
	}

	return nil
}
