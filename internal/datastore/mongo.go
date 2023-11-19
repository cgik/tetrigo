package datastore

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoSession struct {
	client *mongo.Client
}

func ConnectMongo(mongoUri string) *MongoSession {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("Attempting to connect to mongo...")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))

	if err != nil {
		log.Fatal(err)
	}

	if err := pingMongo(*client); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to mongo.")

	mongoClient := &MongoSession{
		client: client,
	}

	return mongoClient
}

func pingMongo(client mongo.Client) error {
	var result bson.M

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		return err
	}

	return nil
}

func (m *MongoSession) Insert(collection string, item interface{}) error {
	return nil
}

func (m *MongoSession) FindById(collection string, document string, id int) error {
	return nil
}
