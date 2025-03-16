package gateways

import (
	"context"

	"lanchonete/domain/repository"
	"lanchonete/domain/entities"
	"lanchonete/infra/database/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

type acompanhamentoRepository struct {
	database mongo.Database
	collection string
}

func NewAcompanhamentoRepository(db mongo.Database, collection string) repository.AcompanhamentoRepository {
	return &acompanhamentoRepository{
		database: db,
		collection: collection,
	}
}

func (ar *acompanhamentoRepository) CriarAcompanhamento(c context.Context, acompanhamento *entities.AcompanhamentoPedido) error {
	collection := ar.database.Collection(ar.collection)

	_, err := collection.InsertOne(c, acompanhamento)
	return err
}

func (ar *acompanhamentoRepository) BuscarPedidos(c context.Context, nroPedido string) (entities.Pedido, error) {
	collection := ar.database.Collection(ar.collection)

	var pedido entities.Pedido

	err := collection.FindOne(c, bson.M{"nroPedido": nroPedido}).Decode(&pedido)
	return pedido, err
}
