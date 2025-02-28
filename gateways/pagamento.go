package gateways

import (
	"context"

	"lanchonete/domain/repository"
	"lanchonete/infra/database/mongo"

)

type pagamentoGateway struct {
	database mongo.Database
	collection string
}

func NewPagamentoGateway(db mongo.Database, collection string) repository.ClienteRepository {
	return &clienteRepository{
		database: db,
		collection: collection,
	}
}

func (pg *pagamentoGateway) RealizarPagamento(c context.Context, pedido struct {
	ID          string
	Title       string
	UnitPrice   int
	Quantity    int
	Description string
	CurrencyID  string
}) error {
	collection := pg.database.Collection(pg.collection)
	_, err := collection.InsertOne(c, pedido)

	return err
}