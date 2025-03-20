package entities

import (
	"errors"
	"strings"

)

type CatProduto string

const (
	Lanche         CatProduto = "Lanche"
	Acompanhamento CatProduto = "Acompanhamento"
	Bebida         CatProduto = "Bebida"
)

type Produto struct {
	Identificacao string
	Nome          string
	Categoria     CatProduto
	Preco         float32
}

func ProdutoNew(identificacao string, nome string, categoria string, preco float32) (*Produto, error) {
	if strings.TrimSpace(identificacao) == "" || strings.TrimSpace(nome) == "" || strings.TrimSpace(categoria) == "" {
		return nil, errors.New("todos os campos são obrigatórios")
	}

	var cat_prod CatProduto

	switch CatProduto(categoria) {
	case Lanche, Acompanhamento, Bebida:
		cat_prod = CatProduto(categoria)
	default:
		return nil, errors.New("categoria inválida")
	}

	// if err != nil {
	// 	return nil, errors.New("preço deve ser um número válido")
	// }

	return &Produto{
		Identificacao: identificacao,
		Nome:          nome,
		Categoria:     cat_prod,
		Preco:         preco,
	}, nil
}
