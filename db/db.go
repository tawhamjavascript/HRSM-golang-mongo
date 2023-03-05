package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

const mongoURI = "mongodb+srv://<USER>:<PASSWORD>@banco.8arn9hb.mongodb.net/?retryWrites=true&w=majority"

func Connect() *MongoInstance {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		panic(fmt.Sprintf("Error to connect with database. Fix it"))

	}
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(fmt.Sprintf("Connect with database did not happen "))

	}
	db := client.Database("fiber-hrms")
	mg := MongoInstance{
		Client: client,
		Db:     db,
	}
	return &mg
}
