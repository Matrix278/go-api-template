package model

type UnauthorizedResponse struct {
	Message string `example:"You are not authorized to access this resource" json:"message"`
	Code    string `example:"UNAUTHORIZED"                                   json:"code"`
}
