package usecase

// import (
// 	"context"
// 	"fmt"
// 	"lanchonete/domain/entities"
// 	"lanchonete/domain/repository"
// 	"lanchonete/domain/usecases"
// )

// type ProdutoIncluirUseCase struct {
// 	produtoRepo repository.ProdutoRepository
// }

// func NewProdutoIncluirUseCase(produtoRepo repository.ProdutoRepository) usecases.ProdutoIncluirUseCase {
// 	return &ProdutoIncluirUseCase{
// 		produtoRepo: produtoRepo,
// 	}
// }

// func (pd *ProdutoIncluirUseCase) Run(c context.Context, identificacao string, nome string, categoria string, descricao string, preco float32) (*entities.Produto, error) {

// 	produto, err := entities.ProdutoNew(identificacao, nome, categoria, descricao, preco)

// 	if err != nil {
// 		return nil, fmt.Errorf("criação de produto inválida: %w", err)
// 	}

// 	err = pd.produtoRepo.AdicionarProduto(c, produto)
// 	if err != nil {
// 		return nil, fmt.Errorf("não foi possível criar produto: %w", err)

// 	}

// 	return produto, nil
// }
