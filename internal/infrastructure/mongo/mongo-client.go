package mongo

import (
	"context"
	"log"

	"github.com/mineamihai2001/game-night/internal/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var instance *MongoClient = nil

type MongoClient struct {
	ctx    context.Context
	client *mongo.Client
	db     *mongo.Database
}

func connect(ctx context.Context) *mongo.Client {
	env := helpers.Env()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(env.Db.ConnectionString))

	defer func() {
		if err == nil {
			return
		}
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal("[ERROR] - disconnected", err)
		}
	}()

	return client
}

func createInstance(database string, context context.Context) *MongoClient {
	client := connect(context)

	db := client.Database(database)

	return &MongoClient{
		ctx:    context,
		client: client,
		db:     db,
	}
}

func GetInstance(database string, context context.Context) *MongoClient {
	if instance == nil {
		instance = createInstance(database, context)
	}

	return instance
}

func GetCollection[T interface{}](client *MongoClient, name string, options ...*options.CollectionOptions) *Collection[T] {
	coll := client.db.Collection(name)

	return &Collection[T]{
		collection: coll,
		ctx:        client.ctx,
	}
}
