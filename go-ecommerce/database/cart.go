package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddProductToCart(ctx context.Context, productCollection, userCollection *mongo.Collection, productId primitive.ObjectID, userId string) error {

}

func RemoveFromCart(ctx context.Context, productCollection, userCollection *mongo.Collection, productId primitive.ObjectID, userId string) error {

}

func BuyItemsFromTheCart(ctx context.Context, userCollection *mongo.Collection, userId string) error {

}
