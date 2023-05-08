package handler

import (
	"net/http"

	"github.com/14mdzk/dev_finance/internal/pkg/reason"
	"github.com/14mdzk/dev_finance/internal/pkg/validator"
	"github.com/gin-gonic/gin"
)

func ResponseError(ctx *gin.Context, statusCode int, message string) {
	resp := ResponseBody{
		Status:  "failed",
		Message: message,
	}
	ctx.JSON(statusCode, resp)
}

func ResponseSuccess(ctx *gin.Context, statusCode int, message string, data interface{}) {
	resp := ResponseBody{
		Status:  "success",
		Message: message,
		Data:    data,
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
