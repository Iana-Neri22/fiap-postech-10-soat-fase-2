package usecases

import (
	"context"
	"lanchonete/domain/entities"
)

type AcompanhamentoUseCase interface {
	CriarAcompanhamento(c context.Context, acompanhamento *entities.AcompanhamentoPedido) error
	BuscarPedidos(c context.Context, nroPedido string) (entities.Pedido, error)
	AdicionarPedido(c context.Context, acompanhamento *entities.AcompanhamentoPedido, pedido *entities.Pedido) error
}