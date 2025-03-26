package handler

import (
	"encoding/json"
	_ "lanchonete/docs"
	response "lanchonete/domain/responses"
	"lanchonete/domain/usecases"
	"lanchonete/presenters"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PedidoHandler struct {
	PedidoIncluirUseCase         usecases.PedidoIncluirUseCase
	PedidoBuscarPorIdUseCase     usecases.PedidoBuscarPorIdUseCase
	PedidoAtualizarStatusUseCase usecases.PedidoAtualizarStatusUseCase
	ProdutoBuscarPorNomeUseCase  usecases.ProdutoBuscaPorNomeUseCase
	PedidoListarTodosUseCase     usecases.PedidoListarTodosUseCase
	ClienteUseCase               usecases.ClienteUseCase
}

func NewPedidoHandler(pedidoIncluirUseCase usecases.PedidoIncluirUseCase,
	pedidoBuscarPorIdUseCase usecases.PedidoBuscarPorIdUseCase,
	pedidoAtualizarStatusUsecase usecases.PedidoAtualizarStatusUseCase,
	produtoBuscarPorNomeUseCase usecases.ProdutoBuscaPorNomeUseCase,
	pedidoListarTodosUseCase usecases.PedidoListarTodosUseCase,
	clienteUseCase usecases.ClienteUseCase) *PedidoHandler {
	return &PedidoHandler{
		PedidoIncluirUseCase:         pedidoIncluirUseCase,
		PedidoBuscarPorIdUseCase:     pedidoBuscarPorIdUseCase,
		PedidoAtualizarStatusUseCase: pedidoAtualizarStatusUsecase,
		ProdutoBuscarPorNomeUseCase:  produtoBuscarPorNomeUseCase,
		PedidoListarTodosUseCase:     pedidoListarTodosUseCase,
		ClienteUseCase:               clienteUseCase,
	}
}

// CriarPedido godoc
// @Summary Cria um pedido
// @Description Cria um pedido
// @Tags pedido
// @Router /pedidos [post]
// @Accept  json
// @Produce  json
// @Param pedido body presenters.PedidoDTO true "Pedido"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
func (h *PedidoHandler) CriarPedido(r *gin.Context) {
	var pedidoDTO presenters.PedidoDTO
	var produtoDTO presenters.ProdutoDTO
	var produtosDTO presenters.ProdutoDTOs
	err := json.NewDecoder(r.Request.Body).Decode(&pedidoDTO)
	if err != nil {
		r.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	for _, nomeProduto := range pedidoDTO.Produtos {
		produto, err := h.ProdutoBuscarPorNomeUseCase.Run(r, nomeProduto)
		if err != nil {
			r.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Produto não Cadastrado!"})
			return
		}
		produtoDTO.ToDTO(produto)
		produtosDTO = append(produtosDTO, produtoDTO)
	}
	produtos := produtosDTO.ConvertProdutoDTOsToEntities()
	cliente, err := h.ClienteUseCase.BuscarCliente(r, pedidoDTO.ClienteCPF)
	if err != nil {
		println("cliente não cadastrado: %w", err)
	}
	pedidoID, err := h.PedidoIncluirUseCase.Run(r, cliente, produtos, pedidoDTO.Personalizacao)
	if err != nil {
		r.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	r.JSON(http.StatusOK, response.SuccessResponse{
		Message: "Pedido criado com sucesso: " + pedidoID,
	})
}

// BuscarPedido godoc
// @Summary Busca um pedido
// @Description Busca um pedido
// @Tags pedido
// @Router /pedidos/{ID} [get]
// @Accept  json
// @Produce  json
// @Param ID path string true "Número do pedido"
// @Success 200 {object} entities.Pedido
// @Failure 400 {object} response.ErrorResponse
func (h *PedidoHandler) BuscarPedido(r *gin.Context) {
	nroPedido := r.Param("nroPedido")
	pedido, err := h.PedidoBuscarPorIdUseCase.Run(r, nroPedido)
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

// ProdutoListarTodos godoc
// @Summary Lista todos os pedidos no banco
// @Description Lista todos os pedidos presentes no banco
// @Tags pedido
// @Router /pedidos/listartodos [POST]
// @Accept  json
// @Produce  json
// @Success 200 {object} []entities.Pedido
// @Failure 400 {object} response.ErrorResponse
func (h *PedidoHandler) ListarTodosOsPedidos(r *gin.Context) {
	pedidos, err := h.PedidoListarTodosUseCase.Run(r)
	if err != nil {
		r.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	r.JSON(http.StatusOK, pedidos)
}
