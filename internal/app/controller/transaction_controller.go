package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/14mdzk/dev_finance/internal/app/schema"
	"github.com/14mdzk/dev_finance/internal/pkg/handler"
	"github.com/14mdzk/dev_finance/internal/pkg/reason"
	"github.com/gin-gonic/gin"
)

type TransactionService interface {
	BrowseAll(userID int, transactionTypeID int) ([]schema.TransactionResp, error)
	Create(userID int, req schema.TransactionReq) error
	Update(userID int, transactionID int, req schema.TransactionReq) error
	Delete(userID int, transactionID int) error
}

type TransactionController struct {
	transactionSvc TransactionService
}

func NewTransactionController(transactionSvc TransactionService) *TransactionController {
	return &TransactionController{
		transactionSvc: transactionSvc,
	}
}

func (ctrl *TransactionController) BrowseTransaction(ctx *gin.Context) {
	sub, _ := ctx.Get("user_id")

	intSub, err := strconv.Atoi(fmt.Sprint(sub))
	if err != nil {
		handler.ResponseError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	transactionType := ctx.Query("transaction_type")
	transactionTypeID, _ := strconv.Atoi(transactionType)

	resp, err := ctrl.transactionSvc.BrowseAll(intSub, transactionTypeID)
	if err != nil {
		handler.ResponseError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, reason.CurrencySuccessBrowse, resp)
}

func (ctrl *TransactionController) CreateTransaction(ctx *gin.Context) {
	sub, _ := ctx.Get("user_id")

	intSub, err := strconv.Atoi(fmt.Sprint(sub))
	if err != nil {
		handler.ResponseError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var req schema.TransactionReq

	if !handler.BindAndValidate(ctx, &req) {
		return
	}

	err = ctrl.transactionSvc.Create(intSub, req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, reason.CurrencySuccessBrowse, nil)
}

func (ctrl *TransactionController) UpdateTransaction(ctx *gin.Context) {
	sub, _ := ctx.Get("user_id")

	intSub, err := strconv.Atoi(fmt.Sprint(sub))
	if err != nil {
		handler.ResponseError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	transactionID, err := strconv.Atoi(fmt.Sprint(ctx.Param("id")))
	if err != nil || transactionID <= 0 {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.InvalidRequest)
		return
	}

	var req schema.TransactionReq

	if !handler.BindAndValidate(ctx, &req) {
		return
	}

	err = ctrl.transactionSvc.Update(intSub, transactionID, req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, reason.CurrencySuccessBrowse, nil)
}

func (ctrl *TransactionController) DeleteTransaction(ctx *gin.Context) {
	sub, _ := ctx.Get("user_id")

	intSub, err := strconv.Atoi(fmt.Sprint(sub))
	if err != nil {
		handler.ResponseError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	transactionID, err := strconv.Atoi(fmt.Sprint(ctx.Param("id")))
	if err != nil || transactionID <= 0 {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.InvalidRequest)
		return
	}

	err = ctrl.transactionSvc.Delete(intSub, transactionID)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.InvalidRequest)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, reason.CurrencySuccessBrowse, nil)
}
