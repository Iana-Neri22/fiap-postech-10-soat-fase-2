package handler

import (
	"encoding/json"
	_ "lanchonete/docs"
	"lanchonete/domain/entities"
	response "lanchonete/domain/responses"
// <<<<<<< HEAD
	// service "lanchonete/domain/services"
// =======
// >>>>>>> Bruno
	"lanchonete/domain/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PedidoHandler struct {
	PedidoIncluirUseCase         usecases.PedidoIncluirUseCase
	PedidoBuscarPorIdUseCase     usecases.PedidoBuscarPorIdUseCase
	PedidoAtualizarStatusUseCase usecases.PedidoAtualizarStatusUseCase
}

func NewPedidoHandler(pedidoIncluirUseCase usecases.PedidoIncluirUseCase, pedidoBuscarPorIdUseCase usecases.PedidoBuscarPorIdUseCase, pedidoAtualizarStatusUsecase usecases.PedidoAtualizarStatusUseCase) *PedidoHandler {
	return &PedidoHandler{
		PedidoIncluirUseCase:         pedidoIncluirUseCase,
		PedidoBuscarPorIdUseCase:     pedidoBuscarPorIdUseCase,
		PedidoAtualizarStatusUseCase: pedidoAtualizarStatusUsecase,
	}
}

// func NewPedidoHandler(pedidoUseCase service.PedidoUseCase) *PedidoHandler {
// 	return &PedidoHandler{
// 		PedidoUseCase:        pedidoUseCase,
// 		PedidoIncluirUseCase: pedidoIncluirUseCase,
// 	}
// }

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

	_, err = h.PedidoIncluirUseCase.Run(r, pedido.Cliente, pedido.Produtos, pedido.Personalizacao)
	if err != nil {
		r.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	// err = h.PedidoUseCase.CriarPedido(r, ped)
	// if err != nil {
	// 	r.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
	// 	return
	// }

	r.JSON(http.StatusOK, response.SuccessResponse{
		Message: "Pedido criado com sucesso",
	})
}

// BuscarPedido godoc
// @Summary Busca um pedido
// @Description Busca um pedido
// @Tags pedido
// <<<<<<< HEAD
// @Router /pedidos/{ID} [get]
// =======
// >>>>>>> Bruno
// @Accept  json
// @Produce  json
// @Param ID path string true "Número do pedido"
// @Success 200 {object} entities.Pedido
// @Failure 400 {object} response.ErrorResponse
func (h *PedidoHandler) BuscarPedido(r *gin.Context) {
// <<<<<<< HEAD
	nroPedido := r.Param("nroPedido")
	// pedido, err := h.PedidoBuscarPorIdUseCase.Run(r, ID)
// =======
	// nroPedido := r.Param("nroPedido")
	pedido, err := h.PedidoBuscarPorIdUseCase.Run(r, nroPedido)
// >>>>>>> Bruno
	if err != nil {
		r.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	r.JSON(http.StatusOK, pedido)

}

// BuscarPedido godoc
// @Summary Atualiza um pedido a partir de sua Identificação
// @Description Atualizar um pedido
// @Tags pedido
// @Router /pedidos/{nroPedido}/status/{status} [put]
// @Accept  json
// @Produce  json
// @Param nroPedido path string true "Número do pedido"
// @Param status path string true "Novo Status do pedido"
// @Success 200 {object} entities.Pedido
// @Failure 400 {object} response.ErrorResponse
func (h *PedidoHandler) AtualizarStatusPedido(r *gin.Context) {
	nroPedido := r.Param("nroPedido")
	status := r.Param("status")
	err := h.PedidoAtualizarStatusUseCase.Run(r, nroPedido, status)
	if err != nil {
		r.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	r.JSON(http.StatusOK, response.SuccessResponse{
		Message: "Status do pedido atualizado com sucesso",
	})
}
