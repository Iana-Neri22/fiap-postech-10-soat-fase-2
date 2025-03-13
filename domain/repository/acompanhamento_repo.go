package repository

import "context"
import "lanchonete/domain/entities"

type AcompanhamentoRepository interface {
	BuscarPedidos(c context.Context, nroPedido string) (entities.AcompanhamentoPedido, error)
	AcompanharPedido(c context.Context, pedido *entities.AcompanhamentoPedido) error
}
