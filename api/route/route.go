package route

import (
	"lanchonete/bootstrap"
	"lanchonete/mongo"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, db mongo.Database, gin *gin.Engine) {
	router := gin.Group("")

	NewClienteRouter(env, db, router)
	NewPagamentoRouter(env, db, router)
}