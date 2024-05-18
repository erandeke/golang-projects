package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		fmt.Fprint(w, "Path is not supported", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		fmt.Fprint(w, "Http method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hey there")

}

func formHandler(w http.ResponseWriter, r *http.Request) {

	//parse Form
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, "Err occurred while parsing", http.StatusBadRequest)
		return
	}

	userName := r.FormValue("name")
	address := r.FormValue("address")
	fmt. (w, "userName  is: \n", userName, "address is : \n", address)

}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Server starting at port 8080\n")

	//create the server

	if err := http.ListenAndServe(":7100", nil); err != nil {
		log.Fatal(err)
	}

}
