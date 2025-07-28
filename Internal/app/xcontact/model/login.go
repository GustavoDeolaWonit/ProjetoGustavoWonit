package model

type Login struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Senha    string `json:"senha"`
	Token    string `json:"token"`
	ExpiraEm int    `json:"expira_em"`
}

type LoginRequestAPI struct {
	Nome  string `json:"nome"`
	Senha string `json:"senha"`
}
