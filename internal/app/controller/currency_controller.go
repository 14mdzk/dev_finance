package controller

import (
	"fmt"

	"github.com/14mdzk/dev_finance/internal/app/schema"
	"github.com/gin-gonic/gin"
)

type CurrencyService interface {
	BrowseAll() ([]schema.CurrencyResp, error)
}

type CurrencyController struct {
	service CurrencyService
}

func NewCurrencyController(service CurrencyService) *CurrencyController {
	return &CurrencyController{
		service: service,
	}
}

func (ctrl *CurrencyController) BrowseAll(ctx *gin.Context) {
	resp, err := ctrl.service.BrowseAll()
	if err != nil {
		return
	}
	fmt.Println(resp)
}
