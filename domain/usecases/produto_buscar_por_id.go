package usecases

import (
	"context"
	"lanchonete/domain/entities"
)

/*
type ProdutoBuscaPorIdUseCase struct {
	produtoRepo repository.ProdutoRepository
}

func (pd *ProdutoBuscaPorIdUseCase) Run(c context.Context, id string) (*entities.Produto, error) {
	produto, err := pd.produtoRepo.BuscarProdutoPorId(c, id)

	if err != nil {
		return nil, fmt.Errorf("não foi possível buscar produto: %w", err)
	}
	return produto, nil
}
	"lanchonete/domain/entities"
)*/

type ProdutoBuscaPorIdUseCase interface {
	Run(c context.Context, id string) (*entities.Produto, error)
}
