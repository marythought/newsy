package newsService

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Article models a top news story
type Article struct {
	ID     int    `json:"id,omitempty"`
	Title  string `json:"title"`
	Source string `json:"source"`
	URL    string `json:"url"`
	// Time   *time.Time `json:"time,omitempty"`
}

// GetArticles is a function that calls Hackernews, TechCrunch and NYT APIs to get, aggregate and return top articles
func GetArticles(w http.ResponseWriter, r *http.Request) {
	articles := getHackerNewsTopStories()
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

// /v0/item/<id>
func getHNArticle(id string) Article {
	// For example, a story: GET https://hacker-news.firebaseio.com/v0/item/8863.json
	// returns {
	// 	"id":8863,
	// 	"score":104,
	// 	"time":1175714200,
	// 	"title":"My YC app: Dropbox - Throw away your USB drive",
	// 	"url":"http://www.getdropbox.com/u/2/screencast.html"
	// }
	art := Article{Source: "HackerNews"}
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/item/" + id + ".json")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	err = json.Unmarshal(body, &art)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(art)
	return art
}

// /v0/topstories  -- returns an array of IDs (up to 500)
func getHackerNewsTopIDs() []int {
	var ids []int
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &ids)
	fmt.Println(ids[:10])
	return ids[:10]
}

func getTechCrunch() {
	// https://newsapi.org/v2/top-headlines?sources=techcrunch&apiKey=
	// returns:
	// {
	// "status": "ok",
	// "totalResults": 10,
	// -"articles": [
	// -{
	// -"source": {
	// "id": "techcrunch",
	// "name": "TechCrunch"
	// },
	// "author": "Matt Burns",
	// "title": "BMW and Lexus look to car subscriptions",
	// "description": "More automakers will soon offer vehicles through subscription services. Lexus today announced its upcoming UX crossover would be available through one and Bloomberg published a report today stating BMW is about to announce a subscription pilot. These automake…",
	// "url": "https://techcrunch.com/2018/03/26/bmw-and-lexus-look-to-car-subscriptions/",
	// "urlToImage": "https://techcrunch.com/wp-content/uploads/2018/03/lexus-ux.jpg?w=649",
	// "publishedAt": "2018-03-26T23:20:16Z"
	// },
	// -{
	// -"source": {
	// "id": "techcrunch",
	// "name": "TechCrunch"
	// },
	// "author": "Katie Roof",
	// "title": "Microsoft surges 8% after Morgan Stanley says it will reach $1 trillion market cap",
	// "description": "The Dow surged 669 points on Monday after trade tensions eased. Tech stocks like Amazon and Apple saw gains, but the biggest winner of all was Microsoft . The Seattle tech giant, which is a Dow 30 company, benefitted not only from the solid stock market day, …",
	// "url": "https://techcrunch.com/2018/03/26/microsoft-surges-8-after-morgan-stanley-says-it-will-reach-1-trillion-market-cap/",
	// "urlToImage": "https://techcrunch.com/wp-content/uploads/2017/07/gettyimages-496394442.jpg?w=643",
	// "publishedAt": "2018-03-26T22:11:38Z"
	// },
	// ]
	// }
}

func getNYT() {

}
