package importer

import (
	"git.sr.ht/~will-clarke/news-api/model"
	"github.com/mmcdole/gofeed"
)

type Importer interface {
	StoreArticle(newArticle model.NewArticle) model.Article
	StoreFeed(newFeed model.NewFeed) model.Feed
}

func Import(newFeed model.NewFeed, importer Importer) (model.Feed, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(newFeed.Url)
	if err != nil {
		return model.Feed{}, err
	}

	// TODO: We need to think about error handling scenarios;
	// do we want to add this feed if any of the items fail?
	storedFeed := importer.StoreFeed(newFeed)
	for _, item := range feed.Items {
		newArticle := model.NewArticle{
			Categories:    newFeed.Categories,
			FeedId:        storedFeed.Id,
			PublishedDate: &item.Published,
			Title:         item.Title,
			Url:           item.Link,
		}

		// nothing to do with the stored articles
		_ = importer.StoreArticle(newArticle)
	}

	return storedFeed, nil
}
