package database

import (
	"context"
	"log"

	"github.com/exilesprx/forge-contrib-api/models"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var client mongo.Client

// Connect Create a client and connect to MongoDB
func Connect() mongo.Client {
	client = createClient()

	var err error = client.Connect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func createClient() mongo.Client {
	client, err := mongo.NewClient("mongodb://mongo:27017")

	if err != nil {
		log.Fatal(err)
	}

	return *client
}

// CreateUser Creates a new user in the database
func CreateUser(user *models.User) interface{} {
	var collection *mongo.Collection = client.Database("forge-contrib").Collection("user")

	res, err := collection.InsertOne(context.Background(), user)

	if err != nil {
		log.Fatal(err)
	}

	return res.InsertedID
}
