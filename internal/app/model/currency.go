package model

import "time"

type Currency struct {
	ID           int        `db:"id"`
	Name         string     `db:"name"`
	Abbreviation string     `db:"abbreviation"`
	CreatedAt    time.Time  `db:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at"`
}
