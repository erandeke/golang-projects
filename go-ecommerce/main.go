package main

import (
	"go-ecommerce/controllers/"
	"go-ecommerce/database"
	"go-ecommerce/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	//get the port from os

	port := os.Getenv("PORT")

	//check if port is empty
	if port == "'" {
		port = "8000"
	}

	//get the product data from product collection and user data from users collection
	app := controllers.NewApplication(database.GetProductData(database.Client, "Products"), database.GetUserData(database.Client, "Users"))

	//create the router

	router := gin.New()

	//set the routes

	routes.UserRoutes(router)

	//attach the middleware

	//router.Use(middleware.Authentication()) todo

	//add the routes

	router.GET("/addToCart", app.AddToCart())
	router.GET("/removeFromCart", app.RemoveFromCart())
	router.GET("cartcheckout", app.BuyFromCart())

	//run the server

	log.Fatal(router.Run(":" + port))

}
