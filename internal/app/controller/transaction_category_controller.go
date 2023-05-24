package controller

import (
	"net/http"

	"github.com/14mdzk/dev_finance/internal/app/schema"
	"github.com/14mdzk/dev_finance/internal/pkg/handler"
	"github.com/14mdzk/dev_finance/internal/pkg/reason"
	"github.com/gin-gonic/gin"
)

type TransactionCategoryService interface {
	BrowseAll(pagination schema.PaginationReq) ([]schema.TransactionCategoryResp, error)
	FindById(transactionCategoryID string) (schema.TransactionCategoryResp, error)
	Create(req schema.TransactionCategoryReq) error
	Update(transactionCategoryID string, req schema.TransactionCategoryReq) error
	Delete(transactionCategoryID string) error
}

type TransactionCategoryController struct {
	service TransactionCategoryService
}

func NewTransactionCategoryController(service TransactionCategoryService) *TransactionCategoryController {
	return &TransactionCategoryController{
		service: service,
	}
}

func (ctrl *TransactionCategoryController) BrowseTransactionCategory(ctx *gin.Context) {
	var pagination schema.PaginationReq
	handler.BindPagination(ctx, &pagination)
	// pagination := schema.PaginationReq{Page: 1, PageSize: 2}

	resp, err := ctrl.service.BrowseAll(pagination)
	if err != nil {
		handler.ResponseError(ctx, http.StatusInternalServerError, reason.InternalServerError)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, reason.TransactionCategorySuccessBrowse, resp, pagination)
}

func (ctrl *TransactionCategoryController) FindTransactionCategory(ctx *gin.Context) {
	transactionCategory, err := ctrl.service.FindById(ctx.Param("id"))
	if err != nil {
		handler.ResponseError(ctx, http.StatusNotFound, reason.TransactionCategoryNotFound)
		return
	}

	resp := schema.TransactionCategoryResp{
		ID:          transactionCategory.ID,
		Name:        transactionCategory.Name,
		Description: transactionCategory.Description,
	}

	handler.ResponseSuccess(ctx, http.StatusOK, reason.TransactionCategoryFound, resp, nil)
}

func (ctrl *TransactionCategoryController) CreateTransactionCategory(ctx *gin.Context) {
	var req schema.TransactionCategoryReq

	if !handler.BindAndValidate(ctx, &req) {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.InvalidRequest)
		return
	}

	err := ctrl.service.Create(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusInternalServerError, reason.InternalServerError)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, reason.TransactionCategorySuccessCreate, nil, nil)
}

func (ctrl *TransactionCategoryController) UpdateTransactionCategory(ctx *gin.Context) {
	var req schema.TransactionCategoryReq

	if !handler.BindAndValidate(ctx, &req) {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.InvalidRequest)
		return
	}

	err := ctrl.service.Update(ctx.Param("id"), req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusInternalServerError, reason.InternalServerError)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, reason.TransactionCategorySuccessUpdate, nil, nil)
}

func (ctrl *TransactionCategoryController) DeleteTransactionCategory(ctx *gin.Context) {
	err := ctrl.service.Delete(ctx.Param("id"))
	if err != nil {
		handler.ResponseError(ctx, http.StatusInternalServerError, reason.InternalServerError)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, reason.TransactionCategorySuccessDelete, nil, nil)
}
