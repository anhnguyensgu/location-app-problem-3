package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Event struct {
	Timezone  string `validate:"required"`
	Email     string `validate:"required"`
	IpAddress string `validate:"required"`
	Latitude  float64
	Longitude float64
	createdAt primitive.Timestamp
}

func ConnectMongodb() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, error := mongo.Connect(context.TODO(), clientOptions)

	if error != nil {
		log.Fatal("error")
	}

	err := client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal("err 2")
	}
	return client
}
