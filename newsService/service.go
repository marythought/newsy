package newsService

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db *mgo.Database
var dao = NewsyDAO{
	Server:   "localhost",
	Database: COLLECTION,
}

const (
	COLLECTION = "newsy"
)

func (m *NewsyDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *NewsyDAO) FindAll() ([]Article, error) {
	m.Connect()
	var articles []Article
	fmt.Println(db)
	err := db.C(COLLECTION).Find(bson.M{}).All(&articles)
	if err != nil {
		fmt.Println(err)
	}
	return articles, err
}

func (m *NewsyDAO) FindById(id string) (Article, error) {
	var article Article
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&article)
	if err != nil {
		fmt.Println(err)
	}
	return article, err
}

func (m *NewsyDAO) Insert(article Article) error {
	err := db.C(COLLECTION).Insert(&article)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (m *NewsyDAO) Delete(article Article) error {
	err := db.C(COLLECTION).Remove(&article)
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
