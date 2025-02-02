package model

type BadRequestWithValidationErrorsResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Errors  []struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	} `json:"errors"`
}
