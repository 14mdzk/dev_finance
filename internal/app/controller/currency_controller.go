package controller

import (
	"net/http"

	"github.com/14mdzk/dev_finance/internal/app/schema"
	"github.com/14mdzk/dev_finance/internal/pkg/handler"
	"github.com/14mdzk/dev_finance/internal/pkg/reason"
	"github.com/gin-gonic/gin"
)

type CurrencyService interface {
	BrowseAll() ([]schema.CurrencyResp, error)
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
	resp, err := ctrl.service.BrowseAll()
	if err != nil {
		handler.ResponseError(ctx, http.StatusInternalServerError, reason.InternalServerError)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, reason.CurrencySuccessBrowse, resp)
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

	handler.ResponseSuccess(ctx, http.StatusOK, reason.CurrencyFound, resp)
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

	handler.ResponseSuccess(ctx, http.StatusCreated, reason.CurrencySuccessCreate, nil)
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

	handler.ResponseSuccess(ctx, http.StatusOK, reason.CurrencySuccessUpdate, nil)
}

func (ctrl *CurrencyController) DeleteCurrency(ctx *gin.Context) {
	err := ctrl.service.Delete(ctx.Param("id"))
	if err != nil {
		handler.ResponseError(ctx, http.StatusInternalServerError, reason.InternalServerError)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, reason.CurrencySuccessDelete, nil)
}
