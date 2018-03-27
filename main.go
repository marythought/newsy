package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/marythought/newsy/newsService"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	newsService.Setup()

	router := mux.NewRouter()
	router.HandleFunc("/news", newsService.GetArticles).Methods("GET")
	router.HandleFunc("/crawl", newsService.CrawlArticles).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
