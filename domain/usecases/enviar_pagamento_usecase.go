package usecases

import (
	"context"
	"lanchonete/domain/entities"
)

type PagamentoUseCase interface {
	EnviarPagamento(c context.Context, pagamento *entities.Pagamento) error
}