package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/marythought/newsy/newsService"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	newsService.Setup()

	router := mux.NewRouter()
	router.HandleFunc("/news", newsService.GetArticles).Methods("GET")
	router.HandleFunc("/crawl", newsService.CrawlArticles).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
