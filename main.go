package main

import (
	route "lanchonete/api/route"
	"lanchonete/bootstrap"
	_"lanchonete/docs"
	"github.com/gin-gonic/gin"
)

// @title Lanchonete API - Tech Challenge 2
// @version 1.0
// @description API para o Tech Challenge 2 da FIAP - SOAT

// @host localhost:8080
// @BasePath /
func main() {

	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	ginEngine := gin.Default()

	route.Setup(env, db, ginEngine)

	ginEngine.Run(env.ServerAddress)
}