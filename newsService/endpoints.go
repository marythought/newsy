package newsService

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

// GetArticles is an endpoint that returns top articles from the db
func GetArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, articles)
}

// CrawlArticles is an endpoint that calls Hackernews, TechCrunch and NYT APIs to get, aggregate and save top articles in the db
func CrawlArticles(w http.ResponseWriter, r *http.Request) {
	articles := getHackerNewsArticles()
	articles = append(articles, getTechCrunchArticles()...)
	articles = append(articles, getNYTArticles()...)
	// save articles in db
	for _, a := range articles {
		a.ID = bson.NewObjectId()
		if err := dao.Insert(a); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	// TODO: delete old articles?
	respondWithJSON(w, http.StatusCreated, articles)
}
