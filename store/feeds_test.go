package store_test

import (
	"testing"

	"git.sr.ht/~will-clarke/news-api/model"
	"git.sr.ht/~will-clarke/news-api/store"
	"github.com/stretchr/testify/assert"
)

var exampleTitle = "an example title"

func TestStore_StoreFeed(t *testing.T) {
	tests := []struct {
		name    string
		newFeed model.NewFeed
		want    model.Feed
	}{
		{
			name:    "first feed has Id 1",
			newFeed: model.NewFeed{Url: "https://example.com"},
			want: model.Feed{FeedAllOf: model.FeedAllOf{Id: 1},
				NewFeed: model.NewFeed{Url: "https://example.com"}},
		},
		{
			name:    "second feed has Id 2",
			newFeed: model.NewFeed{Url: "https://example.com"},
			want: model.Feed{FeedAllOf: model.FeedAllOf{Id: 2},
				NewFeed: model.NewFeed{Url: "https://example.com"}},
		},
		{
			name:    "thrid feed has Id 3 and persists the title",
			newFeed: model.NewFeed{Url: "https://example.com", Title: &exampleTitle},
			want: model.Feed{FeedAllOf: model.FeedAllOf{Id: 3},
				NewFeed: model.NewFeed{Url: "https://example.com", Title: &exampleTitle}},
		},
	}
	s := store.NewStore()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storedFeed := s.StoreFeed(tt.newFeed)

			assert.Equal(t, tt.want, storedFeed)
		})
	}
}
