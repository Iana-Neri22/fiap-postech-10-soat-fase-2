package route

import (
	handler "lanchonete/api/handlers"
	"lanchonete/bootstrap"
	"lanchonete/domain/usecases"
	"lanchonete/gateways"
	"lanchonete/infra/database/mongo"

	"github.com/gin-gonic/gin"
)

func NewPedidoRouter(env *bootstrap.Env, db mongo.Database, router *gin.RouterGroup) {
	collection := "pedido"
	pr := gateways.NewPedidoGateway(db, collection)
	puc := &handler.PedidoHandler{
		PedidoIncluirUseCase:         *usecases.NewPedidoIncluirUseCase(pr),
		PedidoBuscarPorIdUseCase:     *usecases.NewPedidoBuscarPorIdUseCase(pr),
		PedidoAtualizarStatusUseCase: *usecases.NewPedidoAtualizarStatusUseCase(pr),
	}

	router.POST("/pedidos", puc.CriarPedido)
	router.GET("/pedidos/:nroPedido", puc.BuscarPedido)
	router.PUT("/pedidos/:nroPedido/status/:status", puc.AtualizarStatusPedido)
}
