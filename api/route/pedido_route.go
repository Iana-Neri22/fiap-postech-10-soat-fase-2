package route

import (
	"lanchonete/api/handlers"
	"lanchonete/bootstrap"
	"lanchonete/gateways"
	"lanchonete/infra/database/mongo"
	"lanchonete/application/usecases"

	"github.com/gin-gonic/gin"
)

func NewPedidoRouter(env *bootstrap.Env, db mongo.Database, router *gin.RouterGroup) {
	collection := "pedido"
	pr := gateways.NewPedidoGateway(db, collection)
	puc := &handler.PedidoHandler{
		PedidoUseCase: usecase.NewPedidoUseCase(pr),
	}

	router.POST("/pedidos", puc.CriarPedido)
	router.GET("/pedidos/:ID", puc.BuscarPedido)
}