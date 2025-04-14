package route

import (
	handler "lanchonete/internal/interfaces/http/handlers"
	"lanchonete/internal/application/usecases"
	"lanchonete/bootstrap"
	"lanchonete/gateways"
	"lanchonete/infra/database/mongo"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// NewPagamentoRouter creates and configures all pagamento-related routes
func NewPagamentoRouter(env *bootstrap.Env, db mongo.Database, router *gin.RouterGroup) {
	redisAddress := os.Getenv("REDIS_ADDRESS")
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pr := gateways.NewPagamentoRepository(db, mongo.CollectionPagamento)
	pc := &handler.PagamentoHandler{
		EnviarPagamentoUseCase: usecases.NewEnviarPagamentoUseCase(pr),
		ConfirmarPagamentoUseCase: usecases.NewConfirmarPagamentoUseCase(pr),
	}

	gateways.RegisterPaymentRoutes(redisClient)

	router.POST("/pagamento", pc.EnviarPagamento)
	router.POST("/pagamento/confirmar", pc.ConfirmarPagamento)
}