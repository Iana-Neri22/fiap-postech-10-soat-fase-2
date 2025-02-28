package route

import (
	"lanchonete/bootstrap"
	"lanchonete/infra/database/mongo"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, db mongo.Database, gin gin.IRouter) {
	router := gin.Group("")
	
	NewClienteRouter(env, db, router)
	NewPagamentoRouter(env, db, router)
}