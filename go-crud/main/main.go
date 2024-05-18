package main

import (
	"go-crud/movieroutes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	log.Println("Here I am in the crud app in GO langauage....")
	r := mux.NewRouter()
	movieroutes.RegisterRoute(r)
	http.Handle("/", r) //not sure on this
	log.Fatal(http.ListenAndServe(":7200", r))

}
