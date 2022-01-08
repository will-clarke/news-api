package store

import (
	"sync"

	"git.sr.ht/~will-clarke/news-api/model"
)

// Store implements the ServerInterface (model/server.gen.go)
// This specific implementation is really very bad, ineficcient
// and won't scale properly (eg. it'll break if not load-balanced properly)
// ... but maybe it's okay for a demo.

type Store struct {
	articles      map[int64]model.Article
	feeds         map[int64]model.Feed
	nextArticleID int64
	nextFeedID    int64
	lock          sync.Mutex
}

func NewStore() *Store {
	return &Store{
		articles:      make(map[int64]model.Article),
		feeds:         make(map[int64]model.Feed),
		nextArticleID: 1,
		nextFeedID:    1,
	}
}
