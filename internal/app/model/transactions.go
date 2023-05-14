package model

import "time"

type Transaction struct {
	ID                    int        `db:"id"`
	UserID                int        `db:"user_id"`
	TransactionType       string     `db:"transaction_type"`
	TransactionTypeID     int        `db:"transaction_type_id"`
	TransactionCategory   string     `db:"transaction_category"`
	TransactionCategoryID int        `db:"transaction_category_id"`
	TransactionDate       string     `db:"transaction_date"`
	Currency              string     `db:"currency"`
	CurrencyID            int        `db:"currency_id"`
	Amount                int        `db:"amount"`
	Note                  string     `db:"note"`
	CreatedAt             time.Time  `db:"created_at"`
	UpdatedAt             *time.Time `db:"updated_at"`
}
