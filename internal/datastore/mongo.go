package datastore

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
	"time"
)

type DataStore struct {
	client   *mongo.Client
	database *mongo.Database
}

func ConnectMongo(mongoUri string, database string) *DataStore {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	slog.Info("Attempting to connect to mongo...")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))

	if err != nil {
		slog.Error("error: ", err)
	}

	if err := pingMongo(*client); err != nil {
		slog.Error("error: ", err)
	}

	slog.Info("Connected to mongo.")

	mongoClient := &DataStore{
		client:   client,
		database: client.Database(database),
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

func (m *DataStore) SetupDatabase(collections []string) {
	for _, collection := range collections {
		if err := m.createCollection(collection); err != nil {
			panic(err)
		}
	}
}

func (m *DataStore) createCollection(collection string) error {
	if err := m.database.CreateCollection(context.Background(), collection); err != nil {
		return err
	}

	return nil
}

func (m *DataStore) Insert(collection string, item interface{}) error {
	_, err := m.database.Collection(collection).InsertOne(context.Background(), item)

	if err != nil {
		return err
	}

	return nil
}

func (m *DataStore) FindById(collection string, document string, id int) (bson.M, error) {
	filter := bson.D{{Key: "_id", Value: id}}
	var result bson.M

	if err := m.database.Collection(collection).FindOne(context.Background(), filter).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
