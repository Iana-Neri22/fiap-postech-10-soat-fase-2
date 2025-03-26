package route

import (
	handler "lanchonete/api/handlers"
	usecase "lanchonete/application/usecases"
	"lanchonete/bootstrap"
	"lanchonete/domain/usecases"
	"lanchonete/gateways"
	"lanchonete/infra/database/mongo"

	"github.com/gin-gonic/gin"
)

func NewPedidoRouter(env *bootstrap.Env, db mongo.Database, router *gin.RouterGroup) {
	collection := "pedido"
	pr := gateways.NewPedidoGateway(db, collection)
	cr := gateways.NewClienteRepository(db, collectionCliente)
	prod := gateways.NewProdutoGateway(db, "produto")
	puc := &handler.PedidoHandler{
		PedidoIncluirUseCase:         *usecases.NewPedidoIncluirUseCase(pr),
		PedidoBuscarPorIdUseCase:     *usecases.NewPedidoBuscarPorIdUseCase(pr),
		PedidoAtualizarStatusUseCase: *usecases.NewPedidoAtualizarStatusUseCase(pr),
		PedidoListarTodosUseCase:     *usecases.NewPedidoListarTodosUseCase(pr),
		ProdutoBuscarPorNomeUseCase:  *usecases.NewProdutoBuscaPorNomeUseCase(prod),
		ClienteUseCase:               usecase.NewClienteUseCase(cr),
	}

	router.POST("/pedidos", puc.CriarPedido)
	router.GET("/pedidos/:nroPedido", puc.BuscarPedido)
	router.PUT("/pedidos/:nroPedido/status/:status", puc.AtualizarStatusPedido)
	router.POST("/pedidos/listartodos", puc.ListarTodosOsPedidos)
}
