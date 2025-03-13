package usecases

import (
	"context"
	"lanchonete/domain/entities"
)

type AcompanhamentoUseCase interface {
	BuscarPedidos(c context.Context, nroPedido string) (entities.Pedido, error)
	AcompanharPedido(c context.Context, pedido *entities.Pedido) error
}