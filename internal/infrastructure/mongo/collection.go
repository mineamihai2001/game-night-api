package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Schema interface{}

type Collection[T Schema] struct {
	collection *mongo.Collection
	ctx        context.Context
}

func CreateCollection[T interface{}](ctx context.Context, collection *mongo.Collection) *Collection[T] {
	return &Collection[T]{
		collection: collection,
		ctx:        ctx,
	}
}

func (coll *Collection[T]) FindOne(filter interface{}, opts ...*options.FindOneOptions) (T, error) {
	data := coll.collection.FindOne(coll.ctx, filter, opts...)

	var doc T
	err := data.Decode(&doc)

	return doc, err
}

func (coll *Collection[T]) Find(filter interface{}, opts ...*options.FindOptions) ([]T, error) {
	cursor, err := coll.collection.Find(coll.ctx, filter, opts...)
	if err != nil {
		return nil, err
	}

	var result []T = make([]T, 0)
	if err = cursor.All(coll.ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (coll *Collection[T]) InsertOne(document T, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return coll.collection.InsertOne(coll.ctx, document, opts...)
}

func (coll *Collection[T]) InsertMany(documents []T, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	interfaces := make([]interface{}, len(documents))
	for i := range documents {
		interfaces[i] = documents[i]
	}
	return coll.collection.InsertMany(coll.ctx, interfaces, opts...)
}

func (coll *Collection[T]) UpdateOne(filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return coll.collection.UpdateOne(coll.ctx, filter, update, opts...)
}

func (coll *Collection[T]) UpdateMany(filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return coll.collection.UpdateMany(coll.ctx, filter, update, opts...)
}

func (coll *Collection[T]) DeleteOne(filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return coll.collection.DeleteOne(coll.ctx, filter, opts...)
}

func (coll *Collection[T]) DeleteMany(filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return coll.collection.DeleteMany(coll.ctx, filter, opts...)
}

func (coll *Collection[T]) Aggregate(pipeline interface{}, opts ...*options.AggregateOptions) ([]T, error) {
	cursor, err := coll.collection.Aggregate(coll.ctx, pipeline, opts...)
	if err != nil {
		return nil, err
	}

	var result []T = make([]T, 0)
	if err = cursor.All(coll.ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}
