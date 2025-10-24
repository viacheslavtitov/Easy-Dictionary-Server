package domain

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SuccessIdResponse struct {
	Id int `json:"id"`
}
