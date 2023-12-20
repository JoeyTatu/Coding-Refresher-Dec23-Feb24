package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
	MongoDb := "mongodb://localhost:27017"
	fmt.Print(MongoDb)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
		return nil
	}

	fmt.Println("Connected to MongoDB.")

	return client
}

var Client *mongo.Client = DBInstance()

var FoodCollection *mongo.Collection = OpenCollection(Client, "food")
var UserCollection *mongo.Collection = OpenCollection(Client, "user")
var InvoiceCollection *mongo.Collection = OpenCollection(Client, "invoice")
var MenuCollection *mongo.Collection = OpenCollection(Client, "menu")
var OrderCollection *mongo.Collection = OpenCollection(Client, "order")
var OrderItemCollection *mongo.Collection = OpenCollection(Client, "orderItem")
var TableCollection *mongo.Collection = OpenCollection(Client, "table")

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant").Collection(collectionName)
	return collection
}
