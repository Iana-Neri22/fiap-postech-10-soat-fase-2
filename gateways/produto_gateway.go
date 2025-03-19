package gateways

import (
	"context"
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
	"lanchonete/infra/database/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

type ProdutoGateway struct {
	database   mongo.Database
	collection string
}

func NewProdutoGateway(db mongo.Database, collection string) repository.ProdutoRepository {
	return &ProdutoGateway{
		database:   db,
		collection: collection,
	}
}

func (pd *ProdutoGateway) AdicionarProduto(c context.Context, produto *entities.Produto) error {
	collection := pd.database.Collection(pd.collection)
	_, err := collection.InsertOne(c, produto)

	return err
}

func (pd *ProdutoGateway) BuscarProdutoPorId(c context.Context, id string) (*entities.Produto, error) {
	collection := pd.database.Collection(pd.collection)

	var produto *entities.Produto

	err := collection.FindOne(c, bson.M{"identificacao": id}).Decode(&produto)
	return produto, err
}
