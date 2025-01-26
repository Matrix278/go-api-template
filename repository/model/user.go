package model

import (
	"time"

	"github.com/go-openapi/strfmt"
)

type User struct {
	ID        strfmt.UUID4 `db:"id"`
	Username  string       `db:"username"`
	Email     string       `db:"email"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt *time.Time   `db:"updated_at"`
}
