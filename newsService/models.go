package newsService

import (
	"fmt"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"
)

const dateForm = "2006-01-02"

// Article models a top news story for saving and retriving from the db
type Article struct {
	ID     bson.ObjectId `bson:"_id" json:"id,omitempty"`
	Title  string        `bson:"title" json:"title"`
	Source string        `bson:"source" json:"source"`
	URL    string        `bson:"url" json:"url"`
	Time   *time.Time    `json:"time,omitempty"`
}

// NewsyDAO is a Data Access Object to manage database operations
type NewsyDAO struct {
	Server   string
	Database string
}

// HNArticle models the JSON of an article from HackerNews API
type HNArticle struct {
	Title  string `json:"title"`
	Source string `json:"source"`
	URL    string `json:"url"`
	Time   uint64 `json:"time"`
}

func (a HNArticle) toArticle() (art Article) {
	art.Source = a.Source
	art.Title = a.Title
	art.URL = a.URL
	s := strconv.FormatUint(a.Time, 10)
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	tm := time.Unix(i, 0)
	art.Time = &tm
	return art
}

// TCArticles models the JSON of top articles from TechCrunch API
type TCArticles struct {
	Status   string      `json:"status"`
	Articles []tCArticle `json:"articles"`
}

type tCArticle struct {
	Source source     `json:"source"`
	Author string     `json:"author"`
	Title  string     `json:"title"`
	URL    string     `json:"url"`
	Time   *time.Time `json:"publishedAt"`
}

type source struct {
	Name string `json:"name"`
}

func (a tCArticle) toArticle() (art Article) {
	art.Source = a.Source.Name
	art.Title = a.Title
	art.URL = a.URL
	art.Time = a.Time
	return art
}

// NYTArticles models the JSON of top articles from New York Times API
type NYTArticles struct {
	Articles []nYTArticle `json:"results"`
}

type nYTArticle struct {
	URL    string `json:"url"`
	Title  string `json:"title"`
	Source string `json:"source"`
	Time   string `json:"published_date"`
}

func (a nYTArticle) toArticle() (art Article) {
	art.Source = a.Source
	art.Title = a.Title
	art.URL = a.URL
	t, err := time.Parse(dateForm, a.Time)
	if err != nil {
		fmt.Println(err)
	}
	art.Time = &t
	return art
}
