package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marythought/newsy/newsService"
)

// Setup help via https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/news", newsService.GetArticles).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
