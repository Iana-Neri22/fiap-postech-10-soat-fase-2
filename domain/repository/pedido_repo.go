package repository

import (
	"context"
	"lanchonete/domain/entities"
)

type PedidoRepository interface {
	CriarPedido(c context.Context, pedido *entities.Pedido) error
	BuscarPedido(c context.Context, nroPedido string) (entities.Pedido, error)
}
