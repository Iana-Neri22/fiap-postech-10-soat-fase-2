package usecases

import (
	"context"
	"lanchonete/domain/entities"

	"github.com/google/uuid"
)

type PedidoUseCase interface {
	CriarPedido(c context.Context, pedido *entities.Pedido) error
	BuscarPedido(c context.Context, IdDoPedido uuid.UUID) (*entities.Pedido, error)
}
