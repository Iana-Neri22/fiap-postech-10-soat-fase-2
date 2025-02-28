package handler

import (
	"net/http"

	"lanchonete/domain/services"
	"lanchonete/domain/entities"
	"lanchonete/domain/responses"
	
	"github.com/gin-gonic/gin"
)

type ClienteHandler struct {
	ClienteUseCase service.ClienteUseCase
}

func (cc *ClienteHandler) CriarCliente(c *gin.Context) {
	var cliente entities.Cliente

	err := c.ShouldBind(&cliente)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	err = cc.ClienteUseCase.CriarCliente(c, &cliente)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Message: "Cliente criado com sucesso",
	})
}

func (cc *ClienteHandler) BuscarCliente(c *gin.Context) {
	CPF := c.Param("CPF")
	cliente, err := cc.ClienteUseCase.BuscarCliente(c, CPF)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, cliente)
}