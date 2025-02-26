package handler

import (
	"net/http"

	"lanchonete/domain/entities"
	"lanchonete/domain/usecases"
	"lanchonete/domain/responses"
	"github.com/gin-gonic/gin"
)

type ClienteHandler struct {
	ClienteUseCase usecases.ClienteUseCase
}

func (cc *ClienteHandler) CriarCliente(c *gin.Context) {
	var cliente entities.Cliente

	err := c.ShouldBind(&cliente)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: err.Error()})
		return
	}

	err = cc.ClienteUseCase.CriarCliente(c, &cliente)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Cliente criado com sucesso",
	})
}

func (cc *ClienteHandler) BuscarCliente(c *gin.Context) {
	CPF := c.Param("CPF")
	cliente, err := cc.ClienteUseCase.BuscarCliente(c, CPF)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, cliente)
}
