package handler

import (
	"lanchonete/domain/entities"
	"lanchonete/domain/responses"
	"lanchonete/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AcompanhamentoHandler struct {
	AcompanhamentoUseCase service.AcompanhamentoUseCase
}

func NewAcompanhamentoHandler(auc service.AcompanhamentoUseCase) *AcompanhamentoHandler {
	return &AcompanhamentoHandler{
		AcompanhamentoUseCase: auc,
	}
}

func (ah *AcompanhamentoHandler) CriarAcompanhamento(a *gin.Context) {
	var acompanhamento entities.AcompanhamentoPedido

	err := a.ShouldBind(&acompanhamento)
	if err != nil {
		a.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	err = ah.AcompanhamentoUseCase.CriarAcompanhamento(a, &acompanhamento)
	if err != nil {
		a.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	a.JSON(http.StatusOK, response.SuccessResponse{
		Message: "Cliente criado com sucesso",
	})
}

func (ah *AcompanhamentoHandler) BuscarPedido(p *gin.Context) {
	ID := p.Param("ID")
	pedido, err := ah.AcompanhamentoUseCase.BuscarPedido(p, ID)
	if err != nil {
		p.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	p.JSON(http.StatusOK, pedido)
}