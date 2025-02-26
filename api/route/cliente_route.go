package route

import (
	"lanchonete/bootstrap"
	"lanchonete/api/handler"
	"lanchonete/adapters/usecases"
	"lanchonete/adapters/mongo"
	"lanchonete/gateways"
	"github.com/gin-gonic/gin"
)

func NewClienteRouter(env *bootstrap.Env, db *mongo.Database, group *gin.RouterGroup) {
	cr := gateways.NewClienteGateway(db)
	cc := &handler.ClienteHandler{
		ClienteUseCase: usecases.NewClienteUseCase(cr),
	}

	group.GET("/cliente/:CPF", cc.BuscarCliente)
	group.POST("/cliente", cc.CriarCliente)
}