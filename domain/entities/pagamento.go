package entities

import (
	"errors"
	"strings"
)

type Pagamento struct {
	ID  string
	Valor float64
	Status   string
	DataCriacao string
}

func PagamentoNew(id string, valor float64, status string, dataCriacao string) (*Pagamento, error) {

	if  strings.TrimSpace(id) == "" || strings.TrimSpace(status) == "" || strings.TrimSpace(dataCriacao) == "" {
		return nil, errors.New("nenhum dos campos podem estar em branco")
	}	

	return &Pagamento{
		ID:  id,
		Valor: valor,
		Status: status,	
		DataCriacao: dataCriacao,
	}, nil
}