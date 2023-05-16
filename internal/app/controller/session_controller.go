package controller

import (
	"net/http"
	"strconv"

	"github.com/14mdzk/dev_finance/internal/app/schema"
	"github.com/14mdzk/dev_finance/internal/pkg/handler"
	"github.com/14mdzk/dev_finance/internal/pkg/reason"
	"github.com/gin-gonic/gin"
)

type ISessionService interface {
	Authenticate(req schema.LoginReq) (schema.LoginResp, error)
	RefreshToken(req schema.RefreshTokenReq) (schema.RefreshTokenResp, error)
	Destroy(userID int) error
}

type ITokenService interface {
	VerifyRefreshToken(tokenString string) (string, error)
}

type SessionController struct {
	sessionSvc ISessionService
	tokenSvc   ITokenService
}

func NewSessionController(sessionSvc ISessionService, tokenSvc ITokenService) *SessionController {
	return &SessionController{
		sessionSvc: sessionSvc,
		tokenSvc:   tokenSvc,
	}
}

func (ctrl *SessionController) Login(ctx *gin.Context) {
	var req schema.LoginReq

	if !handler.BindAndValidate(ctx, &req) {
		return
	}

	resp, err := ctrl.sessionSvc.Authenticate(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, reason.LoginSuccess, resp, nil)
}

func (ctrl *SessionController) Logout(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.GetString("user_id"))

	err := ctrl.sessionSvc.Destroy(userID)
	if err != nil {
		handler.ResponseError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, reason.LogoutSuccess, nil, nil)
}

func (ctrl *SessionController) RefreshToken(ctx *gin.Context) {
	refreshToken := ctx.GetHeader("refresh_token")

	sub, err := ctrl.tokenSvc.VerifyRefreshToken(refreshToken)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	intSub, _ := strconv.Atoi(sub)
	req := schema.RefreshTokenReq{
		RefreshToken: refreshToken,
		UserID:       intSub,
	}

	resp, err := ctrl.sessionSvc.RefreshToken(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, reason.RefreshToken, resp, nil)
}
