# Newsy App

This is a RESTful Go API that returns top news articles from several tech news sources.

It supports three endpoints currently:
1. https://newsy-bits.herokuapp.com/crawl -- Calls news source APIs and saves articles in MongoDB database, returns a count of number of articles added to DB
2. https://newsy-bits.herokuapp.com/news -- Gets all top news stories saved in the MongoDB database and return articles as JSON in decending date order
3. https://newsy-bits.herokuapp.com/clean -- Removes entries from DB 30+ days old, returns a count of number of articles deleted

See it in action: http://www.marydickson.info/news

In the future, I would add a job or recurring task to regularly crawl for new articles and clean up DB (currently only works on-demand), and add pagination / only return a specified number of articles.

TODO:
- limit number of responses for get news endpoint
- add caching, and/or some way of checking when APIs were last crawled (add to the JSON?)
- add a job that crawls the APIs on a recurring schedule (2x per day?) and removes old entries from DB
- better error handling and logging
- add tests

Questions for Consideration:
- do we want an even mix of the news sources, maybe the latest 10 for each? if yes, return the appropriate json
- how often to crawl for new news and remove old news?

Resources:
- https://www.mongodb.com/blog/post/building-your-first-application-mongodb-creating-rest-api-using-mean-stack-part-1
- http://www.blog.labouardy.com/build-restful-api-in-go-and-mongodb/
- https://medium.com/@IndianGuru/go-mongodb-mongolab-mgo-and-heroku-d411b5ac53f9
- https://github.com/kardianos/govendor (required for Heroku)
- https://godoc.org/labix.org/v2/mgo
- https://docs.mongodb.com/getting-started/shell/query/
- https://github.com/michaeltreat/Mongo_quickstart
