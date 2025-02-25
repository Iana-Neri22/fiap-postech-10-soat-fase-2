package route

import (
	"lanchonete/bootstrap"
	"lanchonete/mongo"
	"lanchonete/mercado-pago"
	"lanchonete/pagamento"
	"github.com/gin-gonic/gin"
)

func NewPagamentoRouter(env *bootstrap.Env, db mongo.Database, group *gin.RouterGroup) {
	
	group.POST("/pagamento", func(c *gin.Context) {
		access_token := env.AccessToken
		mercadopago.MercadoPagoHandler(c, access_token)
	})
	group.POST("/pagamento/atualizado", mercadopago.WebhookHandler) 
	group.GET("/pagamento/atualizado", pagamentohandler.WebhookHandlerGET)
	group.GET("/pagamento/sucesso", pagamentohandler.SucessoHandler)
	group.GET("/pagamento/pendente", pagamentohandler.PendenteHandler)
	group.GET("/pagamento/falha", pagamentohandler.FalhaHandler)
}