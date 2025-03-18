package gateways

import (
	"context"

	"lanchonete/domain/repository"
	"lanchonete/domain/entities"
	"lanchonete/infra/database/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

type PedidoRepository struct {
	database mongo.Database
	collection string
}

func NewPedidoRepository(db mongo.Database, collection string) repository.PedidoRepository {
	return &PedidoRepository{
		database: db,
		collection: collection,
	}
}

func (pr *PedidoRepository) CriarPedido(c context.Context, pedido *entities.Pedido) error {
	collection := pr.database.Collection(pr.collection)

	_, err := collection.InsertOne(c, pedido)
	return err
}

func (pr *PedidoRepository) BuscarPedido(c context.Context, nroPedido string) (entities.Pedido, error) {
	collection := pr.database.Collection(pr.collection)

	var pedido entities.Pedido

	err := collection.FindOne(c, bson.M{"nroPedido": nroPedido}).Decode(&pedido)
	return pedido, err
}
