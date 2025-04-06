package model

type PartialSuccessResponse struct {
	Message string `example:"Some of the requested operations were successful" json:"message"`
	Code    string `example:"PARTIAL_SUCCESS"                                  json:"code"`
}
