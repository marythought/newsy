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
	router.HandleFunc("/news", newsService.GetArticles).Methods("GET", "OPTIONS")
	router.HandleFunc("/crawl", newsService.CrawlArticles).Methods("GET", "OPTIONS")
	router.HandleFunc("/clean", newsService.CleanDatabase).Methods("GET")
	log.Printf("Now listening on TCP port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
