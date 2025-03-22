package repository

import (
	"context"
	"lanchonete/domain/entities"
)

type ProdutoRepository interface {
	AdicionarProduto(c context.Context, produto *entities.Produto) error
	BuscarProdutoPorId(c context.Context, nome string) (*entities.Produto, error)
	ListarTodosOsProdutos(c context.Context) ([]*entities.Produto, error)
}
