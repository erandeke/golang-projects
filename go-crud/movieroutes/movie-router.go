package movieroutes

import (
	"go-crud/moviecontrollers"

	"github.com/gorilla/mux"
)

func RegisterRoute(routes *mux.Router) {

	routes.HandleFunc("/movies", moviecontrollers.CreateMovies).Methods("POST")
	//routes.HandleFunc("/movies/{id}", moviecontrollers.GetMoviesById).Methods("GET")
	//routes.HandleFunc("/movies{name}", moviecontrollers.GetMoviesByName).Methods("GET")

}
