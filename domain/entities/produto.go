package entities

import (
	"errors"
	"strings"
// <<<<<<< HEAD

// =======
// >>>>>>> Bruno
)

type CatProduto string

const (
	Lanche         CatProduto = "Lanche"
	Acompanhamento CatProduto = "Acompanhamento"
	Bebida         CatProduto = "Bebida"
	Sobremesa      CatProduto = "Sobremesa"
)

type Produto struct {
	Identificacao string
	Nome          string
	Categoria     CatProduto
// <<<<<<< HEAD
// 	Preco         float32
// }

// func ProdutoNew(identificacao string, nome string, categoria string, preco float32) (*Produto, error) {
// 	if strings.TrimSpace(identificacao) == "" || strings.TrimSpace(nome) == "" || strings.TrimSpace(categoria) == "" {
// =======
	Descricao     string
	Preco         float32
}

func ProdutoNew(identificacao string, nome string, categoria string, descricao string, preco float32) (*Produto, error) {
	if strings.TrimSpace(identificacao) == "" || strings.TrimSpace(nome) == "" || preco == 0 || strings.TrimSpace(categoria) == "" {
// >>>>>>> Bruno
		return nil, errors.New("todos os campos são obrigatórios")
	}

	var cat_prod CatProduto

	switch CatProduto(categoria) {
	case Lanche, Acompanhamento, Bebida, Sobremesa:
		cat_prod = CatProduto(categoria)
	default:
		return nil, errors.New("categoria inválida")
	}

// <<<<<<< HEAD
// 	// if err != nil {
// 	// 	return nil, errors.New("preço deve ser um número válido")
// 	// }

// =======
// >>>>>>> Bruno
	return &Produto{
		Identificacao: identificacao,
		Nome:          nome,
		Categoria:     cat_prod,
// <<<<<<< HEAD
// =======
		Descricao:     descricao,
// >>>>>>> Bruno
		Preco:         preco,
	}, nil
}
