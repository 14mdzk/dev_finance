package model

import "time"

type User struct {
	ID        int        `db:"id"`
	Username  string     `db:"username"`
	Password  string     `db:"password,omitempty"`
	FullName  string     `db:"full_name"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}
