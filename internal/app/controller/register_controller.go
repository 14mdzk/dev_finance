package controller

import (
	"net/http"

	"github.com/14mdzk/dev_finance/internal/app/schema"
	"github.com/14mdzk/dev_finance/internal/pkg/handler"
	"github.com/14mdzk/dev_finance/internal/pkg/reason"
	"github.com/gin-gonic/gin"
)

type RegisterService interface {
	Register(req schema.UserRegisterReq) error
}

type RegisterController struct {
	svc RegisterService
}

func NewRegisterController(svc RegisterService) *RegisterController {
	return &RegisterController{
		svc: svc,
	}
}

func (ctrl *RegisterController) RegisterUser(ctx *gin.Context) {
	var req schema.UserRegisterReq
	if !handler.BindAndValidate(ctx, &req) {
		return
	}

	err := ctrl.svc.Register(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusNotFound, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, reason.RegisterSuccess, nil, nil)
}
