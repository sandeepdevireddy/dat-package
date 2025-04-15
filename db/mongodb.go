package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var PackagesCollection *mongo.Collection
var PackageComponentsCollection *mongo.Collection

// SetupDatabase initializes the MongoDB client and collection
func SetupDatabase() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://verifyhire:vh12345@verifyhire.outzf6f.mongodb.net/?retryWrites=true&w=majority&appName=verifyhire")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	Client = client
	PackagesCollection = client.Database("verifyhireDB").Collection("packages")
	PackageComponentsCollection = client.Database("verifyhireDB").Collection("package_components")
}
