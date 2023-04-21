package dto

type SuccessResponse struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"OK"`
}

type ResultResponse struct {
	Result bool `json:"result" example:"true"`
}
