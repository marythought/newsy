package newsService

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

// GetArticles is an endpoint that returns top articles from the db
func GetArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := dao.findAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, articles)
}

// CrawlArticles is an endpoint that calls Hackernews, TechCrunch and NYT APIs to get, aggregate and save top articles in the db
func CrawlArticles(w http.ResponseWriter, r *http.Request) {
	// TODO: could this happen in parallel? save the articles while getting the next news source...
	articles := getHackerNewsArticles()
	articles = append(articles, getTechCrunchArticles()...)
	articles = append(articles, getNYTArticles()...)
	for _, a := range articles {
		// don't save articles that have already been crawled
		// TODO: not always working, figure out why...
		dups, _ := dao.findByURL(a.URL)
		if len(dups) > 0 {
			continue
		}
		a.ID = bson.NewObjectId()
		if err := dao.insert(a); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	respondWithJSON(w, http.StatusCreated, articles)
}

func CleanArticles(w http.ResponseWriter, r *http.Request) {
	// TODO: delete old articles?

}
