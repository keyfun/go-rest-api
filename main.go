package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Server Start on localhost:8000")
	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
