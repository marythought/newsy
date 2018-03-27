package newsService

import (
	"encoding/json"
	"net/http"
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
	// delete old articles?
	json.NewEncoder(w).Encode(articles)
}

// func CreateArticle(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()
// 	var movie Movie
// 	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}
// 	movie.ID = bson.NewObjectId()
// 	if err := dao.Insert(movie); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondWithJson(w, http.StatusCreated, movie)
// }
