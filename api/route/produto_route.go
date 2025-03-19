package route

import (
	handler "lanchonete/api/handlers"
	usecase "lanchonete/application/usecases"
	"lanchonete/bootstrap"
	"lanchonete/gateways"
	"lanchonete/infra/database/mongo"

	"github.com/gin-gonic/gin"
)

func NewProdutoRouter(env *bootstrap.Env, db mongo.Database, router *gin.RouterGroup) {

	pr := gateways.NewProdutoGateway(db, CollectionProduto)
	pc := &handler.ProdutoHandler{
		ProdutoIncluirUseCase:     usecase.NewProdutoIncluirUseCase(pr),
		ProdutoBuscarPorIdUseCase: usecase.NewProdutoBuscarPorIdUseCase(pr),
	}

	router.POST("/produto", pc.ProdutoIncluir)
	router.GET("/produto/:id", pc.ProdutoBuscarPorId)

}