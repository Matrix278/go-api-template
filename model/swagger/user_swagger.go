package swagger

import "time"

// UserSwagger represents the user model for Swagger documentation.
type UserSwagger struct {
	ID        string     `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Username  string     `json:"username" example:"JohnDoe"`
	Email     string     `json:"email" example:"test@test.com"`
	CreatedAt time.Time  `json:"created_at" example:"2021-01-01T00:00:00Z"`
	UpdatedAt *time.Time `json:"updated_at" example:"2021-01-01T00:00:00Z"`
}
