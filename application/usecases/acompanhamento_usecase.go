package usecase

import (
	"context"
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
	"lanchonete/domain/services"
)

type acompanhamentoUseCase struct {
	acompanhamentoRepo repository.AcompanhamentoRepository
}

func NewAcompanhamentoUseCase(acompanhamentoRepo repository.AcompanhamentoRepository) service.AcompanhamentoUseCase {
	return &acompanhamentoUseCase{
		acompanhamentoRepo: acompanhamentoRepo,
	}
}

func (uc *acompanhamentoUseCase) CriarAcompanhamento(c context.Context, acompanhamento *entities.AcompanhamentoPedido) error {
	return uc.acompanhamentoRepo.CriarAcompanhamento(c, acompanhamento)
}

func (uc *acompanhamentoUseCase) BuscarPedido(c context.Context, nroPedido string) (entities.Pedido, error) {
	return uc.acompanhamentoRepo.BuscarPedidos(c, nroPedido)
}

func (uc *acompanhamentoUseCase) AdicionarPedido(c context.Context, acompanhamento *entities.AcompanhamentoPedido, pedido *entities.Pedido) error {
	return uc.acompanhamentoRepo.AdicionarPedido(c, acompanhamento, pedido)
}

func (uc *acompanhamentoUseCase) BuscarAcompanhamento(c context.Context, nroPedido string) (entities.AcompanhamentoPedido, error) {
    return uc.acompanhamentoRepo.BuscarAcompanhamento(c, nroPedido)
}
