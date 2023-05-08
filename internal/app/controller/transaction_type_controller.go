package controller

import (
	"net/http"

	"github.com/14mdzk/dev_finance/internal/app/schema"
	"github.com/14mdzk/dev_finance/internal/pkg/handler"
	"github.com/14mdzk/dev_finance/internal/pkg/reason"
	"github.com/gin-gonic/gin"
)

type TransactionTypeService interface {
	BrowseAll() ([]schema.TransactionTypeResp, error)
	FindById(TransactionTypeID string) (schema.TransactionTypeResp, error)
	Create(req schema.TransactionTypeReq) error
	Update(TransactionTypeID string, req schema.TransactionTypeReq) error
	Delete(TransactionTypeID string) error
}

type TransactionTypeController struct {
	service TransactionTypeService
}

func NewTransactionTypeController(service TransactionTypeService) *TransactionTypeController {
	return &TransactionTypeController{
		service: service,
	}
}

func (ctrl *TransactionTypeController) BrowseTransactionType(ctx *gin.Context) {
	resp, err := ctrl.service.BrowseAll()
	if err != nil {
		handler.ResponseError(ctx, http.StatusInternalServerError, reason.InternalServerError)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, reason.TransactionTypeSuccessBrowse, resp)
}

func (ctrl *TransactionTypeController) FindTransactionType(ctx *gin.Context) {
	transactionType, err := ctrl.service.FindById(ctx.Param("id"))
	if err != nil {
		handler.ResponseError(ctx, http.StatusNotFound, reason.TransactionTypeNotFound)
		return
	}

	resp := schema.TransactionTypeResp{
		ID:          transactionType.ID,
		Name:        transactionType.Name,
		Description: transactionType.Description,
	}

	handler.ResponseSuccess(ctx, http.StatusOK, reason.TransactionTypeFound, resp)
}

func (ctrl *TransactionTypeController) CreateTransactionType(ctx *gin.Context) {
	var req schema.TransactionTypeReq

	if !handler.BindAndValidate(ctx, &req) {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.InvalidRequest)
		return
	}

	err := ctrl.service.Create(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusInternalServerError, reason.InternalServerError)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, reason.TransactionTypeSuccessCreate, nil)
}

func (ctrl *TransactionTypeController) UpdateTransactionType(ctx *gin.Context) {
	var req schema.TransactionTypeReq

	if !handler.BindAndValidate(ctx, &req) {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.InvalidRequest)
		return
	}

	err := ctrl.service.Update(ctx.Param("id"), req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusInternalServerError, reason.InternalServerError)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, reason.TransactionTypeSuccessUpdate, nil)
}

func (ctrl *TransactionTypeController) DeleteTransactionType(ctx *gin.Context) {
	err := ctrl.service.Delete(ctx.Param("id"))
	if err != nil {
		handler.ResponseError(ctx, http.StatusInternalServerError, reason.InternalServerError)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, reason.TransactionTypeSuccessDelete, nil)
}
