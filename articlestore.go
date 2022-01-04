package main

import (
	"git.sr.ht/~will-clarke/news-api/article"
	"github.com/labstack/echo/v4"
)

type ArticleStore struct {
}

func NewArticleStore() *ArticleStore {
	return &ArticleStore{}
}

// (GET /articles)
func (a *ArticleStore) GetArticles(ctx echo.Context, params article.GetArticlesParams) error {
	return nil
}

// (POST /articles)
func (a *ArticleStore) PostArticle(ctx echo.Context) error {
	return nil
}

// (DELETE /articles/{id})
func (a *ArticleStore) DeleteArticle(ctx echo.Context, id int64) error {
	return nil
}

// (GET /articles/{id})
func (a *ArticleStore) GetArticle(ctx echo.Context, id int64) error {
	return nil
}
