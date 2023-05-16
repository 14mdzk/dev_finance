package handler

import (
	"fmt"
	"net/http"

	"github.com/14mdzk/dev_finance/internal/app/schema"
	"github.com/14mdzk/dev_finance/internal/pkg/reason"
	"github.com/14mdzk/dev_finance/internal/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ResponseError(ctx *gin.Context, statusCode int, message string) {
	resp := ResponseBody{
		Status:  "failed",
		Message: message,
	}
	ctx.JSON(statusCode, resp)
}

func ResponseSuccess(ctx *gin.Context, statusCode int, message string, data interface{}, pagination interface{}) {
	resp := ResponseBody{
		Status:   "success",
		Message:  message,
		Data:     data,
		MetaData: pagination,
	}
	ctx.JSON(statusCode, resp)
}

func BindAndValidate(ctx *gin.Context, data interface{}) bool {
	err := ctx.ShouldBind(data)
	if err != nil {
		ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return false
	}

	isError := validator.Check(data)
	if isError {
		ResponseError(ctx, http.StatusUnprocessableEntity, reason.InvalidRequest)
		return false
	}

	return true
}

func BindPagination(ctx *gin.Context, data *schema.PaginationReq) {
	_ = ctx.ShouldBind(data)
	logrus.Error(fmt.Println(data))
	if data.Page <= 0 {
		data.Page = 1
	}

	if data.PageSize <= 0 {
		data.PageSize = 10
	}
}
