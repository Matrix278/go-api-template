package model

import "time"

// UserSwagger represents the user model for Swagger documentation.
type UserSwagger struct {
	ID        string     `example:"550e8400-e29b-41d4-a716-446655440000" json:"id"`
	Username  string     `example:"JohnDoe"                              json:"username"`
	Email     string     `example:"test@test.com"                        json:"email"`
	CreatedAt time.Time  `example:"2021-01-01T00:00:00Z"                 json:"created_at"`
	UpdatedAt *time.Time `example:"2021-01-01T00:00:00Z"                 json:"updated_at"`
}
