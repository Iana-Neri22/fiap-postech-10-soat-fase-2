package mercadopago

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"lanchonete/bootstrap"
	"log"
	"net/http"

	//"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mercadopago/sdk-go/pkg/config"

	//"github.com/mercadopago/sdk-go/pkg/merchantorder"

	//"github.com/mercadopago/sdk-go/pkg/payment"
	//"github.com/mercadopago/sdk-go/pkg/merchantorder"
	"github.com/mercadopago/sdk-go/pkg/preference"
)

func WebhookHandler(c *gin.Context) {
	access_token := bootstrap.NewEnv().AccessToken
	paymentID := c.Query("id")

	if paymentID == "" {
		println("Payment ID is empty")
	}

	url := fmt.Sprintf("https://api.mercadopago.com/v1/payments/%s", paymentID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", access_token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
		return
	}
	defer resp.Body.Close()

	// Check for non-200 responses
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Request failed with status: %d - %s", resp.StatusCode, string(body))
		
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
		return
	}

	//fmt.Println("Raw response body:", string(body)) // Debugging step

	var responseMap map[string]interface{}
	if err := json.Unmarshal(body, &responseMap); err != nil {
		log.Printf("Error unmarshaling response body: %v. Raw response: %s", err, string(body))
		return
	}

	if responseMap["items"] != nil {
		fmt.Println("id")
		fmt.Println(responseMap["items"])
	}

	if responseMap["status"] != nil {
		fmt.Println("status")
		fmt.Println(responseMap["status"])
	}
	
}





func MercadoPagoHandler(c *gin.Context, access_token string) {
	cfg, err := config.New(access_token)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// request := payment.Request{
	// 	TransactionAmount: 10,
	// 	Description: "Pedido",
	// 	PaymentMethodID:   "pix",
	// 	Payer: &payment.PayerRequest{
	// 	   Email: "iananwobhm@gmail.com",
	// 	//    Identification: &payment.IdentificationRequest{
	// 	// 	  Type: "CPF",
	// 	// 	  Number: "19119119100",
	// 	//    },
	// 	},		 
	//  }

	// //payment_response = sdk.payment().create(payment_data, request_options)
	// //payment = payment_response["response"]

	// client := payment.NewClient(cfg)
	// result, err := client.Create(context.Background(), request)
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error()})
	// 	return
	// }

	// c.JSON(200, gin.H{"res": result})


	request := preference.Request{
		Items: []preference.ItemRequest{
			{
				ID:          "123",
				Title:       "Pedido",
				UnitPrice:   1,
				Quantity:    1,
				Description: "Combo",
				CurrencyID: "BRL",
				CategoryID: "food",
			},
		},
		PaymentMethods: &preference.PaymentMethodsRequest{
			ExcludedPaymentMethods: []preference.ExcludedPaymentMethodRequest{
				{
					ID: "bolbradesco",
				},
			},
		},
		NotificationURL: "https://in-drake-currently.ngrok-free.app/pagamento/atualizado",
		
		// AutoReturn: "all",
		// BackURLs: &preference.BackURLsRequest{
		// 	Success: "http://localhost:8080/pagamento/sucesso",
		// 	Pending: "http://localhost:8080/pagamento/pendente",
		// 	Failure: "http://localhost:8080/pagamento/falha",
		// },	
	}

	client := preference.NewClient(cfg)
	resource, err := client.Create(context.Background(), request)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"init_point": resource.InitPoint})
}

// func BuscarPagamento(c *gin.Context, access_token string, id string) {
// 	cfg, err := config.New(access_token)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}

// 	client := preference.NewClient(cfg)
// 	resource, err := client.Get(context.Background(), id)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(200, gin.H{"status": resource.InitPoint})
// }

func PagamentoAtualizado(c *gin.Context) {
	c.JSON(200, gin.H{"response": "success"})
}

