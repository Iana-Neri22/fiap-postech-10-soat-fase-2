package route

import (
	"lanchonete/api/handlers"
	"lanchonete/application/usecases"
	"lanchonete/bootstrap"
	"lanchonete/domain/services"
	"lanchonete/gateways"
	"lanchonete/infra/database/mongo"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func NewPagamentoRouter(env *bootstrap.Env, db mongo.Database, router *gin.RouterGroup) {

	redisAddress := os.Getenv("REDIS_ADDRESS")
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})


	pr := gateways.NewPagamentoRepository(db, service.CollectionPagamento)
	pc := &handler.PagamentoHandler{
		PagamentoUseCase: usecase.NewPagamentoUseCase(pr),
	}

	gateways.RegisterPaymentRoutes(redisClient)

	router.POST("/pagamento", pc.EnviarPagamento)
	router.POST("/pagamento/confirmar", pc.ConfirmarPagamento)

}