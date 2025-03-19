package usecases

import (
	"context"
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
)

type PedidoBuscarPorIdUseCase struct {
	pedidoGateway repository.PedidoRepository
}

func (pduc *PedidoBuscarPorIdUseCase) Run(c context.Context, identificacao string) (entities.Pedido, error) {
	pedido, err := pduc.pedidoGateway.BuscarPedido(c, identificacao)
	if err != nil {
		return entities.Pedido{}, err
	}
	return pedido, nil
}
