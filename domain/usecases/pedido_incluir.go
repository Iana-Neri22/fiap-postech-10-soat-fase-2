package usecases

import (
	"context"
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
)

type PedidoIncluirUseCase struct {
	pedidoRepository repository.PedidoRepository
}

func (pduc *PedidoIncluirUseCase) Run(c context.Context, indentificacao string, produtos []entities.Produto) (*entities.Pedido, error) {
	pedido, err := entities.PedidoNew(indentificacao, produtos)
	if err != nil {
		return nil, err
	}
	err = pduc.pedidoRepository.CriarPedido(c, pedido)
	if err != nil {
		return nil, err
	}
	return pedido, nil
}
