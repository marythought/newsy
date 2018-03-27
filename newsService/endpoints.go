package newsService

import (
	"encoding/json"
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
	err := insertArticles(articles)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	respondWithJSON(w, http.StatusCreated, articles)
}

func insertArticles(articles []Article) error {
	for _, a := range articles {
		// don't save articles that have already been crawled
		dups, _ := dao.findByURL(a.URL)
		if len(dups) > 0 {
			continue
		}
		a.ID = bson.NewObjectId()
		err := dao.insert(a)
		if err != nil {
			return err
		}
	}
	return nil
}

func CleanArticles(w http.ResponseWriter, r *http.Request) {
	// TODO: delete old articles?
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(response)
}
