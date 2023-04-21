package dto

type (
	LoginRequest struct {
		Username string `json:"username" validate:"required" example:"username"`
		Password string `json:"password" validate:"required" example:"password"`
	}
	LoginResponse struct {
		Token string `json:"token" example:"Y29weSB0byBjbGlwYm9hcmQuCg=="`
	}
)
