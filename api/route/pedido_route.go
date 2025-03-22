package route

import (
	handler "lanchonete/api/handlers"
	//	usecase "lanchonete/application/usecases"
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
		//PedidoUseCase: usecase.NewPedidoUseCase(pr),
		PedidoIncluirUseCase:         *usecases.NewPedidoIncluirUseCase(pr),
		PedidoBuscarPorIdUseCase:     *usecases.NewPedidoBuscarPorIdUseCase(pr),
		PedidoAtualizarStatusUseCase: *usecases.NewPedidoAtualizarStatusUseCase(pr),
	}

	router.POST("/pedidos", puc.CriarPedido)
// <<<<<<< HEAD
// 	router.GET("/pedidos/:ID", puc.BuscarPedido)
// }
// =======
	router.GET("/pedidos/:nroPedido", puc.BuscarPedido)
	router.PUT("/pedidos/:nroPedido/status/:status", puc.AtualizarStatusPedido)
}
// >>>>>>> Bruno
