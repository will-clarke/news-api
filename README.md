# News Article API (via RSS)

This app is an article aggregator. You can submit your favourite RSS feeds,
categorise them and then query the API for specific articles or articles
matching certain categories or feeds.

This golang webapp is build around an openapi interface. The contract exposed in
[api/openapi.yaml](api/openapi.yaml) is a basic RESTful API.

Here's a quick look at how you can interact with the API

``` sh
## start the server in the background (or other tab)
$ make run &

## add a feed
$ curl localhost:8080/feeds -u top:secret -H 'content-type: application/json' -d '{"url": "http://feeds.bbci.co.uk/news/technology/rss.xml", "categories": ["tech", "super-cool"]}'
{"categories":["tech","super-cool"],"url":"http://feeds.bbci.co.uk/news/technology/rss.xml","id":1}

## list all articles
$ curl localhost:8080/articles -u top:secret 
[{"categories":["tech","super-cool"],"feedId":1,"publishedDate":"Sat, 27 Nov 2021 11:24:06 GMT","title":"What are quantum computers and what are they used for?","url":"https://www.bbc.co.uk/news/technology-59300301?at_medium=RSS\u0026at_campaign=KARANGA","id":14},{"categories":["tech","super-cool"],"feedId":1,"publishedDate":"Fri, 07 Jan 2022 10:56:36 GMT","title":"Bitcoin prices fall to lowest in months after US Fed remarks","url":"https://www.bbc.co.uk/news/techn..}]

## list specific article
$ curl localhost:8080/articles/3 -u top:secret
{"categories":["tech","super-cool"],"feedId":1,"publishedDate":"Fri, 07 Jan 2022 14:05:36 GMT","title":"France fines Google and Facebook over cookies","url":"https://www.bbc.co.uk/news/technology-59909647?at_medium=RSS\u0026at_campaign=KARANGA","id":3}

## list all feeds
$ curl localhost:8080/feeds -u top:secret
[{"categories":["tech","super-cool"],"url":"http://feeds.bbci.co.uk/news/technology/rss.xml","id":1}]

## adding another feed
$ curl localhost:8080/feeds -u top:secret -H 'content-type: application/json' -d '{"url": "http://feeds.skynews.com/feeds/rss/uk.xml", "categories": ["sky"]}'
{"categories":["sky"],"url":"http://feeds.skynews.com/feeds/rss/uk.xml","id":2}

## filtering articles by category (cn accept more than one input)
$ curl localhost:8080/articles\?categories=sky -u top:secret
[{"categories":["sky"],"feedId":2,"publishedDate":"Sun, 09 Jan 2022 17:02:00 +0000","title":"Morrisons to scrap 'use by' dates on milk to help reduce food waste","url":"http://news.sky.com/s

## filtering articles by feedID
$ curl localhost:8080/articles\?feedIDs=2 -u top:secret
[{"categories":["sky"],"feedId":2,"publish..}]

## filtering articles by more than one feedID
$ curl "localhost:8080/articles?feedIDs=1&feedIDs=2" -u top:secret
[{"categories":["tech","super-cool"],"feedId":1,"publ...}]
```

## Is this a pointless app?! :D

- Honestly I reckon most RSS clients should be able to do all this themselves..
and there may not be *that* much value in creating a web service just to read
articles from an RSS feed. We've basically just transcoded the RSS spec into
JSON (albeit with a smaller interface...). I'd question the purpose of this app!

## Limitations

- Error handing throughout could be better
- I'm not logging anything. Or monitoring anything.
- The "database" is just in-memory. Because of this we can't really have more
than one app running.
- I should have tested WAY better.
I would have implemented some mocked-out unit tests for the Stores, using the
Get and Set methods to make sure we were properly persisting the models.
Given all the autogenerated server code too, it would probably be worth having
a suite of integration tests, making sure some end-to-end scenarios work. We'd
want to stub the real HTTP requests by either injecting the HTTP client in
somehow (the autogenerated client code created an `HttpRequestDoer` which would
be how you'd stub that), or by using something like HTTPMock.
- I didn't get round to implementing any HTML-rendered pages; it's just JSON.
- The Go code isn't always pretty. A fair chunk of it is autogenerated. And the
re  st of it has to fit in with the framework oapi-codegen provides.
- I wanted to split up my openapi autogenerated models (& potentially the spec).
But the autogeneration wants us to just use one file, which is okay but a
little restrictive.
- (Continuing the point above...) Autogeneration is very restrictive. There's
  only one way to generate the code, and if requirements change or we need to do
  something totally different with the server, we may have to drop
  autogenerating the server.
  We should be safe to always rely on the autogenerated types though.

## Good stuff

- OpenAPI is super cool and provides you with a great foundation for any API. We
can generate laods of code and documentation from it.
It should be easy to extend this app by adding whole new APIs to the openApi spec.
- Go's a fastish language so we should be able to handle decent throughput.
- We're using a popular web framework which is easily extensible (eg. easy to
chuck middleware at the stack).
- It's super secure. BASIC AUTH works
