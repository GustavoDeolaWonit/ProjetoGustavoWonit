package dto

type LoginResponse struct {
	Email    string `json:"email"`
	Token    string `json:"Token"`
	ExpiraEm int    `json:"expira_em"`
}
