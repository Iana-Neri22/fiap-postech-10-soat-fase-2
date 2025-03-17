package handler

import (
	"encoding/json"
	"fmt"
	"io"
	_ "lanchonete/docs"
	"lanchonete/domain/entities"
	"lanchonete/domain/responses"
	"lanchonete/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PagamentoHandler struct {
	PagamentoUseCase service.PagamentoUseCase
}

func (ph *PagamentoHandler) EnviarPagamento(c *gin.Context) {
	var pagamento entities.Pagamento



	err := c.ShouldBind(&pagamento)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	err = ph.PagamentoUseCase.EnviarPagamento(c, &pagamento)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Message: "Pagaamento enviado com sucesso",
	})
}

func (ph *PagamentoHandler) ConfirmarPagamento(c *gin.Context) {
	// Read the request body
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Failed to read request body: " + err.Error()})
		return
	}
	defer c.Request.Body.Close()

	// Print the raw request body
	fmt.Println("Received webhook data:", string(body))

	// Bind the request body to the pagamento struct
	var pagamento entities.Pagamento
	if err := json.Unmarshal(body, &pagamento); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Failed to unmarshal request body: " + err.Error()})
		return
	}

	// Print the parsed pagamento struct
	fmt.Printf("Parsed pagamento: %+v\n", pagamento)

	// Call the use case to confirm the payment
	if err := ph.PagamentoUseCase.ConfirmarPagamento(c, &pagamento); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, response.SuccessResponse{
		Message: "Pagamento confirmado com sucesso",
	})
}
