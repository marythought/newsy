package newsService

import (
	"fmt"
	"log"
	"os"
	"time"

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

// remove articles from the db 30+ days old
func (m *NewsyDAO) deleteOld() (int, error) {
	m.connect()
	info, err := db.C(m.Database).RemoveAll(bson.M{"time": bson.M{"$lt": time.Now().Add(-time.Hour * 24 * 30)}})
	if err != nil {
		fmt.Println(err)
	}
	return info.Removed, err
}
