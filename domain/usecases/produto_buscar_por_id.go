package usecases

import (
	"context"
	"lanchonete/domain/entities"
)

type ProdutoBuscaPorIdUseCase interface {
	Run(c context.Context, id string) (*entities.Produto, error)
}
