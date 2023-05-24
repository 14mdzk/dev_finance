package middleware

import (
	"net/http"

	"github.com/14mdzk/dev_finance/internal/pkg/handler"
	"github.com/14mdzk/dev_finance/internal/pkg/reason"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func Authorize(obj string, act string, enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sub, _ := ctx.Get("user_id")

		if sub != "1" {
			sub = "usr"
		}

		ok, err := enforce(sub.(string), obj, act, enforcer)

		if err != nil || !ok {
			handler.ResponseError(ctx, http.StatusForbidden, reason.Unauthorized)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func enforce(sub string, obj string, act string, enforcer *casbin.Enforcer) (bool, error) {
	err := enforcer.LoadPolicy()
	if err != nil {
		return false, err
	}
	ok, err := enforcer.Enforce(sub, obj, act)

	return ok, err
}
