package usecases

import (
	"context"
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
)

type PedidoIncluirUseCase struct {
	pedidoGateway repository.PedidoRepository
}

func (pduc *PedidoIncluirUseCase) Run(c context.Context, cliente entities.Cliente, produtos []entities.Produto) (*entities.Pedido, error) {

	pedido, err := entities.PedidoNew(cliente, produtos)
	if err != nil {
		return nil, err
	}
	err = pduc.pedidoGateway.CriarPedido(c, pedido)
	if err != nil {
		return nil, err
	}
	return pedido, nil
}
