package dto

type LoginRequest struct {
	Email string `json:"email" binding:"required" example:"guszinho@gmail.com"`
	Senha string `json:"senha" binding:"required" example:"123456"`
}
