package usecase

import (
	"context"
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
	"lanchonete/domain/usecases"
)

type AcompanhamentoUseCase struct {
	acompanhamentoRepo repository.AcompanhamentoRepository
}

func NewAcompanhamentoUseCase(acompanhamentoRepo repository.AcompanhamentoRepository) usecases.AcompanhamentoUseCase {
	return &AcompanhamentoUseCase{
		acompanhamentoRepo: acompanhamentoRepo,
	}
}

func (uc *AcompanhamentoUseCase) BuscarPedidos(c context.Context, nroPedido string) (entities.Pedido, error) {
	return uc.acompanhamentoRepo.BuscarPedidos(c, nroPedido)
}

func (uc *AcompanhamentoUseCase) AcompanharPedido(c context.Context, pedido *entities.Pedido) error {
	return uc.acompanhamentoRepo.AcompanharPedido(c, pedido)
}
