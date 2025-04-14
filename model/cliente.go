package model

type Cliente struct {
	ID       int    `json:"id_cliente"`
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Telefone string `json:"telefone"`
	CPF      string `json:"cpf"`
	Senha 	 string `json:"senha"`
}