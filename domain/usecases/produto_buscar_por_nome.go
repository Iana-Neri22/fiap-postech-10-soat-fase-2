package usecases

import (
	"context"
	"fmt"
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
)

type ProdutoBuscaPorNomeUseCase struct {
	produtoRepo repository.ProdutoRepository
}

func NewProdutoBuscaPorNomeUseCase(produtoRepo repository.ProdutoRepository) *ProdutoBuscaPorNomeUseCase {
	return &ProdutoBuscaPorNomeUseCase{
		produtoRepo: produtoRepo,
	}
}

func (pd *ProdutoBuscaPorNomeUseCase) Run(c context.Context, nome string) (*entities.Produto, error) {
	Produto, err := pd.produtoRepo.ProdutoBuscarPorNome(c, nome)
	if err != nil {
		return nil, fmt.Errorf("não foi possível buscar produto: %w", err)
	}
	return Produto, nil
}
