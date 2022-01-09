package store

import (
	"net/http"

	"git.sr.ht/~will-clarke/news-api/importer"
	"git.sr.ht/~will-clarke/news-api/model"
	"github.com/labstack/echo/v4"
)

// (GET /Feeds)
func (s *Store) GetFeeds(ctx echo.Context) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	feeds := []model.Feed{}

	for _, feed := range s.feeds {
		feeds = append(feeds, feed)
	}
	return ctx.JSON(http.StatusOK, feeds)
}

// (POST /Feeds)
func (s *Store) PostFeed(ctx echo.Context) error {
	var newFeed model.NewFeed
	err := ctx.Bind(&newFeed)
	if err != nil {
		FeedErr := model.Error{
			Code:    int32(http.StatusBadRequest),
			Message: "unable to unmarshal newFeed",
		}
		err := ctx.JSON(http.StatusBadRequest, FeedErr)
		return err
	}

	f := importer.Import(newFeed, s)

	return ctx.JSON(http.StatusCreated, f)
}

func (s *Store) StoreFeed(newFeed model.NewFeed) model.Feed {
	s.lock.Lock()
	defer s.lock.Unlock()

	f := model.Feed{}
	f.Url = newFeed.Url
	f.Categories = newFeed.Categories
	f.Id = s.nextFeedID

	s.nextFeedID += 1

	s.feeds[f.Id] = f
	return f
}
