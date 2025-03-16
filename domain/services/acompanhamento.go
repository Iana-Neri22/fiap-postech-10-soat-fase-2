package service

import (
	"context"
	"lanchonete/domain/entities"
)

const (
	CollectionPedido = "pedido"
)

type AcompanhamentoUseCase interface {
	CriarAcompanhamento(C context.Context, acompanhamento *entities.AcompanhamentoPedido) error
	BuscarPedido(C context.Context, ID string) (entities.Pedido, error)
	AdicionarPedido(C context.Context, pedido *entities.Pedido) error
}
