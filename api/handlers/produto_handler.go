// <<<<<<< HEAD
// package handler

// import (
// 	_ "lanchonete/docs"
// 	"lanchonete/domain/entities"
// 	"lanchonete/domain/responses"
// 	"lanchonete/domain/usecases"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type ProdutoHandler struct {
// 	ProdutoIncluirUseCase usecases.ProdutoIncluirUseCase
// 	ProdutoBuscarPorIdUseCase usecases.ProdutoBuscaPorIdUseCase
// }

// func (ph *ProdutoHandler) ProdutoIncluir(c *gin.Context) {
// 	var produto entities.Produto

// 	err := c.ShouldBind(&produto)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
// 		return
// 	}


// 	prd, err := ph.ProdutoIncluirUseCase.Run(c, produto.Identificacao, produto.Nome, string(produto.Categoria), produto.Preco)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, response.SuccessResponse{
// 		Message: "Produto incluído com sucesso: " + prd.Nome,
// 	})
// }

// func (ph *ProdutoHandler) ProdutoBuscarPorId(c *gin.Context) {
// 	id := c.Param("id")

// 	prd, err := ph.ProdutoBuscarPorIdUseCase.Run(c, id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, prd)
// }

// =======
package handler

import (
	_ "lanchonete/docs"
	"lanchonete/domain/entities"
	response "lanchonete/domain/responses"
	"lanchonete/domain/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProdutoHandler struct {
	ProdutoIncluirUseCase     usecases.ProdutoIncluirUseCase
	ProdutoBuscarPorIdUseCase usecases.ProdutoBuscaPorIdUseCase
}

// CriarProduto godoc
// @Summary Cria um produto
// @Description Cria um produto
// @Tags produto
// @Router /produto [post]
// @Accept  json
// @Produce  json
// @Param cliente body entities.Produto true "Produto"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
func (ph *ProdutoHandler) ProdutoIncluir(c *gin.Context) {
	var produto entities.Produto

	err := c.ShouldBind(&produto)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	prd, err := ph.ProdutoIncluirUseCase.Run(c, produto.Identificacao, produto.Nome, string(produto.Categoria), produto.Descricao, produto.Preco)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Message: "Produto incluído com sucesso: " + prd.Nome,
	})
}

// BuscarProduto godoc
// @Summary Busca um produto
// @Description Busca um produto
// @Tags produto
// @Router /produto/{id} [get]
// @Accept  json
// @Produce  json
// @Param id path string true "id do produto"
// @Success 200 {object} entities.Produto
// @Failure 400 {object} response.ErrorResponse
func (ph *ProdutoHandler) ProdutoBuscarPorId(c *gin.Context) {
	id := c.Param("id")

	prd, err := ph.ProdutoBuscarPorIdUseCase.Run(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, prd)
}
// >>>>>>> Bruno
