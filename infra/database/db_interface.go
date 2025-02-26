package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

// Define the Database interface
type Database interface {
	InsertOne(ctx context.Context, collection string, document interface{}) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, collection string, filter interface{}, result interface{}) error
}
