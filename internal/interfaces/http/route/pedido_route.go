package route

import (
	handler "lanchonete/internal/interfaces/http/handlers"
	"lanchonete/bootstrap"
	"lanchonete/internal/application/usecase"
	"lanchonete/internal/infrastructure/repository"
	"lanchonete/infra/database/mongo"

	"github.com/gin-gonic/gin"
)

// NovoPedidoRouter cria um novo roteador para pedidos
func NovoPedidoRouter(env *bootstrap.Env, db mongo.Database, router *gin.RouterGroup) {
	// Criar fábricas
	repositorioFactory := repository.NovoRepositorioFactory(db)
	casoDeUsoFactory := usecase.NovoCasoDeUsoFactory(repositorioFactory)
	
	// Criar casos de uso
	pedidoIncluirUseCase := casoDeUsoFactory.CriarPedidoIncluirUseCase()
	pedidoBuscarPorIdUseCase := casoDeUsoFactory.CriarPedidoBuscarPorIdUseCase()
	pedidoAtualizarStatusUseCase := casoDeUsoFactory.CriarPedidoAtualizarStatusUseCase()
	produtoBuscarPorIdUseCase := casoDeUsoFactory.CriarProdutoBuscarPorIdUseCase()
	pedidoListarTodosUseCase := casoDeUsoFactory.CriarPedidoListarTodosUseCase()
	
	// Criar handler
	puc := &handler.PedidoHandler{
		PedidoIncluirUseCase:         *pedidoIncluirUseCase,
		PedidoBuscarPorIdUseCase:     *pedidoBuscarPorIdUseCase,
		PedidoAtualizarStatusUseCase: *pedidoAtualizarStatusUseCase,
		ProdutoBuscarPorIdUseCase:    *produtoBuscarPorIdUseCase,
		PedidoListarTodosUseCase:     *pedidoListarTodosUseCase,
	}

	// Configurar rotas
	router.POST("/pedidos", puc.CriarPedido)
	router.GET("/pedidos/:nroPedido", puc.BuscarPedido)
	router.PUT("/pedidos/:nroPedido/status/:status", puc.AtualizarStatusPedido)
	router.POST("/pedidos/listartodos", puc.ListarTodosOsPedidos)
}
