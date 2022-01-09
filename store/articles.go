package store

import (
	"fmt"
	"net/http"

	"git.sr.ht/~will-clarke/news-api/model"
	"github.com/labstack/echo/v4"
)

// (GET /articles)
// This one is super inefficient and should probably be destroyed (or at least properly indexed)
func (s *Store) GetArticles(ctx echo.Context, params model.GetArticlesParams) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	articles := []model.Article{}

	if params.Categories == nil && params.FeedIDs == nil {
		// no filtering so I'm guessing we should just return EVERYTHING now, performance be damned.
		// We can add pagination later if we want.
		for _, art := range s.articles {
			articles = append(articles, art)
		}
		return ctx.JSON(http.StatusOK, articles)
	}

	feedIDParams := []int64{}
	categoryParams := []string{}
	if params.FeedIDs != nil {
		feedIDParams = *params.FeedIDs
	}
	if params.Categories != nil {
		categoryParams = *params.Categories
	}

	// starting a horrific for loop.. hang on!!!
NEXTARTICLE:
	for _, art := range s.articles {
		for _, feedParam := range feedIDParams {
			if feedParam == art.FeedId {
				articles = append(articles, art)
				continue NEXTARTICLE
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
func (s *Store) PostArticle(ctx echo.Context) error {
	var newArticle model.NewArticle
	err := ctx.Bind(&newArticle)
	if err != nil {
		return sendArticleStoreError(ctx, http.StatusBadRequest, "Invalid format for NewPet")
	}

	art := s.StoreArticle(newArticle)

	return ctx.JSON(http.StatusCreated, art)
}

func (s *Store) StoreArticle(newArticle model.NewArticle) model.Article {
	s.lock.Lock()
	defer s.lock.Unlock()

	art := model.Article{}
	art.Categories = newArticle.Categories
	art.FeedId = newArticle.FeedId
	art.PublishedDate = newArticle.PublishedDate
	art.Url = newArticle.Url
	art.Title = newArticle.Title
	art.Id = s.nextArticleID

	s.nextArticleID += 1

	s.articles[art.Id] = art
	return art
}

// (DELETE /articles/{id})
func (s *Store) DeleteArticle(ctx echo.Context, id int64) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.articles, id)

	_, found := s.articles[id]
	if !found {
		return sendArticleStoreError(ctx, http.StatusNotFound,
			fmt.Sprintf("unable to find article #%d", id))
	}

	return ctx.NoContent(http.StatusNoContent)
}

// (GET /articles/{id})
func (s *Store) GetArticle(ctx echo.Context, id int64) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	article, found := s.articles[id]
	if !found {
		return sendArticleStoreError(ctx, http.StatusNotFound,
			fmt.Sprintf("unable to find article #%d", id))
	}

	return ctx.JSON(http.StatusOK, article)
}

// This function wraps sending of an error in the Error format, and
// handling the failure to marshal that.
func sendArticleStoreError(ctx echo.Context, code int, message string) error {
	articleErr := model.Error{
		Code:    int32(code),
		Message: message,
	}
	err := ctx.JSON(code, articleErr)
	return err
}
