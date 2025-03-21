package route

import (
	handler "lanchonete/api/handlers"
	usecase "lanchonete/application/usecases"
	"lanchonete/bootstrap"
	"lanchonete/gateways"
	"lanchonete/infra/database/mongo"

	"fmt"

	"github.com/gin-gonic/gin"
)

func NewAcompanhamentoRouter(env *bootstrap.Env, db mongo.Database, router *gin.RouterGroup) {

	ar := gateways.NewAcompanhamentoRepository(db, CollectionAcompanhamento)
	auc := &handler.AcompanhamentoHandler{
		AcompanhamentoUseCase: usecase.NewAcompanhamentoUseCase(ar),
	}

	// Debug print to confirm route registration
	fmt.Printf("Registrando rotas do acompanhamento\n")

	router.POST("/acompanhamento", auc.CriarAcompanhamento)
	router.GET("/acompanhamento/show", auc.BuscarAcompanhamento)
	router.GET("/acompanhamento/:ID", auc.BuscarPedido)
	router.POST("/acompanhamento/:IDAcompanhamento/:IDPedido", auc.AdicionarPedido)
	router.PUT("/acompanhamento/:ID/status", auc.AtualizarStatusPedido)
}
