package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
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
	Identificacao     uuid.UUID
	Cliente           Cliente
	Produtos          []Produto
	Status            StatusPedido
	TimeStamp         time.Time
	UltimaAtualizacao time.Time
	Total             float32
}

func PedidoNew(cliente Cliente, produtos []Produto) (*Pedido, error) {
	id := uuid.New()

	if len(produtos) == 0 {
		return nil, errors.New("o pedido precisa ter ao menos um produto")
	}

	temLanche := false
	for _, produto := range produtos {
		if produto.Categoria == Lanche {
			temLanche = true
			break
		}
	}
	if !temLanche {
		return nil, errors.New("o pedido precisa ter ao menos um lanche")
	}

	total := float32(0)
	for _, produto := range produtos {
		total += produto.Preco
	}

	return &Pedido{
		Identificacao:     id,
		Cliente:           cliente,
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
