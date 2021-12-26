package main

import (
	"context"
	"log"
	"time"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func GetDB() *mongo.Database {
	clientOptions := options.Client().
    ApplyURI(os.Getenv("CONNECTION"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
	    log.Fatal(err)
	}
	return client.Database(os.Getenv("DB_NAME"))
}