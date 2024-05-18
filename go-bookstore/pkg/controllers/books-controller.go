package controllers

import (
	"encoding/json"
	"fmt"
	"go-bookstore/pkg/models"
	"go-bookstore/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// marshalling is converting go struct to json
// unmarshalling is converting json to go struct
func GetBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside GetBooks")
	w.Header().Set("Content-Type", "application/json")
	newBooks := models.GetAllBooks()
	json.NewEncoder(w).Encode(newBooks) //here in project marshalling is used check how its a difference
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	//get it from request
	params := mux.Vars(r)
	bookId := params["id"]          
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	json.NewEncoder(w).Encode(bookDetails)

}

func CreateBooks(w http.ResponseWriter, r *http.Request) {
	//create book var with pointer since we need to get it parsed and then call model to create book

	book := &models.Book{}
	utils.ParseBody(r, book) //convert json received into body of post method to go struct book
	b := book.CreateBook()   //add to db or insert to db
	json.NewEncoder(w).Encode(b)

}
