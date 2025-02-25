package bootstrap

import (
	"context"
	"log"

	"lanchonete/mongo"
)

func NewDatabase(env *Env) mongo.Client {

	client, err := mongo.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func CloseConnection(client mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to database closed.")
}