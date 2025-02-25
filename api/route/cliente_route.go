package route

import (
	"lanchonete/bootstrap"
	"lanchonete/api/controller"
	"lanchonete/usecases"
	"lanchonete/mongo"
	"lanchonete/gateways"
	"github.com/gin-gonic/gin"
)

func NewClienteRouter(env *bootstrap.Env, db mongo.Database, group *gin.RouterGroup) {
	cr := gateways.NewClienteGateway(db)
	cc := &controller.ClienteController{
		ClienteUseCase: usecases.NewClienteUseCase(cr),
	}

	group.GET("/cliente/:CPF", cc.BuscarCliente)
	group.POST("/cliente", cc.CriarCliente)
}