package middleware

import (
	"net/http"
	"strings"

	"github.com/14mdzk/dev_finance/internal/pkg/handler"
	"github.com/14mdzk/dev_finance/internal/pkg/reason"
	"github.com/gin-gonic/gin"
)

type AccessTokenVerifier interface {
	VerifyAccessToken(tokenString string) (string, error)
}

func SessionMiddleware(tokenMaker AccessTokenVerifier) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := tokenFromHeader(ctx)
		if accessToken == "" {
			handler.ResponseError(ctx, http.StatusUnauthorized, reason.Unauthenticated)
			ctx.Abort()
			return
		}

		sub, err := tokenMaker.VerifyAccessToken(accessToken)
		if err != nil {
			handler.ResponseError(ctx, http.StatusUnauthorized, reason.Unauthenticated)
			ctx.Abort()
			return
		}

		ctx.Set("user_id", sub)
		ctx.Next()
	}
}

func tokenFromHeader(ctx *gin.Context) string {
	accessToken := ""

	bearerToken := ctx.Request.Header.Get("Authorization")
	fields := strings.Fields(bearerToken)

	if len(fields) == 2 && fields[0] == "Bearer" {
		accessToken = fields[1]
	}

	return accessToken
}
