package usecases

import (
	"context"
	"lanchonete/internal/domain/entities"
	"lanchonete/internal/domain/repository"
)

type PedidoBuscarPorIdUseCase struct {
	pedidoRepository repository.PedidoRepository
}

func NewPedidoBuscarPorIdUseCase(pedidoRepository repository.PedidoRepository) *PedidoBuscarPorIdUseCase {
	return &PedidoBuscarPorIdUseCase{
		pedidoRepository: pedidoRepository,
	}
}

func (pduc *PedidoBuscarPorIdUseCase) Run(c context.Context, identificacao string) (*entities.Pedido, error) {

	pedido, err := pduc.pedidoRepository.BuscarPedido(c, identificacao)
	if err != nil {
		return nil, err
	}
	return pedido, nil
}
