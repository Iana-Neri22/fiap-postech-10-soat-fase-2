package repository

import "context"
import "lanchonete/domain/entities"

type AcompanhamentoRepository interface {
	CriarAcompanhamento(c context.Context, acompanhamento *entities.AcompanhamentoPedido) error
	BuscarPedidos(c context.Context, nroPedido string) (entities.Pedido, error)
}
