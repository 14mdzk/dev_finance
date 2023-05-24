package middleware

import (
	"net/http"

	"github.com/14mdzk/dev_finance/internal/pkg/handler"
	"github.com/14mdzk/dev_finance/internal/pkg/reason"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				handler.ResponseError(ctx, http.StatusInternalServerError, reason.InternalServerError)
			}
		}()

		ctx.Next()
	}
}
