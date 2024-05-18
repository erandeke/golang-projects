package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Isbn     string    `json:"isbn"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"LastName"`
}

//define slice of the movies

var movies []Movie

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	//set the headers to json
	log.Print("Inside gel all movies")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies) //read the w and encode it to stream of movies

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) //check this out the code for test-slice.go explained with an example
			break
		}
	}
	json.NewEncoder(w).Encode(movies)

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//loop over the movies and once the move found of that id simply encode in json and retruns it

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie

	//decode the body having movie from the request
	json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			//delete
			movies = append(movies[:index], movies[index+1:]...)
			//add one
			var movie Movie
			json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = strconv.Itoa(rand.Intn(1000))
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}

func main() {

	r := mux.NewRouter()

	//populate the movies
	movies = append(movies, Movie{ID: "1", Title: "Gotch", Isbn: "AA", Director: &Director{FirstName: "Raj", LastName: "Tevar"}})
	movies = append(movies, Movie{ID: "2", Title: "Game of thrones", Isbn: "AB", Director: &Director{FirstName: "Digaaz", LastName: "Tevar"}})
	movies = append(movies, Movie{ID: "3", Title: "Man vs wild", Isbn: "ABC", Director: &Director{FirstName: "Vipul", LastName: "Tevar"}})

	r.HandleFunc("/movies", getAllMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at %v\n", 7900)
	err := http.ListenAndServe(":7900", r)
	if err != nil {
		log.Fatal("Error ocurred while starting the web server %v\n", err)
	}

}
