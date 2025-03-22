package repository

import (
	"context"
	"lanchonete/domain/entities"
)

type PedidoRepository interface {
	CriarPedido(c context.Context, pedido *entities.Pedido) error
// <<<<<<< HEAD
	BuscarPedido(c context.Context, identificacao string) (entities.Pedido, error)
// =======
// 	BuscarPedido(c context.Context, identificacao string) (*entities.Pedido, error)
// >>>>>>> Bruno
	AtualizarStatusPedido(c context.Context, identificacao string, status string) error
}
