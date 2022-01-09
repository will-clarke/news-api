package importer

import (
	"fmt"

	"git.sr.ht/~will-clarke/news-api/model"
	"github.com/mmcdole/gofeed"
)

type Importer interface {
	StoreArticle(newArticle model.NewArticle) model.Article
	StoreFeed(newFeed model.NewFeed) model.Feed
}

func Import(newFeed model.NewFeed, importer Importer) model.Feed {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(newFeed.Url)

	// TODO: We need to think about error handling scenarios;
	// do we want to add this feed if any of the items fail?
	storedFeed := importer.StoreFeed(newFeed)
	for _, item := range feed.Items {
		fmt.Printf("item.Title: %+v\n", item.Title)
		fmt.Printf("item.Link: %+v\n", item.Link)
		fmt.Printf("item.Published: %+v\n", item.Published)

		newArticle := model.NewArticle{
			Categories:    newFeed.Categories,
			FeedId:        storedFeed.Id,
			PublishedDate: &item.Published,
			Url:           item.Link,
		}

		// nothing to do with the stored articles
		_ = importer.StoreArticle(newArticle)
	}

	return storedFeed
}
