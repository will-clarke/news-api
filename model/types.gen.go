// Package model provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package model

// Article defines model for Article.
type Article struct {
	// Embedded struct due to allOf(#/components/schemas/NewArticle)
	NewArticle `yaml:",inline"`
	// Embedded struct due to allOf(#/components/schemas/Article_allOf)
	ArticleAllOf `yaml:",inline"`
}

// ArticleAllOf defines model for Article_allOf.
type ArticleAllOf struct {
	Id int64 `json:"id"`
}

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

// NewArticle defines model for NewArticle.
type NewArticle struct {
	Categories    *[]string `json:"categories,omitempty"`
	Feed          *string   `json:"feed,omitempty"`
	PublishedDate *string   `json:"publishedDate,omitempty"`
	Url           string    `json:"url"`
}

// NewFeed defines model for NewFeed.
type NewFeed struct {
	Categories *[]string `json:"categories,omitempty"`
	Url        string    `json:"url"`
}

// GetArticlesParams defines parameters for GetArticles.
type GetArticlesParams struct {
	// feeds to filter by
	Feeds *[]string `json:"feeds,omitempty"`

	// categories to filter by
	Categories *[]string `json:"categories,omitempty"`
}

// PostArticleJSONBody defines parameters for PostArticle.
type PostArticleJSONBody NewArticle

// PostFeedJSONBody defines parameters for PostFeed.
type PostFeedJSONBody NewFeed

// PostArticleJSONRequestBody defines body for PostArticle for application/json ContentType.
type PostArticleJSONRequestBody PostArticleJSONBody

// PostFeedJSONRequestBody defines body for PostFeed for application/json ContentType.
type PostFeedJSONRequestBody PostFeedJSONBody
