package main

import (
	"reflect"
	"sync"
	"testing"

	"git.sr.ht/~will-clarke/news-api/article"
	"github.com/labstack/echo/v4"
)

func TestNewArticleStore(t *testing.T) {
	tests := []struct {
		name string
		want *ArticleStore
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArticleStore(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArticleStore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArticleStore_GetArticles(t *testing.T) {
	type fields struct {
		articles   map[int64]article.Article
		nextNumber int64
		lock       sync.Mutex
	}
	type args struct {
		ctx    echo.Context
		params article.GetArticlesParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArticleStore{
				articles:   tt.fields.articles,
				nextNumber: tt.fields.nextNumber,
				lock:       tt.fields.lock,
			}
			if err := a.GetArticles(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("ArticleStore.GetArticles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestArticleStore_PostArticle(t *testing.T) {
	type fields struct {
		articles   map[int64]article.Article
		nextNumber int64
		lock       sync.Mutex
	}
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArticleStore{
				articles:   tt.fields.articles,
				nextNumber: tt.fields.nextNumber,
				lock:       tt.fields.lock,
			}
			if err := a.PostArticle(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("ArticleStore.PostArticle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestArticleStore_DeleteArticle(t *testing.T) {
	type fields struct {
		articles   map[int64]article.Article
		nextNumber int64
		lock       sync.Mutex
	}
	type args struct {
		ctx echo.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArticleStore{
				articles:   tt.fields.articles,
				nextNumber: tt.fields.nextNumber,
				lock:       tt.fields.lock,
			}
			if err := a.DeleteArticle(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("ArticleStore.DeleteArticle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestArticleStore_GetArticle(t *testing.T) {
	type fields struct {
		articles   map[int64]article.Article
		nextNumber int64
		lock       sync.Mutex
	}
	type args struct {
		ctx echo.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArticleStore{
				articles:   tt.fields.articles,
				nextNumber: tt.fields.nextNumber,
				lock:       tt.fields.lock,
			}
			if err := a.GetArticle(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("ArticleStore.GetArticle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_sendArticleStoreError(t *testing.T) {
	type args struct {
		ctx     echo.Context
		code    int
		message string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := sendArticleStoreError(tt.args.ctx, tt.args.code, tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("sendArticleStoreError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
