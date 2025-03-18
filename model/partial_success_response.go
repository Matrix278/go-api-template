package model

type PartialSuccessResponse struct {
	Message string `json:"message" example:"Some of the requested operations were successful"`
	Code    string `json:"code" example:"PARTIAL_SUCCESS"`
}
