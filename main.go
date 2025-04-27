package main

import (
	"lanchonete/bootstrap"
	"lanchonete/internal/interfaces/http"
	_"lanchonete/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @title Lanchonete API - Tech Challenge 2
// @version 1.0
// @description API para o Tech Challenge 2 da FIAP - SOAT

// @host localhost:8080
// @BasePath /
//go:generate go run github.com/swaggo/swag/cmd/swag@latest init
func main() {
	// Initialize application
	app := bootstrap.App()
	defer app.CloseDBConnection()

	// Setup HTTP server
	ginEngine := gin.Default()
	ginEngine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Setup routes
	http.Setup(app.Env, app.Mongo.Database(app.Env.DBName), ginEngine)

	// Start server
	ginEngine.Run(app.Env.ServerAddress)
}