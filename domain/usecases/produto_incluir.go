package usecases

import (
	"context"
	/*	"fmt"
			"lanchonete/domain/entities"
			"lanchonete/domain/repository"

			"github.com/shopspring/decimal"
		)

		type ProdutoIncluirUseCase struct {
			produtoRepo repository.ProdutoRepository
		}

		func (pd *ProdutoIncluirUseCase) Run(c context.Context, identificacao string, nome string, categoria string, preco string) (*entities.Produto, error) {
			Preco, err := decimal.NewFromString(preco)
			if err != nil {
				return nil, fmt.Errorf("preço deve ser um número válido: %w", err)
			}
			produto, err := entities.ProdutoNew(identificacao, nome, categoria, Preco)

			if err != nil {
				return nil, fmt.Errorf("criação de produto inválida: %w", err)
			}

			err = pd.produtoRepo.AdicionarProduto(c, produto)
			if err != nil {
				return nil, fmt.Errorf("não foi possível criar produto: %w", err)

			}

			return produto, nil
	*/
	"lanchonete/domain/entities"
)

type ProdutoIncluirUseCase interface {
	Run(c context.Context, identificacao string, nome string, categoria string, preco string) (*entities.Produto, error)
}
