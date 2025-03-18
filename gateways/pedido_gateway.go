package gateways

import (
	"context"
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
	"lanchonete/infra/database/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

type PedidoGateway struct {
	database   mongo.Database
	collection string
}

func NewPedidoGateway(db mongo.Database, collection string) repository.PedidoRepository {
	return &PedidoGateway{
		database:   db,
		collection: collection,
	}
}

func (pg *PedidoGateway) CriarPedido(c context.Context, pedido *entities.Pedido) error {
	collection := pg.database.Collection(pg.collection)
	_, err := collection.InsertOne(c, pedido)

	return err
}

func (pg *PedidoGateway) BuscarPedido(c context.Context, identificacao string) (entities.Pedido, error) {
	collection := pg.database.Collection(pg.collection)

	var pedido entities.Pedido

	err := collection.FindOne(c, bson.M{"identificacao": identificacao}).Decode(&pedido)
	return pedido, err
}

func (pg *PedidoGateway) AtualizarStatusPedido(c context.Context, identificacao string, status string) error {
	collection := pg.database.Collection(pg.collection)

	filter := bson.M{"identificacao": identificacao}
	update := bson.M{"$set": bson.M{"status": status}}

	_, err := collection.UpdateOne(c, filter, update)
	if err != nil {
		return err
	}

	var pedido entities.Pedido
	err = collection.FindOne(c, filter).Decode(&pedido)
	if err != nil {
		return err
	}

	return nil
}
