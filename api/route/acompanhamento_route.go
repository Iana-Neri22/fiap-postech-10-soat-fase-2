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

func NewAcompanhamentoRouter(env *bootstrap.Env, db mongo.Database, router *gin.RouterGroup) {
	ar := gateways.NewAcompanhamentoRepository(db, service.CollectionPedido)
	auc := &handler.AcompanhamentoHandler{
		AcompanhamentoUseCase: usecase.NewAcompanhamentoUseCase(ar),
	}

	router.POST("/acompanhamento", auc.CriarAcompanhamento)
	router.GET("/acompanhamento/:nroPedido", auc.BuscarPedido)
	router.POST("/acompanhamento/:nroPedido", auc.AdicionarPedido)
}
