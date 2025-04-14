package route

import (
	handler "lanchonete/internal/interfaces/http/handlers"
	"lanchonete/bootstrap"
	"lanchonete/usecases"
	"lanchonete/gateways"
	"lanchonete/infra/database/mongo"

	"github.com/gin-gonic/gin"
)

func NewPedidoRouter(env *bootstrap.Env, db mongo.Database, router *gin.RouterGroup) {
	pr := gateways.NewPedidoGateway(db, mongo.CollectionPedido)
	prod := gateways.NewProdutoGateway(db, mongo.CollectionProduto)
	puc := &handler.PedidoHandler{
		PedidoIncluirUseCase:         *usecases.NewPedidoIncluirUseCase(pr),
		PedidoBuscarPorIdUseCase:     *usecases.NewPedidoBuscarPorIdUseCase(pr),
		PedidoAtualizarStatusUseCase: *usecases.NewPedidoAtualizarStatusUseCase(pr),
		ProdutoBuscarPorIdUseCase:    *usecases.NewProdutoBuscaPorIdUseCase(prod),
		PedidoListarTodosUseCase:     *usecases.NewPedidoListarTodosUseCase(pr),
	}

	router.POST("/pedidos", puc.CriarPedido)
	router.GET("/pedidos/:nroPedido", puc.BuscarPedido)
	router.PUT("/pedidos/:nroPedido/status/:status", puc.AtualizarStatusPedido)
	router.POST("/pedidos/listartodos", puc.ListarTodosOsPedidos)
}
