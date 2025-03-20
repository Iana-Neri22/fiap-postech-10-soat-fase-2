package repository

import (
	"context"
	"lanchonete/domain/entities"

	"github.com/google/uuid"
)

type PedidoRepository interface {
	CriarPedido(c context.Context, pedido *entities.Pedido) error
	BuscarPedido(c context.Context, identificacao uuid.UUID) (*entities.Pedido, error)
	AtualizarStatusPedido(c context.Context, identificacao string, status string) error
}
