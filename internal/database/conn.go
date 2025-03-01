package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongoDB() (*mongo.Client, error) {
	clienetOption := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clienetOption)
	if err != nil {
		return &mongo.Client{}, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return &mongo.Client{}, err
	}
	log.Println("Connected to MongoDB!")
	return client, nil
}
