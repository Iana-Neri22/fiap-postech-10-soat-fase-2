package route

import (
	"lanchonete/adapters/mercado_pago"
	handler "lanchonete/api/handlers"
	"lanchonete/bootstrap"
	"lanchonete/infra/database/mongo"

	//"lanchonete/api/handlers"
	"github.com/gin-gonic/gin"
)

func NewPagamentoRouter(env *bootstrap.Env, db mongo.Database, group *gin.RouterGroup) {
	
	group.POST("/pagamento", func(c *gin.Context) {
		access_token := env.AccessToken
		mercadopago.MercadoPagoHandler(c, access_token)
	})
	group.POST("/pagamento/atualizado", mercadopago.WebhookHandler) 
	group.GET("/pagamento/atualizado", handler.WebhookHandlerGET)
	group.GET("/pagamento/sucesso", handler.SucessoHandler)
	group.GET("/pagamento/pendente", handler.PendenteHandler)
	group.GET("/pagamento/falha", handler.FalhaHandler)
}