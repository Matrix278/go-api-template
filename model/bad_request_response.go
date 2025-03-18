package model

type BadRequestResponse struct {
	Message string `json:"message" example:"Bad request"`
	Code    string `json:"code" example:"BAD_REQUEST"`
	Errors  []struct {
		Field   string `json:"field" example:"name"`
		Message string `json:"message" example:"Name is required"`
	} `json:"errors"`
}
