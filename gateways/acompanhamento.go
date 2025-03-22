package gateways

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
	"lanchonete/infra/database/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type acompanhamentoRepository struct {
	database          mongo.Database
	collection        string
	pedidosCollection string
}

func NewAcompanhamentoRepository(db mongo.Database, collection string) repository.AcompanhamentoRepository {
	return &acompanhamentoRepository{
		database:          db,
		collection:        collection,
		pedidosCollection: "pedido",
	}
}

func (ar *acompanhamentoRepository) CriarAcompanhamento(c context.Context, acompanhamento *entities.AcompanhamentoPedido) error {
	collection := ar.database.Collection(ar.collection)
	_, err := collection.InsertOne(c, acompanhamento)
	return err
}

func (ar *acompanhamentoRepository) BuscarPedidos(c context.Context, ID string) (entities.Pedido, error) {
	collection := ar.database.Collection(ar.pedidosCollection)
	fmt.Printf("Collection name: %s\n", ar.pedidosCollection)

	fmt.Printf("Buscando pedido com ID: %s\n", ID)

	filter := bson.M{
		"$or": []bson.M{
			{"identificacao": ID},
		},
	}

	var pedido entities.Pedido
	err := collection.FindOne(c, filter).Decode(&pedido)

	if err != nil {
		fmt.Printf("Erro ao buscar pedido: %v\n", err)
		return entities.Pedido{}, fmt.Errorf("pedido não encontrado: %v", err)
	}

	fmt.Printf("Pedido encontrado: %+v\n", pedido)
	return pedido, nil
}

func (ar *acompanhamentoRepository) AdicionarPedido(c context.Context, acompanhamento *entities.AcompanhamentoPedido, p *entities.Pedido) error {
	fmt.Printf("Tentando adicionar pedido %s ao acompanhamento %s\n", p.Identificacao, acompanhamento.ID)

	// First get the pedido
	pedido, err := ar.BuscarPedidos(c, p.Identificacao)
	if err != nil {
		return fmt.Errorf("erro ao buscar pedido: %v", err)
	}
	fmt.Printf("Pedido encontrado: %+v\n", pedido)

	// Get the current acompanhamento
	currentAcompanhamento, err := ar.BuscarAcompanhamento(c, acompanhamento.ID)
	if err != nil {
		return fmt.Errorf("acompanhamento não encontrado: %v", err)
	}
	fmt.Printf("Acompanhamento encontrado: %+v\n", currentAcompanhamento)

	// Add the pedido
	currentAcompanhamento.AdicionarPedido(pedido)
	fmt.Printf("Pedidos após adicionar: %+v\n", currentAcompanhamento.Pedidos.Listar())

	// Update in MongoDB
	filter := bson.M{"ID": acompanhamento.ID} // Try just the simple ID field first

	update := bson.M{
		"$set": bson.M{
			"pedidos":           currentAcompanhamento.Pedidos,
			"ultimaAtualizacao": time.Now(),
		},
	}

	fmt.Printf("Tentando atualizar com filter: %+v\n", filter)
	fmt.Printf("Update: %+v\n", update)

	result, err := ar.database.Collection(ar.collection).UpdateOne(c, filter, update)
	if err != nil {
		return fmt.Errorf("erro ao atualizar acompanhamento: %v", err)
	}

	fmt.Printf("Result: %+v\n", result)

	if result.ModifiedCount == 0 {
		// If first attempt failed, try with other ID field names
		filter = bson.M{
			"$or": []bson.M{
				{"_id": acompanhamento.ID},
				{"id": acompanhamento.ID},
				{"ID": acompanhamento.ID},
			},
		}
		result, err = ar.database.Collection(ar.collection).UpdateOne(c, filter, update)
		if err != nil {
			return fmt.Errorf("erro ao atualizar acompanhamento (segunda tentativa): %v", err)
		}
		if result.ModifiedCount == 0 {
			return fmt.Errorf("acompanhamento com ID %s não encontrado após múltiplas tentativas", acompanhamento.ID)
		}
	}

	return nil
}

func GenerateShortID() string {
	// Get current timestamp in Unix milliseconds
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	// Convert to base36 (numbers + lowercase letters) and take last 6 chars
	// This will give us something like "a1b2c3"
	id := fmt.Sprintf("acc%s", strconv.FormatInt(timestamp, 36)[len(strconv.FormatInt(timestamp, 36))-6:])
	return id
}

