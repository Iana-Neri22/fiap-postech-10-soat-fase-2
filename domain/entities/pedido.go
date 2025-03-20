package entities

import (
	"errors"
	"time"

)

type StatusPedido string

const (
	Pendente     StatusPedido = "Pendente"
	Recebido     StatusPedido = "Recebido"
	EmPreparacao StatusPedido = "Em preparação"
	Pronto       StatusPedido = "Pronto"
	Finalizado   StatusPedido = "Finalizado"
)

type Pedido struct {
	CPF				  string			`json:"cpf"`	
	Identificacao     string			`json:"identificacao"`
	Produtos          []Produto			`json:"produtos"`
	Status            StatusPedido		`json:"status"`
	TimeStamp         time.Time			`json:"timeStamp" format:"date-time"`
	UltimaAtualizacao time.Time			`json:"ultimaAtualizacao" format:"date-time"`
	Total             float32			`json:"total"`
}

func PedidoNew(identificacao string, produtos []Produto) (*Pedido, error) {
	if identificacao == "" || len(produtos) == 0 {
		return nil, errors.New("todos os campos são obrigatórios")
	}

	var total float32
	for _, produto := range produtos {
		total += produto.Preco
	}

	return &Pedido{
		Identificacao:     identificacao,
		Produtos:          produtos,
		Status:            Pendente,
		TimeStamp:         time.Now(),
		UltimaAtualizacao: time.Now(),
		Total:             total,
	}, nil
}

func (p *Pedido) UpdateStatus(status StatusPedido) error {
	switch status {
	case Recebido, EmPreparacao, Pronto, Finalizado:
		p.Status = status
	default:
		return errors.New("status inválido")
	}
	p.UltimaAtualizacao = time.Now()

	return nil
}
