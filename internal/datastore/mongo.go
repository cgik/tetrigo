package datastore

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	client *mongo.Client
}

func ConnectMongo() *MongoStore {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Print(err)
		panic(err)
	}

	var result bson.M

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		log.Print(err)
		panic(err)
	}

	log.Println("Connected to mongo.")

	mongo := &MongoStore{
		client: client,
	}

	return mongo
}

func (s *MongoStore) Insert(item interface{}) error {
	return nil
}

func (s *MongoStore) FindById(document string, id int) error {
	return nil
}
