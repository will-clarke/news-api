package store

import (
	"net/http"
	"sync"

	"git.sr.ht/~will-clarke/news-api/feed"
	"github.com/labstack/echo/v4"
)

type FeedStore struct {
	feeds      map[int64]feed.Feed
	nextNumber int64
	lock       sync.Mutex
}

func NewFeedStore() *FeedStore {
	return &FeedStore{
		feeds:      make(map[int64]feed.Feed),
		nextNumber: 1,
	}
}

// (GET /Feeds)
func (s *FeedStore) GetFeeds(ctx echo.Context) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	feeds := []feed.Feed{}

	for _, feed := range s.feeds {
		feeds = append(feeds, feed)
	}
	return ctx.JSON(http.StatusOK, feeds)
}

// (POST /Feeds)
func (s *FeedStore) PostFeed(ctx echo.Context) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	var newFeed feed.NewFeed
	err := ctx.Bind(&newFeed)
	if err != nil {
		FeedErr := feed.Error{
			Code:    int32(http.StatusBadRequest),
			Message: "unable to unmarshal newFeed",
		}
		err := ctx.JSON(http.StatusBadRequest, FeedErr)
		return err
	}

	f := feed.Feed{}
	f.Url = newFeed.Url
	f.Categories = newFeed.Categories
	f.Id = s.nextNumber

	s.nextNumber += 1

	s.feeds[f.Id] = f

	return ctx.JSON(http.StatusCreated, f)
}
