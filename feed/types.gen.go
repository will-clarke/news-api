// Package feed provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package feed

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// Feed defines model for Feed.
type Feed struct {
	// Embedded struct due to allOf(#/components/schemas/NewFeed)
	NewFeed `yaml:",inline"`
	// Embedded struct due to allOf(#/components/schemas/Feed_allOf)
	FeedAllOf `yaml:",inline"`
}

// FeedAllOf defines model for Feed_allOf.
type FeedAllOf struct {
	Id int64 `json:"id"`
}

// NewFeed defines model for NewFeed.
type NewFeed struct {
	Categories *[]string `json:"categories,omitempty"`
	Url        string    `json:"url"`
}

// PostFeedJSONBody defines parameters for PostFeed.
type PostFeedJSONBody NewFeed

// PostFeedJSONRequestBody defines body for PostFeed for application/json ContentType.
type PostFeedJSONRequestBody PostFeedJSONBody

