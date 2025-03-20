package swagger

type UnprocessableEntityResponse struct {
	Message string `json:"message" example:"Posted data is not valid"`
	Code    string `json:"code" example:"UNPROCESSABLE_ENTITY"`
}
