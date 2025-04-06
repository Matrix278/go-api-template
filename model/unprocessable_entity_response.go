package model

type UnprocessableEntityResponse struct {
	Message string `example:"Posted data is not valid" json:"message"`
	Code    string `example:"UNPROCESSABLE_ENTITY"     json:"code"`
}