func (ar *acompanhamentoRepository) BuscarAcompanhamento(ctx context.Context, ID string) (entities.AcompanhamentoPedido, error) {
	var acompanhamento entities.AcompanhamentoPedido

	fmt.Printf("Buscando acompanhamento na collection: %s\n", ar.collection)

	filter := bson.M{}
	if ID != "" {
		filter["_id"] = ID // Use _id instead of ID
	}

	fmt.Printf("Usando filter: %+v\n", filter)

	err := ar.database.Collection(ar.collection).FindOne(ctx, filter).Decode(&acompanhamento)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			newAcompanhamento := entities.AcompanhamentoPedido{
				Pedidos:           entities.NewFilaPedidos(),
				ID:                GenerateShortID(),
				TempoEstimado:     time.Duration(15 * time.Minute),
				UltimaAtualizacao: time.Now(),
			}

			// Save the new acompanhamento
			_, err := ar.database.Collection(ar.collection).InsertOne(ctx, newAcompanhamento)
			if err != nil {
				return entities.AcompanhamentoPedido{}, fmt.Errorf("failed to create new acompanhamento: %v", err)
			}

			return newAcompanhamento, nil
		}
		return entities.AcompanhamentoPedido{}, fmt.Errorf("failed to find acompanhamento: %v", err)
	}

	fmt.Printf("Acompanhamento encontrado: %+v\n", acompanhamento)
	fmt.Printf("Pedidos na fila: %+v\n", acompanhamento.Pedidos.Listar())
	fmt.Printf("Número de pedidos na fila: %d\n", acompanhamento.Pedidos.Tamanho())

	acompanhamento.Pedidos.EnsureMutex()
	return acompanhamento, nil
}

func (ar *acompanhamentoRepository) AtualizarStatusPedido(ctx context.Context, acompanhamentoID string, identificacao string, novoStatus entities.StatusPedido) error {
	// Adicionar logs para debug
	fmt.Printf("Tentando atualizar pedido %s no acompanhamento %s para status %s\n",
		identificacao, acompanhamentoID, novoStatus)

	// Primeiro buscar o documento para confirmar a estrutura
	var doc bson.M
	err := ar.database.Collection(ar.collection).FindOne(ctx, bson.M{"ID": acompanhamentoID}).Decode(&doc)
	if err != nil {
		return fmt.Errorf("erro ao buscar acompanhamento: %v", err)
	}
	fmt.Printf("Documento encontrado: %+v\n", doc)

	// Tentar atualizar usando diferentes estruturas de caminho
	updates := []bson.M{
		// Tentativa 1: estrutura aninhada direta
		bson.M(
			bson.M{"identificacao": identificacao},
		)}
	// Tentativa 2: estrutura plana
	// {
	// 	"$set": bson.M{
	// 		"pedidos.$[elem].status":            novoStatus,
	// 		"pedidos.$[elem].ultimaAtualizacao": time.Now(),
	// 		"ultimaAtualizacao":                 time.Now(),
	// 	},
	// },

	// arrayFilters := options.Update().SetArrayFilters(options.ArrayFilters{
	// 	Filters: []interface{}{
	// 		bson.M{"elem.identificacao": identificacao},
	// 	},
	// })

	filter := bson.M{"ID": acompanhamentoID}

	// Tentar cada estrutura de update até uma funcionar
	var result *mongo.UpdateResult
	for i, update := range updates {
		fmt.Printf("Tentando update %d: %+v\n", i+1, update)
		result, err = ar.database.Collection(ar.collection).UpdateOne(ctx, filter, update, arrayFilters)
		if err == nil && result.ModifiedCount > 0 {
			break
		}
		fmt.Printf("Tentativa %d falhou: %v\n", i+1, err)
	}

	if err != nil {
		return fmt.Errorf("todas as tentativas de atualização falharam: %v", err)
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("pedido não encontrado na fila")
	}

	// Se chegou aqui, a atualização foi bem sucedida
	fmt.Printf("Atualização bem sucedida. ModifiedCount: %d\n", result.ModifiedCount)

	// Remover se finalizado
	if novoStatus == entities.Finalizado {
		removeUpdate := bson.M{
			"$pull": bson.M{
				"pedidos.pedidos": bson.M{"identificacao": identificacao},
			},
		}
		_, err = ar.database.Collection(ar.collection).UpdateOne(ctx, filter, removeUpdate)
		if err != nil {
			return fmt.Errorf("erro ao remover pedido finalizado: %v", err)
		}
	}

	return nil
}
