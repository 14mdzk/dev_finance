package controller

import (
	"net/http"

	"github.com/14mdzk/dev_finance/internal/app/schema"
	"github.com/14mdzk/dev_finance/internal/pkg/handler"
	"github.com/14mdzk/dev_finance/internal/pkg/reason"
	"github.com/gin-gonic/gin"
)

type CurrencyService interface {
	BrowseAll(pagination schema.PaginationReq) ([]schema.CurrencyResp, error)
	FindById(currencyID string) (schema.CurrencyResp, error)
	Create(req schema.CurrencyReq) error
	Update(currencyID string, req schema.CurrencyReq) error
	Delete(currencyID string) error
}

type CurrencyController struct {
	service CurrencyService
}

func NewCurrencyController(service CurrencyService) *CurrencyController {
	return &CurrencyController{
		service: service,
	}
}

func (ctrl *CurrencyController) BrowseCurrency(ctx *gin.Context) {
	var pagination schema.PaginationReq
	handler.BindPagination(ctx, &pagination)

	resp, err := ctrl.service.BrowseAll(pagination)
	if err != nil {
		handler.ResponseError(ctx, http.StatusInternalServerError, reason.InternalServerError)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, reason.CurrencySuccessBrowse, resp, pagination)
}

func (ctrl *CurrencyController) FindCurrency(ctx *gin.Context) {
	currency, err := ctrl.service.FindById(ctx.Param("id"))
	if err != nil {
		handler.ResponseError(ctx, http.StatusInternalServerError, reason.InternalServerError)
		return
	}

	resp := schema.CurrencyResp{
		ID:           currency.ID,
		Name:         currency.Name,
		Abbreviation: currency.Abbreviation,
	}

	handler.ResponseSuccess(ctx, http.StatusOK, reason.CurrencyFound, resp, nil)
}

func (ctrl *CurrencyController) CreateCurrency(ctx *gin.Context) {
	var req schema.CurrencyReq

	if !handler.BindAndValidate(ctx, &req) {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.InvalidRequest)
		return
	}

	err := ctrl.service.Create(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusInternalServerError, reason.InternalServerError)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, reason.CurrencySuccessCreate, nil, nil)
}

func (ctrl *CurrencyController) UpdateCurrency(ctx *gin.Context) {
	var req schema.CurrencyReq

	if !handler.BindAndValidate(ctx, &req) {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.InvalidRequest)
		return
	}

	err := ctrl.service.Update(ctx.Param("id"), req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusInternalServerError, reason.InternalServerError)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, reason.CurrencySuccessUpdate, nil, nil)
}

func (ctrl *CurrencyController) DeleteCurrency(ctx *gin.Context) {
	err := ctrl.service.Delete(ctx.Param("id"))
	if err != nil {
		handler.ResponseError(ctx, http.StatusInternalServerError, reason.InternalServerError)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, reason.CurrencySuccessDelete, nil, nil)
}
