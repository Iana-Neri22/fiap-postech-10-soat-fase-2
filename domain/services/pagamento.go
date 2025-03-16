package service

import (
	"context"
	"lanchonete/domain/entities"
)

const (
	CollectionPagamento = "pagamento"
)

type PagamentoUseCase interface {
	EnviarPagamento(C context.Context, pagamento *entities.Pagamento) error
	ConfirmarPagamento(C context.Context, pagamento *entities.Pagamento) error
}

