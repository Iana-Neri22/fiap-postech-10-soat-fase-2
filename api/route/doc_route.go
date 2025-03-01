package route

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

func NewDocRouter(router *gin.RouterGroup) {
	
	// go install github.com/swaggo/swag/cmd/swag@latest
	// swag init 
	// Swagger http://localhost:8080/docs/index.html
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}