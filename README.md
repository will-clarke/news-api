# News Article API (via RSS)

Here's an example API spec that I'll implement towards

## Endpoints

### GET /articles
[{id, body, categories, feed, publishdate}]

#### Filtering:
### GET /articles?feeds=:id_1,:id_2
### GET /articles?categories=:category_1,:category_2

### GET /articles/:id
JSON: {id, body, categories, feed, publishdate}
HTML: ? Just the raw html page ?
Or a template would actually be cool



### GET /feeds
[{URL, categories}]

### POST /feeds
{URL, categories}

## Store interface

``` go
type Feed struct {
feedURL string
categories []string
}


type FeedStore interface {
StoreFeed(feed Feed) (err error, id int)
GetFeeds() (err error, []Feed)
}

type ArticleStore interface {
GetArticles() (err error, []Article)
GetArticlesByCategory(category string) (err error, []Article)
GetArticlesByFeed(feed int) (err error, []Article)
GetArticlesByID(id int) (err error, Article)
SetArticle(a Article) error
}
```


## Other stuff to think about
- Pagination

- Example feeds:
http://feeds.bbci.co.uk/news/uk/rss.xml
http://feeds.bbci.co.uk/news/technology/rss.xml
http://feeds.skynews.com/feeds/rss/uk.xml
http://feeds.skynews.com/feeds/rss/technology.xml
