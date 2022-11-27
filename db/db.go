package db

import (
	"TodoAPI/configs"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoInstance struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func NewDBConnection() mongoInstance {
	client, err := mongo.NewClient(options.Client().ApplyURI(configs.GetEnvMongoURI()))
	if err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}

	mongoInstance := mongoInstance{
		Client:   client,
		Database: client.Database(configs.GetEnvDatabaseName()),
	}

	return mongoInstance
}
