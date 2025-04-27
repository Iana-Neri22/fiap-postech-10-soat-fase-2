package main

import (
	"lanchonete/bootstrap"
	"lanchonete/internal/interfaces/http/server"
	_"lanchonete/docs"
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

	// Create and configure HTTP server
	srv := server.NewServer(app.Env, app.Mongo.Database(app.Env.DBName))
	srv.SetupRoutes()

	// Start server
	if err := srv.Start(); err != nil {
		panic(err)
	}
}