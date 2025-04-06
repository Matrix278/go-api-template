package model

type ForbiddenResponse struct {
	Message string `example:"You don't have permission to access this resource" json:"message"`
	Code    string `example:"FORBIDDEN"                                         json:"code"`
}
