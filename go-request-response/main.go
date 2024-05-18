package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	http.Handle("/", r)


	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hey there")
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		d, err := ioutil.ReadAll(r.Body) //Read from src till the EOF
		if err != nil {
			w.WriteHeader(http.StatusBadRequest) //Writeheader sends and Http response header provided with the status code
			return
		}
		fmt.Fprintf(w, "Data %s", d)
	})

	http.ListenAndServe(":9090", nil)

}
