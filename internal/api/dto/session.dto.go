package dto

type Session struct {
	ID        string `json:"id" validate:"required"`
	Username  string `json:"username" validate:"required"`
	StaffID   uint32 `json:"staffId" validate:"required"`
	CreatedAt int64  `json:"createdAt" validate:"required"`
}
