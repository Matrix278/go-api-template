package swagger

type StatusBadRequest struct {
	Message string `json:"message" example:"Bad request"`
	Code    string `json:"code" example:"BAD_REQUEST"`
}
