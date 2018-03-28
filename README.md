# Newsy App

This is a RESTful Go API that returns top news articles from several tech news sources.

It supports two endpoints currently:
1. https://newsy-bits.herokuapp.com/crawl -- Crawls news sources and saves top news stories in MongoDB database
2. https://newsy-bits.herokuapp.com/news -- Gets all top news stories saved in the MongoDB database and returns as JSON

See it in action: http://www.marydickson.info/news

In the future, I would add a job or recurring task to regularly crawl for new articles (currently only works on-demand), and figure out a way to delete old news from the database and/or only return a certain number of articles.

TODO:
- limit number of responses for get news endpoint
- add cacheing, and/or some way of checking when APIs were last crawled (add to the JSON?)
- add a job that crawls the APIs on a recurring schedule (2x per day?) and removes old entries from DB
- better error handling and logging
- add tests

Questions for Consideration:
- do we want an even mix of the news sources, maybe the latest 10 for each? if yes, return the appropriate json
- how often to crawl for new news and remove old news?

Resources:
- https://www.mongodb.com/blog/post/building-your-first-application-mongodb-creating-rest-api-using-mean-stack-part-1
- http://www.blog.labouardy.com/build-restful-api-in-go-and-mongodb/
- https://docs.mongodb.com/getting-started/shell/query/
- https://godoc.org/labix.org/v2/mgo
- https://medium.com/@IndianGuru/go-mongodb-mongolab-mgo-and-heroku-d411b5ac53f9
- https://github.com/kardianos/govendor (required for Heroku)
- https://github.com/michaeltreat/Mongo_quickstart
