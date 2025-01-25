package swagger

type StatusInternalError struct {
	Message string `json:"message" example:"Internal server error"`
	Code    string `json:"code" example:"INTERNAL_ERROR"`
}
