package service

import (
	"context"
	"lanchonete/domain/entities"
)

const (
	CollectionAcompanhamento = "acompanhamento"
)

type AcompanhamentoUseCase interface {
	CriarAcompanhamento(C context.Context, acompanhamento *entities.AcompanhamentoPedido) error
	BuscarPedido(C context.Context, nroPedido string) (entities.Pedido, error)
	AdicionarPedido(C context.Context, acompanhamento *entities.AcompanhamentoPedido, pedido *entities.Pedido) error
	BuscarAcompanhamento(C context.Context, nroPedido string) (entities.AcompanhamentoPedido, error)
	AtualizarStatusPedido(C context.Context, acompanhamento *entities.AcompanhamentoPedido, identificacao string, novoStatus entities.StatusPedido) error
}
