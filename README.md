TODO:
- limit number of responses for get news endpoint
- add cacheing, and/or some way of checking when APIs were last crawled (add to the JSON?)
- add a job that crawls the APIs on a recurring schedule (2x per day?) and removes old entries from DB
- better error handling and logging
- add tests
- deploy!

Questions for Consideration:
- do we want an even mix of the news sources, maybe the latest 10 for each? if yes, return the appropriate json
- how often to crawl for new news and remove old news?

Resources:
- https://www.mongodb.com/blog/post/building-your-first-application-mongodb-creating-rest-api-using-mean-stack-part-1
- http://www.blog.labouardy.com/build-restful-api-in-go-and-mongodb/
- https://docs.mongodb.com/getting-started/shell/query/
- https://medium.com/@IndianGuru/go-mongodb-mongolab-mgo-and-heroku-d411b5ac53f9