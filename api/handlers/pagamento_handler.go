package handler

import (
	"github.com/gin-gonic/gin"

)

func SucessoHandler(c *gin.Context) {
	c.JSON(200, gin.H{"response": "success"})
}

func PendenteHandler(c *gin.Context) {
	c.JSON(200, gin.H{"response": "success"})
}

func FalhaHandler(c *gin.Context) {
	c.JSON(200, gin.H{"response": "success"})
}




func WebhookHandlerGET(c *gin.Context) {
	if c.Request.Method != "GET" {
		c.JSON(405, gin.H{"error": "Method not allowed"})
		return
	}
	println("Teste 2")
	println(c.Request.Body)
	c.JSON(200, gin.H{"response": "webhook"})
}
