package controllers

import (
	"context"
	"errors"
	"go-ecommerce/database"
	"go-ecommerce/models"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	productCollection *mongo.Collection
	userCollection    *mongo.Collection
}

func NewApplication(productCollection, userCollection *mongo.Collection) *Application {
	return &Application{
		productCollection: productCollection,
		userCollection:    userCollection,
	}
}

func (app *Application) AddToCart() gin.HandlerFunc {

	return func(c *gin.Context) {

		//for every add to cart operation we need user id and product id that will be added to cart
		//each user -> product
		productID := c.Query("id") // query return the url query value otherwise it returns an empty string
		if productID == "" {
			log.Println("product id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
			return
		}

		//get the userId from the Query
		userId := c.Query("userid")
		if userId == "" {
			log.Println("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

		productObjectId, err := primitive.ObjectIDFromHex(productID)

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()

		errAdd := database.AddProductToCart(ctx, app.productCollection, app.userCollection, productObjectId, userId)
		if errAdd != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}

		c.IndentedJSON(200, "Successfully Added to the cart")
	}

}

func (app *Application) RemoveFromcart() gin.HandlerFunc {

	return func(c *gin.Context) {

		//for every add to cart operation we need user id and product id that will be added to cart
		//each user -> product
		productID := c.Query("id") // query return the url query value otherwise it returns an empty string
		if productID == "" {
			log.Println("product id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
			return
		}

		//get the userId from the Query
		userId := c.Query("userId")
		if userId == "" {
			log.Println("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

		productObjectId, err := primitive.ObjectIDFromHex(productID)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = database.RemoveFromCart(ctx, app.productCollection, app.userCollection, productObjectId, userId)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}

		c.IndentedJSON(200, "Successfully Removed from the cart")

	}
}

func (app *Application) GetItemsFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Query("id")
		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid id"})
			c.Abort()
			return

		}

		usert_id, _ := primitive.ObjectIDFromHex(user_id)

		cont, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()

		var filledCart models.Users

		err := UserCollection.FindOne(cont, bson.D{primitive.E{Key: "_id", Value: usert_id}}).Decode(&filledCart)

		if err != nil {
			log.Println(err)
			c.IndentedJSON(500, "not id found")
			return
		}

		//$match in mongo returns the document that matches the given query and it passes to next pipeline stage
		filter_match := bson.D{{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: usert_id}}}}
		unwind := bson.D{{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$usercart"}}}}
		grouping := bson.D{{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value: "$_id"}, {Key: "total", Value: bson.D{primitive.E{Key: "$sum", Value: "$usercart.price"}}}}}}
		pointcursor, err := UserCollection.Aggregate(cont, mongo.Pipeline{filter_match, unwind, grouping})
		if err != nil {
			log.Println(err)
		}
		var listing []bson.M
		if err = pointcursor.All(c, &listing); err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		for _, json := range listing {
			c.IndentedJSON(200, json["total"])
			c.IndentedJSON(200, filledCart.UserCart)
		}
		cont.Done()

	}
}

func (app *Application) InstantBuy() gin.HandlerFunc {

}

func (app *Application) BuyFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {

		userQueryId := c.Query("userId")
		if userQueryId == "" {
			log.Panic("user id is empty")
			c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		err := database.BuyItemsFromTheCart(ctx, app.userCollection, userQueryId)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(200, "Successfully Placed the order")

	}

}
