package threads

import (
	"context"
	"disspace/business/threads"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBThreadRepository struct {
	Conn *mongo.Database
}

func NewMongoDBThreadRepository(conn *mongo.Database) threads.Repository {
	return &MongoDBThreadRepository{
		Conn: conn,
	}
}

func (repository *MongoDBThreadRepository) GetAll(ctx context.Context) ([]threads.Domain, error) {
	var result []threads.Domain
	cursor, err := repository.Conn.Collection("threads").Find(ctx, bson.M{})
	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &result); err != nil {
		return []threads.Domain{}, err
	}
	
	return result, nil
}

func (repository *MongoDBThreadRepository) Create(ctx context.Context, threadDomain *threads.Domain) (threads.Domain, error) {
	thread := FromDomain(*threadDomain)

	cursor, err := repository.Conn.Collection("threads").InsertOne(ctx, thread)
	if err != nil {
		return threads.Domain{}, err
	}

	threadId := cursor.InsertedID.(primitive.ObjectID).Hex()
	return threads.Domain{ID: threadId}, nil
}















func (repository *MongoDBThreadRepository) Delete(ctx context.Context, id string) error {
	convert, errorConvert := primitive.ObjectIDFromHex(id)
	if errorConvert != nil {
		return errorConvert
	}
	filter := bson.D{{Key: "_id", Value: convert}}
	_, err := repository.Conn.Collection("threads").DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}