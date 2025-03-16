package handler

import (
	_ "lanchonete/docs"
	"lanchonete/domain/entities"
	"lanchonete/domain/responses"
	"lanchonete/domain/services"
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"

)

type PagamentoHandler struct {
	PagamentoUseCase service.PagamentoUseCase
}

func (ph *PagamentoHandler) EnviarPagamento(c *gin.Context) {
	var pagamento entities.Pagamento



	err := c.ShouldBind(&pagamento)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	err = ph.PagamentoUseCase.EnviarPagamento(c, &pagamento)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Message: "Pagaamento enviado com sucesso",
	})
}

func (ph *PagamentoHandler) ConfirmarPagamento(c *gin.Context) {
	var pagamento entities.Pagamento

	err := c.ShouldBind(&pagamento)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	err = ph.PagamentoUseCase.ConfirmarPagamento(c, &pagamento)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

		c.JSON(http.StatusOK, response.SuccessResponse{
			//Message: "Pagamento confirmado com sucesso " + fmt.Sprintf("%+v", pagamento),
			Message: "Pagamento confirmado com sucesso " + fmt.Sprintf("%+v", c.Request.Body),
		})
		println("Teste")
		println(c.Request.Body)
}

