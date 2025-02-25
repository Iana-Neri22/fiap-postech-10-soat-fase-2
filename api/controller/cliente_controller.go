package controller

import (
	"net/http"

	"lanchonete/domain/entities"
	"lanchonete/domain/usecases"
	"lanchonete/domain/responses"
	"github.com/gin-gonic/gin"
)

type ClienteController struct {
	ClienteUseCase usecases.ClienteUseCase
}

func (cc *ClienteController) CriarCliente(c *gin.Context) {
	var cliente entities.Cliente

	err := c.ShouldBind(&cliente)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: err.Error()})
		return
	}

	err = cc.ClienteUseCase.CriarCliente(c, &cliente)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Cliente criado com sucesso",
	})
}

func (cc *ClienteController) BuscarCliente(c *gin.Context) {
	CPF := c.Param("CPF")
	cliente, err := cc.ClienteUseCase.BuscarCliente(c, CPF)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, cliente)
}