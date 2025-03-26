package presenters

import "lanchonete/domain/entities"

type ProdutoDTO struct {
	Identificacao string  `json:"identificacao"`
	Nome          string  `json:"nome"`
	Preco         float32 `json:"preco"`
	Categoria     string  `json:"categoria"`
	Descricao     string  `json:"descricao"`
}

func NewProdutoDTO() *ProdutoDTO {
	return &ProdutoDTO{}
}

func (p *ProdutoDTO) ToDTO(produto *entities.Produto) *ProdutoDTO {
	p.Identificacao = produto.Identificacao
	p.Nome = produto.Nome
	p.Preco = produto.Preco
	p.Categoria = string(produto.Categoria)
	p.Descricao = produto.Descricao

	return p
}

func (p *ProdutoDTO) ToEntity() *entities.Produto {
	return &entities.Produto{
		Identificacao: p.Identificacao,
		Nome:          p.Nome,
		Preco:         p.Preco,
		Categoria:     entities.CatProduto(p.Categoria),
		Descricao:     p.Descricao,
	}
}

type ProdutoDTOs []ProdutoDTO

func (p ProdutoDTOs) ConvertProdutoDTOsToEntities() []entities.Produto {
	var produtos []entities.Produto
	for _, dto := range p {
		produto := dto.ToEntity()
		produtos = append(produtos, *produto)
	}
	return produtos
}
