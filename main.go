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

	ginEngine := gin.Default()

	route.Setup(env, db, ginEngine)

	ginEngine.Run(env.ServerAddress)
}