package route

import (
	handler "lanchonete/internal/interfaces/http/handlers"
	"lanchonete/bootstrap"
	"lanchonete/gateways"
	"lanchonete/infra/database/mongo"
	appusecases "lanchonete/internal/application/usecases"
	"lanchonete/usecases"

	"fmt"

	"github.com/gin-gonic/gin"
)

// NewAcompanhamentoRouter creates and configures all acompanhamento-related routes
func NewAcompanhamentoRouter(env *bootstrap.Env, db mongo.Database, router *gin.RouterGroup) {
	ar := gateways.NewAcompanhamentoRepository(db, mongo.CollectionAcompanhamento)
	pr := gateways.NewPedidoGateway(db, mongo.CollectionPedido)
	auc := &handler.AcompanhamentoHandler{
		AcompanhamentoUseCase:        appusecases.NewAcompanhamentoUseCase(ar),
		PedidoAtualizarStatusUseCase: *usecases.NewPedidoAtualizarStatusUseCase(pr),
	}

	fmt.Printf("Registrando rotas do acompanhamento\n")

	router.POST("/acompanhamento", auc.CriarAcompanhamento)
	router.GET("/acompanhamento/show", auc.BuscarAcompanhamento)
	router.GET("/acompanhamento/:ID", auc.BuscarPedido)
	router.POST("/acompanhamento/:IDAcompanhamento/:IDPedido", auc.AdicionarPedido)
	router.PUT("acompanhamento/:IDAcompanhamento/:IDPedido/:status", auc.AtualizarStatusPedido)
}
