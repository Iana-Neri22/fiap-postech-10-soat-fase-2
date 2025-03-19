package usecase

import (
	"context"
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
)

type produtoUseCase struct {
	produtoRepository repository.ProdutoRepository
}

func NewProdutoUseCase(produtoRepository repository.ProdutoRepository) repository.ProdutoRepository {
	return &produtoUseCase{
		produtoRepository: produtoRepository,
	}
}

func (puc *produtoUseCase) AdicionarProduto(c context.Context, produto *entities.Produto) error {
	return puc.produtoRepository.AdicionarProduto(c, produto)
}

func (puc *produtoUseCase) BuscarProdutoPorId(c context.Context, Id string) (*entities.Produto, error) {
	return puc.produtoRepository.BuscarProdutoPorId(c, Id)
}
