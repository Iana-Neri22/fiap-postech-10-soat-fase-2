package route

import (
	"lanchonete/api/handlers"
	"lanchonete/application/usecases"
	"lanchonete/bootstrap"
	"lanchonete/gateways"
	"lanchonete/infra/database/mongo"

	"github.com/gin-gonic/gin"
)

func NewAcompanhamentoRouter(env *bootstrap.Env, db mongo.Database, router *gin.RouterGroup) {
	
	ar := gateways.NewAcompanhamentoRepository(db, CollectionPedido)
	auc := &handler.AcompanhamentoHandler{
		AcompanhamentoUseCase: usecase.NewAcompanhamentoUseCase(ar),
	}

	router.POST("/acompanhamento", auc.CriarAcompanhamento)
	router.GET("/acompanhamento/:nroPedido", auc.BuscarPedido)
}
