package handler

import (
	_ "lanchonete/docs"
	"lanchonete/domain/entities"
	response "lanchonete/domain/responses"
	service "lanchonete/domain/services"
	"lanchonete/presenters"
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

// CriarAcompanhamento godoc
// @Summary Cria um acompanhamento
// @Description Cria um acompanhamento
// @Tags acompanhamento
// @Router /acompanhamento [post]
// @Accept  json
// @Produce  json
// @Param acompanhamento body entities.AcompanhamentoPedido true "Acompanhamento"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
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
		Message: "Acompanhamento criado com sucesso",
	})
}

// BuscarPedido godoc
// @Summary Busca um pedido
// @Description Busca um pedido
// @Tags acompanhamento
// @Router /acompanhamento/{nroPedido} [get]
// @Accept  json
// @Produce  json
// @Param nroPedido path string true "ID do pedido"
// @Success 200 {object} entities.AcompanhamentoPedido
// @Failure 400 {object} response.ErrorResponse
func (ah *AcompanhamentoHandler) BuscarPedido(p *gin.Context) {
	nroPedido := p.Param("ID")
	pedido, err := ah.AcompanhamentoUseCase.BuscarPedido(p, nroPedido)
	if err != nil {
		p.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	p.JSON(http.StatusOK, pedido)
}

// AdicionarPedido godoc
// @Summary Adiciona um pedido
// @Description Adiciona um pedido
// @Tags acompanhamento
// @Router /acompanhamento/{nroPedido} [post]
// @Accept  json
// @Produce  json
// @Param ID path string true "ID do pedido"
// @Param acompanhamento body entities.AcompanhamentoPedido true "Acompanhamento"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
func (ah *AcompanhamentoHandler) AdicionarPedido(p *gin.Context) {
	nroPedido := p.Param("ID")
	pedido, err := ah.AcompanhamentoUseCase.BuscarPedido(p, nroPedido)
	if err != nil {
		p.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	var acompanhamento entities.AcompanhamentoPedido
	err = p.ShouldBind(&acompanhamento)
	if err != nil {
		p.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	err = ah.AcompanhamentoUseCase.AdicionarPedido(p, &acompanhamento, &pedido)
	if err != nil {
		p.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	p.JSON(http.StatusOK, response.SuccessResponse{
		Message: "Pedido adicionado com sucesso",
	})
}

// BuscarAcompanhamento godoc
// @Summary Busca um acompanhamento
// @Description Busca um acompanhamento
// @Tags acompanhamento
// @Router /acompanhamento/show [get]
// @Accept  json
// @Produce  json
// @Param nroPedido path string true "Número do pedido"
// @Success 200 {object} presenters.AcompanhamentoDTO
// @Failure 400 {object} response.ErrorResponse
func (ah *AcompanhamentoHandler) BuscarAcompanhamento(c *gin.Context) {
	nroPedido := c.Param("nroPedido")
	acompanhamento, err := ah.AcompanhamentoUseCase.BuscarAcompanhamento(c, nroPedido)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	acompanhamentoDTO := presenters.NewAcompanhamentoDTO(&acompanhamento)
	c.JSON(http.StatusOK, acompanhamentoDTO)
}

// AtualizarStatusPedido godoc
// @Summary Atualiza o status de um pedido
// @Description Atualiza o status de um pedido
// @Tags acompanhamento
// @Router /acompanhamento/{nroPedido}/status [put]
// @Accept  json
// @Produce  json
// @Param nroPedido path string true "ID do pedido"
// @Param status body string true "Novo status"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
func (ah *AcompanhamentoHandler) AtualizarStatusPedido(a *gin.Context) {
	nroPedido := a.Param("nroPedido")

	var requestBody struct {
		Status string `json:"status"`
	}

	if err := a.ShouldBindJSON(&requestBody); err != nil {
		a.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	// Get the acompanhamento
	acompanhamento, err := ah.AcompanhamentoUseCase.BuscarAcompanhamento(a, "")
	if err != nil {
		a.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	// Update the status
	err = ah.AcompanhamentoUseCase.AtualizarStatusPedido(a, &acompanhamento, nroPedido, entities.StatusPedido(requestBody.Status))
	if err != nil {
		a.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	a.JSON(http.StatusOK, response.SuccessResponse{
		Message: "Status atualizado com sucesso",
	})
}
