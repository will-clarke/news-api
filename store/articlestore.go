package store

import (
	"fmt"
	"net/http"
	"sync"

	"git.sr.ht/~will-clarke/news-api/article"
	"github.com/labstack/echo/v4"
)

// ArticleStore implements the ServerInterface (article/server.gen.go)
// This specific implementation is really very bad, ineficcient
// and won't scale properly (it'll break if not load-balanced properly)
// ... but maybe it's okay for a demo.
type ArticleStore struct {
	articles   map[int64]article.Article
	nextNumber int64
	lock       sync.Mutex
}

func NewArticleStore() *ArticleStore {
	return &ArticleStore{
		articles:   make(map[int64]article.Article),
		nextNumber: 1,
	}
}

// (GET /articles)
// This one is super inefficient and should probably be destroyed (or at least properly indexed)
func (a *ArticleStore) GetArticles(ctx echo.Context, params article.GetArticlesParams) error {
	a.lock.Lock()
	defer a.lock.Unlock()

	articles := []article.Article{}

	if params.Categories == nil && params.Feeds == nil {
		// no filtering so I'm guessing we should just return EVERYTHING now, performance be damned.
		// We can add pagination later if we want.
		for _, art := range a.articles {
			articles = append(articles, art)
		}
		return ctx.JSON(http.StatusOK, articles)
	}

	feedParams := []string{}
	categoryParams := []string{}
	if params.Feeds != nil {
		feedParams = *params.Feeds
	}
	if params.Categories != nil {
		categoryParams = *params.Categories
	}

	// starting a horrific for loop.. hang on!!!
NEXTARTICLE:
	for _, art := range a.articles {
		if art.Feed != nil {
			for _, feedParam := range feedParams {
				if feedParam == *art.Feed {
					articles = append(articles, art)
					continue NEXTARTICLE
				}
			}
		}
		if categoryParams == nil || art.Categories == nil {
			continue
		}
		for _, categoryParam := range categoryParams {
			for _, category := range *art.Categories {
				if category == categoryParam {
					articles = append(articles, art)
					continue NEXTARTICLE
				}
			}
		}
	}

	return ctx.JSON(http.StatusOK, articles)
}

// (POST /articles)
func (a *ArticleStore) PostArticle(ctx echo.Context) error {
	a.lock.Lock()
	defer a.lock.Unlock()

	var newArticle article.NewArticle
	err := ctx.Bind(&newArticle)
	if err != nil {
		return sendArticleStoreError(ctx, http.StatusBadRequest, "Invalid format for NewPet")
	}

	art := article.Article{}
	art.Categories = newArticle.Categories
	art.Feed = newArticle.Feed
	art.PublishedDate = newArticle.PublishedDate
	art.Url = newArticle.Url
	art.Id = a.nextNumber

	a.nextNumber += 1

	a.articles[art.Id] = art

	return ctx.JSON(http.StatusCreated, art)
}

// (DELETE /articles/{id})
func (a *ArticleStore) DeleteArticle(ctx echo.Context, id int64) error {
	a.lock.Lock()
	defer a.lock.Unlock()
	delete(a.articles, id)

	_, found := a.articles[id]
	if !found {
		return sendArticleStoreError(ctx, http.StatusNotFound,
			fmt.Sprintf("unable to find article #%d", id))
	}

	return ctx.NoContent(http.StatusNoContent)
}

// (GET /articles/{id})
func (a *ArticleStore) GetArticle(ctx echo.Context, id int64) error {
	a.lock.Lock()
	defer a.lock.Unlock()

	article, found := a.articles[id]
	if !found {
		return sendArticleStoreError(ctx, http.StatusNotFound,
			fmt.Sprintf("unable to find article #%d", id))
	}

	return ctx.JSON(http.StatusOK, article)
}

// This function wraps sending of an error in the Error format, and
// handling the failure to marshal that.
func sendArticleStoreError(ctx echo.Context, code int, message string) error {
	articleErr := article.Error{
		Code:    int32(code),
		Message: message,
	}
	err := ctx.JSON(code, articleErr)
	return err
}
