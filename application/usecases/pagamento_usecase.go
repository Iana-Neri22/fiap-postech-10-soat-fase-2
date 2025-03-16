package usecase

import (
	"context"
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
	"lanchonete/domain/services"
)

type pagamentoUseCase struct {
	pagamentoRepository repository.PagamentoRepository
}

func NewPagamentoUseCase(pagamentoRepository repository.PagamentoRepository) service.PagamentoUseCase {
	return &pagamentoUseCase{
		pagamentoRepository: pagamentoRepository,
	}
}

func (uc *pagamentoUseCase) EnviarPagamento(c context.Context, pagamento *entities.Pagamento) error {
	return uc.pagamentoRepository.EnviarPagamento(c, pagamento)
}

