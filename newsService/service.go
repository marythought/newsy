package newsService

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Article models a top news story
type Article struct {
	ID     int    `json:"id,omitempty"`
	Title  string `json:"title"`
	Source string `json:"source"`
	URL    string `json:"url"`
	// Time   *time.Time `json:"time,omitempty"`
}

type TCArticles struct {
	Status   string      `json:"status"`
	Articles []TCArticle `json:"articles"`
}

type TCArticle struct {
	Source source     `json:"source"`
	Author string     `json:"author"`
	Title  string     `json:"title"`
	URL    string     `json:"url"`
	Time   *time.Time `json:"publishedAt"`
}

type source struct {
	Name string `json:"name"`
}

func (a TCArticle) toArticle() (art Article) {
	art.ID = 0
	art.Source = a.Source.Name
	art.Title = a.Title
	art.URL = a.URL
	return art
}

type NYTArticles struct {
	Articles []NYTArticle `json:"results"`
}

type NYTArticle struct {
	URL   string `json:"url"`
	Title string `json:"title"`
	// Time   *time.Time `json:"published_date"`
	Source string `json:"source"`
}

func (a NYTArticle) toArticle() (art Article) {
	art.ID = 0
	art.Source = a.Source
	art.Title = a.Title
	art.URL = a.URL
	return art
}

// GetArticles is a function that calls Hackernews, TechCrunch and NYT APIs to get, aggregate and return top articles
func GetArticles(w http.ResponseWriter, r *http.Request) {
	articles := getHackerNewsTopStories()
	articles = append(articles, getTechCrunchArticles()...)
	articles = append(articles, getNYTArticles()...)
	json.NewEncoder(w).Encode(articles)
}

func getHackerNewsTopStories() []Article {
	ids := getHackerNewsTopIDs()
	var articles []Article
	for _, art := range ids {
		articles = append(articles, getHNArticle(strconv.Itoa(art)))
	}
	return articles
}

func getHNArticle(id string) Article {
	art := Article{Source: "HackerNews"}
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
	art.ID = 0
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
