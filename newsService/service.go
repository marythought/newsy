package newsService

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var dao NewsyDAO
var db *mgo.Database

// Setup sets the global Data Access Object with the appropriate values, it is called from main.go after env vars are parsed
func Setup() {
	dao.Server = "mongodb://" + os.Getenv("MONGO_USER") + ":" + os.Getenv("MONGO_PWORD") + "@ds125469.mlab.com:25469/newsy_db"
	dao.Database = "newsy_db"
}

func (m *NewsyDAO) connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *NewsyDAO) findAll() ([]Article, error) {
	m.connect()
	var articles []Article
	err := db.C(m.Database).Find(bson.M{}).Sort("-time").All(&articles)
	if err != nil {
		fmt.Println(err)
	}
	return articles, err
}

func (m *NewsyDAO) findByURL(url string) ([]Article, error) {
	m.connect()
	var articles []Article
	err := db.C(m.Database).Find(bson.M{"url": url}).All(&articles)
	if err != nil {
		fmt.Println(err)
	}
	return articles, err
}

func (m *NewsyDAO) insert(article Article) error {
	m.connect()
	err := db.C(m.Database).Insert(&article)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (m *NewsyDAO) delete(article Article) error {
	m.connect()
	err := db.C(m.Database).Remove(&article)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
