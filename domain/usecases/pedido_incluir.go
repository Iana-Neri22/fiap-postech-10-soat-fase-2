package usecases

import (
	"context"
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
)

type PedidoIncluirUseCase struct {
	pedidoRepository repository.PedidoRepository
}

func NewPedidoIncluirUseCase(pedidoRepository repository.PedidoRepository) *PedidoIncluirUseCase {
	return &PedidoIncluirUseCase{
		pedidoRepository: pedidoRepository,
	}
}

func (pduc *PedidoIncluirUseCase) Run(c context.Context, cliente entities.Cliente, produtos []entities.Produto, personalizacao string) (string, error) {
	pedido, err := entities.PedidoNew(cliente, produtos, personalizacao)
	if err != nil {
		return "", err
	}
	err = pduc.pedidoRepository.CriarPedido(c, pedido)
	if err != nil {
		return "", err
	}
	return pedido.Identificacao, nil
}
