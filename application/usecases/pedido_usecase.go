package usecase

import (
	"context"
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
	"lanchonete/domain/services"
)

type pedidoUseCase struct {
	pedidoRepo repository.PedidoRepository
}

func NewPedidoUseCase(pedidoRepo repository.PedidoRepository) service.PedidoUseCase {
	return &pedidoUseCase{
		pedidoRepo: pedidoRepo,
	}
}

func (uc *pedidoUseCase) CriarPedido(c context.Context, pedido *entities.Pedido) error {
	return uc.pedidoRepo.CriarPedido(c, pedido)
}

func (uc *pedidoUseCase) BuscarPedido(c context.Context, pedidoID string) (entities.Pedido, error) {
	return uc.pedidoRepo.BuscarPedido(c, pedidoID)
}
