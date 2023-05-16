package controller

import (
	"net/http"

	"github.com/14mdzk/dev_finance/internal/app/schema"
	"github.com/14mdzk/dev_finance/internal/pkg/handler"
	"github.com/14mdzk/dev_finance/internal/pkg/reason"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	BrowseAll(pagination schema.PaginationReq) ([]schema.UserResp, error)
	Delete(userID string) error
	ChangePassword(userID string, password string) error
}

type UserController struct {
	service UserService
}

func NewUserController(service UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (ctrl *UserController) BrowseUser(ctx *gin.Context) {
	var pagination schema.PaginationReq
	handler.BindPagination(ctx, &pagination)

	resp, err := ctrl.service.BrowseAll(pagination)
	if err != nil {
		handler.ResponseError(ctx, http.StatusInternalServerError, reason.InternalServerError)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, reason.CurrencySuccessBrowse, resp, pagination)
}

func (ctrl *UserController) DeleteUser(ctx *gin.Context) {
	err := ctrl.service.Delete(ctx.Param("id"))
	if err != nil {
		handler.ResponseError(ctx, http.StatusNotFound, reason.CurrencyNotFound)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, reason.CurrencySuccessDelete, nil, nil)
}

func (ctrl *UserController) ChangePasswordUser(ctx *gin.Context) {
	var req schema.UserChangePasswordReq

	err := ctrl.service.Delete(ctx.Param("id"))
	if err != nil {
		handler.ResponseError(ctx, http.StatusNotFound, reason.CurrencyNotFound)
		return
	}

	if !handler.BindAndValidate(ctx, &req) {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.InvalidRequest)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, reason.CurrencySuccessDelete, nil, nil)
}
