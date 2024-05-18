package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectToDb(Url string) *mongo.Client {

	//set the time out using context

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	//connect to mongo db instance using context

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(Url))

	if err != nil {

		log.Fatal(err)
	}

	defer cancel()

	//test the connection

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	return client
}

var Client *mongo.Client = ConnectToDb("mongodb://localhost:27017")

func CloseConnection(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {

	//cancel the context
	defer cancel()

	//disconnect
	defer func() {

		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}

	}()
}

func GetUserData(client *mongo.Client, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database("Ecommerce").Collection(collectionName)
	return collection

}

func GetProductData(client *mongo.Client, collectionName string) *mongo.Collection {
	var productCollection *mongo.Collection = client.Database("Ecommerce").Collection(collectionName)
	return productCollection

}
