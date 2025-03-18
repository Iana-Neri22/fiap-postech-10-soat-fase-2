package route

import (
	"lanchonete/api/handlers"
	"lanchonete/application/usecases"
	"lanchonete/bootstrap"
	"lanchonete/gateways"
	"lanchonete/infra/database/mongo"
	"lanchonete/domain/services"

	"github.com/gin-gonic/gin"
)

func NewPedidoRouter(env *bootstrap.Env, db mongo.Database, router *gin.RouterGroup) {
	pr := gateways.NewPedidoRepository(db, service.CollectionPedido)
	puc := &handler.PedidoHandler{
		PedidoUseCase: usecase.NewPedidoUseCase(pr),
	}

	router.POST("/pedidos", puc.CriarPedido)
	router.GET("/pedidos/:nroPedido", puc.BuscarPedido)
}