package usecase

import (
	"context"
	
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
	"lanchonete/domain/usecases"
)

type pagamentoUseCase struct {
	pagamentoRepository repository.PagamentoRepository
}

func NewPagamentoUseCase(pagamentoRepository repository.PagamentoRepository) usecases.PagamentoUseCase {
	return &pagamentoUseCase{
		pagamentoRepository: pagamentoRepository,
	}
}

func (puc *pagamentoUseCase) EnviarPagamento(c context.Context, pagamento *entities.Pagamento) error {
	return puc.pagamentoRepository.EnviarPagamento(c, pagamento)
}



