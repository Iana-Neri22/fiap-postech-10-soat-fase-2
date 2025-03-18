package service

import (
	"context"
	"lanchonete/domain/entities"
)

type AcompanhamentoUseCase interface {
	CriarAcompanhamento(C context.Context, acompanhamento *entities.AcompanhamentoPedido) error
	BuscarPedido(C context.Context, nroPedido string) (entities.Pedido, error)
	AdicionarPedido(C context.Context, acompanhamento *entities.AcompanhamentoPedido, pedido *entities.Pedido) error
}
