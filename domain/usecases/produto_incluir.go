package usecases

import (
	"context"
	"lanchonete/domain/entities"
)

type ProdutoIncluirUseCase interface {
	Run(c context.Context, identificacao string, nome string, categoria string, preco string) (*entities.Produto, error)
}