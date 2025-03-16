package repository

import "context"
import "lanchonete/domain/entities"

type PagamentoRepository interface {
	EnviarPagamento(c context.Context, pagamento *entities.Pagamento) error
}
