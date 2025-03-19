package entities

import (
	"errors"
	"time"

	"github.com/shopspring/decimal"
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
	Identificacao     string
	Produtos          []Produto
	Status            StatusPedido
	TimeStamp         time.Time
	UltimaAtualizacao time.Time
	Total             decimal.Decimal
}

func PedidoNew(identificacao string, produtos []Produto) (*Pedido, error) {
	if identificacao == "" || len(produtos) == 0 {
		return nil, errors.New("todos os campos são obrigatórios")
	}

	var total decimal.Decimal
	for _, produto := range produtos {
		total = total.Add(produto.Preco)
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
