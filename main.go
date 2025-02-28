package main

import (
	route "lanchonete/api/route"
	"lanchonete/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	gin := gin.Default()

	route.Setup(env, db, gin)

	gin.Run(env.ServerAddress)
}