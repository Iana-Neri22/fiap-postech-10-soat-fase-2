package handler

import (
	_ "lanchonete/docs"
	"lanchonete/domain/entities"
	"lanchonete/domain/responses"
	"lanchonete/domain/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProdutoHandler struct {
	ProdutoIncluirUseCase usecases.ProdutoIncluirUseCase
	ProdutoBuscarPorIdUseCase usecases.ProdutoBuscaPorIdUseCase
}

func (ph *ProdutoHandler) ProdutoIncluir(c *gin.Context) {
	var produto entities.Produto

	err := c.ShouldBind(&produto)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}


	prd, err := ph.ProdutoIncluirUseCase.Run(c, produto.Identificacao, produto.Nome, string(produto.Categoria), produto.Preco.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Message: "Produto incluído com sucesso: " + prd.Nome,
	})
}

func (ph *ProdutoHandler) ProdutoBuscarPorId(c *gin.Context) {
	id := c.Param("id")

	prd, err := ph.ProdutoBuscarPorIdUseCase.Run(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, prd)
}

