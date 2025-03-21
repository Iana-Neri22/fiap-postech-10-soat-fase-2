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

func (pduc *PedidoIncluirUseCase) Run(c context.Context, cliente entities.Cliente, produtos []entities.Produto, personalizacao string) (*entities.Pedido, error) {
	/*
		for _, id := range produtos {
			_, err := pduc.pedidoRepository.BuscarProdutoPorId(c, id)
			if err != nil {
				return nil, errors.New("produto não encontrado")
			}
		}*/

	pedido, err := entities.PedidoNew(cliente, produtos, personalizacao)
	if err != nil {
		return nil, err
	}
	err = pduc.pedidoRepository.CriarPedido(c, pedido)

	if err != nil {
		return nil, err
	}
	return pedido, nil
}
