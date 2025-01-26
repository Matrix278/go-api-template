package model

import (
	"time"

	"github.com/go-openapi/strfmt"
)

type User struct {
	ID        strfmt.UUID4 `json:"id"`
	Username  string       `json:"username"`
	Email     string       `json:"email"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}
