package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


type PagamentoHandler struct{}

func (p *PagamentoHandler) SucessoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"response": "success"})
}

func (p *PagamentoHandler) PendenteHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"response": "pending"})
}

func (p *PagamentoHandler) FalhaHandler(c *gin.Context) {
	c.JSON(200, gin.H{"response": "falha"})
}

func (p* PagamentoHandler) WebhookHandler(c *gin.Context) {
	if c.Request.Method != "GET" {
		c.JSON(405, gin.H{"error": "Method not allowed"})
		return
	}
	println("Teste 2")
	println(c.Request.Body)
	c.JSON(200, gin.H{"response": "webhook"})
}

