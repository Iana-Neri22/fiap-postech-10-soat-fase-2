package handler

import (
	"encoding/json"
	_ "lanchonete/docs"
	"lanchonete/domain/entities"
	"lanchonete/domain/responses"
	"lanchonete/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PedidoHandler struct {
	PedidoUseCase service.PedidoUseCase
}

func NewPedidoHandler(pedidoUseCase service.PedidoUseCase) *PedidoHandler {
	return &PedidoHandler{
		PedidoUseCase: pedidoUseCase,
	}
}

// CriarPedido godoc
// @Summary Cria um pedido
// @Description Cria um pedido
// @Tags pedido
// @Router /pedidos [post]
// @Accept  json
// @Produce  json
// @Param pedido body entities.Pedido true "Pedido"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
func (h *PedidoHandler) CriarPedido(r *gin.Context) {
	var pedido entities.Pedido
	err := json.NewDecoder(r.Request.Body).Decode(&pedido)
	if err != nil {
		r.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	err = h.PedidoUseCase.CriarPedido(r, &pedido)
	if err != nil {
		r.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	r.JSON(http.StatusOK, response.SuccessResponse{
		Message: "Pedido criado com sucesso",
	})
}

// BuscarPedido godoc
// @Summary Busca um pedido
// @Description Busca um pedido
// @Tags pedido
// @Router /pedidos/{nroPedido} [get]
// @Accept  json
// @Produce  json
// @Param nroPedido path string true "Número do pedido"
// @Success 200 {object} entities.Pedido
// @Failure 400 {object} response.ErrorResponse
func (h *PedidoHandler) BuscarPedido(r *gin.Context) {
	nroPedido := r.Param("nroPedido")
	pedido, err := h.PedidoUseCase.BuscarPedido(r, nroPedido)
	if err != nil {
		r.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	r.JSON(http.StatusOK, pedido)

}
