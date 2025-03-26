package presenters

type ClienteDTO struct {
	CPF      string `json:"cpf"`
	Nome     string `json:"nome"`
	Endereco string `json:"endereco"`
}
