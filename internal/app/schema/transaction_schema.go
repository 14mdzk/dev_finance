package schema

import (
	"time"

	"github.com/14mdzk/dev_finance/internal/pkg/helper"
)

type TransactionReq struct {
	TransactionTypeID     int                  `json:"transaction_type_id"`
	TransactionCategoryID int                  `json:"transaction_category_id"`
	TransactionDate       helper.SqlDateFormat `json:"transaction_date"`
	CurrencyID            int                  `json:"currency_id"`
	Amount                int                  `json:"amount"`
	Note                  string               `json:"note"`
}

type TransactionResp struct {
	ID                  int `json:"id"`
	TransactionType     TransactionTypeResp
	TransactionCategory TransactionCategoryResp
	Currency            CurrencyResp
	TransactionDate     time.Time `json:"transaction_date"`
	Amount              int       `json:"amount"`
	Note                string    `json:"note"`
}
