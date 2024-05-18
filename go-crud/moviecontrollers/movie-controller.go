package moviecontrollers

import (
	"encoding/json"
	"fmt"
	"go-crud/models"
	"go-crud/utils"
	"net/http"
)

func CreateMovies(w http.ResponseWriter, r *http.Request) {

	movie := &models.Movies{}
	//write parser here
	utils.ParserequestToAnyDataModel(r, movie)
	m := movie.CreateMovie()
	fmt.Println("the movie is created as {}", m)
	//have the encoder
	json.NewEncoder(w).Encode(movie)

}
