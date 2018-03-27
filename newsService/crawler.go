package newsService

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func getHackerNewsArticles() []Article {
	ids := getHackerNewsTopIDs()
	var articles []Article
	for _, art := range ids {
		articles = append(articles, getHNArticle(strconv.Itoa(art)).toArticle())
	}
	return articles
}

func getHNArticle(id string) HNArticle {
	art := HNArticle{Source: "HackerNews"}
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/item/" + id + ".json")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &art)
	if err != nil {
		fmt.Println(err)
	}
	return art
}

func getHackerNewsTopIDs() []int {
	var ids []int
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &ids)
	if err != nil {
		fmt.Println(err)
	}
	return ids[:10]
}

func getTechCrunchArticles() []Article {
	articles := []Article{}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	token := os.Getenv("TC_TOKEN")

	m := TCArticles{}
	resp, err := http.Get("https://newsapi.org/v2/top-headlines?sources=techcrunch&apiKey=" + token)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &m)
	if err != nil {
		fmt.Println(err)
	}
	for _, a := range m.Articles {
		articles = append(articles, a.toArticle())
	}
	return articles
}

func getNYTArticles() []Article {
	articles := []Article{}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	token := os.Getenv("NYT_TOKEN")

	m := NYTArticles{}
	resp, err := http.Get("https://api.nytimes.com/svc/mostpopular/v2/mostviewed/Technology/7.json?api-key=" + token)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &m)
	if err != nil {
		fmt.Println(err)
	}
	for _, a := range m.Articles {
		articles = append(articles, a.toArticle())
	}
	return articles
}
