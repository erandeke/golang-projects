package main

import (
	"fmt"
	"go-bookstore/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterRoute(r)
	http.Handle("/", r)
	fmt.Println("Server started gracefully")
	log.Fatal(http.ListenAndServe(":7200", r))

}
