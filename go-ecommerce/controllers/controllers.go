package controllers

import (
	"context"
	"fmt"
	"go-ecommerce/database"
	"go-ecommerce/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *mongo.Collection = database.GetUserData(database.Client, "Users")
var ProductCollection *mongo.Collection = database.GetProductData(database.Client, "Product")

func hashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func verifyPassword(userPassword string, userProvidedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(userProvidedPassword), []byte(userPassword))
	isPasswordValid := true
	msg := ""
	if err != nil {
		msg = "The password is incorrect"
		isPasswordValid = false
		log.Panic(err)
	}
	return isPasswordValid, msg

}

func SignUp() gin.HandlerFunc {

	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var user models.Users
		//it decodes the json received into user struct
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//handle the validation
		//validationErr := Validate.struct(user)

		//check here on querying the database if the user exist by returning the count
		count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		}

		//PhoneNo

		count, err = UserCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Phone is already in use"})
			return
		}

		//generate hash password
		password := hashPassword(*user.Password)
		user.Password = &password

		user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		//new Id
		user.ID = primitive.NewObjectID() // generates new object Id

		user.User_Id = user.ID.Hex() // get the hex of the ID

		user.UserCart = make([]models.ProductUser, 0) //slice of length 0 i,e empty slice
		user.Address_Details = make([]models.Address, 0)
		user.Order_Status = make([]models.Order, 0)
		//insert these details in DB

		_, insertErr := UserCollection.InsertOne(ctx, user)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "not created"})
			return
		}
		defer cancel()
		c.JSON(http.StatusCreated, "Successfully Signed Up!!")

	}

}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var user models.Users
		var foundUser models.Users

		err := UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "login or password incorrect"})
			return
		}
		PasswordValid, msg := verifyPassword(*user.Password, *foundUser.Password)
		if !PasswordValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			fmt.Println(msg)
			return
		}
		c.JSON(http.StatusFound, foundUser)

	}

}

func searchProducts() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var products []models.Product

		//get it from database
		ct, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()

		//get the data in json format cursor here hold the json format
		cursor, err := ProductCollection.Find(ct, bson.D{})
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, "Someting Went Wrong Please Try After Some Time")
			return
		}

		// All iterates over the json document and decodes into product struct
		err = cursor.All(ctx, &products) //TODO check on putting just products
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		defer cursor.Close(ctx)
		ctx.IndentedJSON(200, products)
	}

}

func searchByProductQuery() gin.HandlerFunc {

	return func(c *gin.Context) {

		var searchProducts []models.Product
		product_name := c.Query("name")

		if product_name == "" {
			log.Fatal("Product name is empty")
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		searchQueryRes, err := ProductCollection.Find(ctx, bson.M{"product_name": bson.M{"$regex": product_name}})

		if err != nil {
			c.IndentedJSON(404, "something went wrong in fetching the dbquery")
			return
		}

		err = searchQueryRes.All(ctx, &searchProducts)

		if err != nil {
			log.Println(err)
			c.IndentedJSON(400, "invalid")
			return
		}

		defer searchQueryRes.Close(ctx)

		c.IndentedJSON(200, searchProducts)

	}

}
