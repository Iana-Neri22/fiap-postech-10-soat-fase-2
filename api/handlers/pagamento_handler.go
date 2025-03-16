package handler

import (
	_ "lanchonete/docs"
	"lanchonete/domain/entities"
	"lanchonete/domain/responses"
	"lanchonete/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"

)

type PagamentoHandler struct {
	PagamentoUseCase service.PagamentoUseCase
}

func (cc *PagamentoHandler) EnviarPagamento(c *gin.Context) {
	var pagamento entities.Pagamento



	err := c.ShouldBind(&pagamento)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	err = cc.PagamentoUseCase.EnviarPagamento(c, &pagamento)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Message: "Pagaamento enviado com sucesso",
	})
}


