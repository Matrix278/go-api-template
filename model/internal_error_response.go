package model

type InternalErrorResponse struct {
	Message string `example:"Internal server error" json:"message"`
	Code    string `example:"INTERNAL_ERROR"        json:"code"`
}
