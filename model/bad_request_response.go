package model

type BadRequestResponse struct {
	Message string `example:"Bad request" json:"message"`
	Code    string `example:"BAD_REQUEST" json:"code"`
	Errors  []struct {
		Field   string `example:"name"             json:"field"`
		Message string `example:"Name is required" json:"message"`
	} `json:"errors"`
}
