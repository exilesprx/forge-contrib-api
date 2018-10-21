package database

import (
	"context"
	"log"
	"os"

	"github.com/mongodb/mongo-go-driver/mongo"
)

const database = "forge-contrib"

// Connection holds the mongo connection
type Connection struct {
	client *mongo.Client
	*mongo.Database
}

// CreateConnection creates a new connection
func CreateConnection() Connection {
	connection := Connection{}

	connection.connect()

	return connection
}

// Connect Create a client and connect to MongoDB
func (connection *Connection) connect() {

	client := createClient()

	err := client.Connect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	connection.client = client
	connection.Database = client.Database(database)
}

// Disconnect disconnect the mongo connection
func (connection *Connection) disconnect() {
	connection.client.Disconnect(nil)
}

func createClient() *mongo.Client {
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")

	client, err := mongo.NewClient("mongodb://" + host + ":" + port)

	if err != nil {
		log.Fatal(err)
	}

	return client
}
