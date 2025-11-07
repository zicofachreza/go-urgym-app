package dto

type LoginRequest struct {
	EmailOrUsername string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
}
