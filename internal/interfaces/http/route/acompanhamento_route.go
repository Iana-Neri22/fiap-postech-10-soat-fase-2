package route

import (
	handler "lanchonete/internal/interfaces/http/handlers"
	"lanchonete/bootstrap"
	"lanchonete/internal/application/factories"
	"lanchonete/internal/infrastructure/repository"
	"lanchonete/infra/database/mongo"
	appusecases "lanchonete/internal/application/usecases"

	"fmt"

	"github.com/gin-gonic/gin"
)

// NewAcompanhamentoRouter creates and configures all acompanhamento-related routes
func NewAcompanhamentoRouter(env *bootstrap.Env, db mongo.Database, router *gin.RouterGroup) {
	// Criar fábricas
	repositorioFactory := repository.NovoRepositorioFactory(db)
	casoDeUsoFactory := factory.NovoCasoDeUsoFactory(repositorioFactory)
	
	// Criar casos de uso
	acompanhamentoUseCase := appusecases.NewAcompanhamentoUseCase(repositorioFactory.CriarAcompanhamentoRepository())
	pedidoAtualizarStatusUseCase := casoDeUsoFactory.CriarPedidoAtualizarStatusUseCase()
	
	auc := &handler.AcompanhamentoHandler{
		AcompanhamentoUseCase:        acompanhamentoUseCase,
		PedidoAtualizarStatusUseCase: *pedidoAtualizarStatusUseCase,
	}

	fmt.Printf("Registrando rotas do acompanhamento\n")

	router.POST("/acompanhamento", auc.CriarAcompanhamento)
	router.GET("/acompanhamento/show", auc.BuscarAcompanhamento)
	router.GET("/acompanhamento/:ID", auc.BuscarPedido)
	router.POST("/acompanhamento/:IDAcompanhamento/:IDPedido", auc.AdicionarPedido)
	router.PUT("acompanhamento/:IDAcompanhamento/:IDPedido/:status", auc.AtualizarStatusPedido)
}
