package route

import (
	handler "lanchonete/internal/interfaces/http/handlers"
	"lanchonete/bootstrap"
	"lanchonete/infra/database/mongo"
	"lanchonete/gateways"
	"lanchonete/internal/application/usecases"

	"github.com/gin-gonic/gin"
)

// NewClienteRouter creates and configures all cliente-related routes
func NewClienteRouter(env *bootstrap.Env, db mongo.Database, router *gin.RouterGroup) {
	cr := gateways.NewClienteRepository(db, mongo.CollectionCliente)
	cc := &handler.ClienteHandler{
		ClienteUseCase: usecases.NewClienteUseCase(cr),
	}

	router.GET("/cliente/:CPF", cc.BuscarCliente)
	router.POST("/cliente", cc.CriarCliente)
}