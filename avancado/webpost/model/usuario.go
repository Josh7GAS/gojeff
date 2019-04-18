package model

//Usuario representa um usuario do Sistema
type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}
