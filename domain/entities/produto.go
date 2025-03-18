package entities

import (
	"errors"
	"strings"

	"github.com/shopspring/decimal"
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
	Preco         decimal.Decimal
}

func ProdutoNew(identificacao string, nome string, categoria string, preco decimal.Decimal) (*Produto, error) {
	if strings.TrimSpace(identificacao) == "" || strings.TrimSpace(nome) == "" || strings.TrimSpace(preco.String()) == "" || strings.TrimSpace(categoria) == "" {
		return nil, errors.New("todos os campos são obrigatórios")
	}

	var cat_prod CatProduto

	switch CatProduto(categoria) {
	case Lanche, Acompanhamento, Bebida:
		cat_prod = CatProduto(categoria)
	default:
		return nil, errors.New("categoria inválida")
	}

	decimal, err := decimal.NewFromString(preco.String())

	if err != nil {
		return nil, errors.New("preço deve ser um número válido")
	}

	return &Produto{
		Identificacao: identificacao,
		Nome:          nome,
		Categoria:     cat_prod,
		Preco:         decimal,
	}, nil
}
