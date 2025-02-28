package route

import (
	"lanchonete/api/handlers"
	"lanchonete/bootstrap"
	"lanchonete/domain/services"
	"lanchonete/infra/database/mongo"
	"lanchonete/gateways"
	"lanchonete/application/usecases"

	"github.com/gin-gonic/gin"
)

func NewClienteRouter(env *bootstrap.Env, db mongo.Database, group *gin.RouterGroup) {
	cr := gateways.NewClienteRepository(db, service.CollectionCliente)
	cc := &handler.ClienteHandler{
		ClienteUseCase: usecase.NewClienteUseCase(cr),
	}

	group.GET("/cliente/:CPF", cc.BuscarCliente)
	group.POST("/cliente", cc.CriarCliente)
}