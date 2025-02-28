package gateways

import (
	"context"

	"lanchonete/domain/repository"
	"lanchonete/domain/entities"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

const collection = "clientes"

type clienteGateway struct {
	database mongo.Database
	collection string
}

func NewClienteGateway(db mongo.Database) repository.ClienteRepository {
	return &clienteGateway{
		database: db,
		collection: collection,
	}
}

func (cr *clienteGateway) CriarCliente(c context.Context, cliente *entities.Cliente) error {
	collection := cr.database.Collection(cr.collection)
	_, err := collection.InsertOne(c, cliente)

	return err
}

func (cr *clienteGateway) BuscarCliente(c context.Context, CPF string) (entities.Cliente, error) {
	collection := cr.database.Collection(cr.collection)

	var cliente entities.Cliente

	err := collection.FindOne(c, bson.M{"cpf": CPF}).Decode(&cliente)
	return cliente, err
}