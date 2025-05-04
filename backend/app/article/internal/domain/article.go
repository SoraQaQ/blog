package domain

import (
	"context"
	"github.com/soraQaQ/blog/pkg/entity"
	"github.com/soraQaQ/blog/pkg/errors"
)

type Article struct {
	Id         int64
	Title      string
	Summary    string
	ContentUrl string
	Status     int64
	ViewCount  int64
	Tags       string
	ImageUrl   []string
}

func validateArticle(article *Article) error {
	if article == nil {
		return errors.ErrorArticleNil
	}
	if article.Id <= 0 {
		return errors.ErrorArticleID
	}
	if article.Title == "" {
		return errors.ErrorArticleEmptyTitle
	}
	if article.Summary == "" {
		return errors.ErrorArticleEmptySummary
	}
	if article.ContentUrl == "" {
		return errors.ErrorArticleEmptyContent
	}
	return nil
}

type ArticleRepo interface {
	Save(context.Context, *Article) error
	Get(context.Context, int64) (*Article, error)
	Update(context.Context, *Article, func(context.Context, *Article) (*Article, error)) error
	GetAll(context.Context) ([]*Article, error)
	GetArticlesByTag(context.Context, string) ([]*Article, error)
	Delete(context.Context, int64) error
}

func NewArticle(article *entity.Article) *Article {
	return &Article{
		Id:         article.Id,
		Title:      article.Title,
		Summary:    article.Summary,
		ContentUrl: article.ContentUrl,
		Status:     article.Status,
		ViewCount:  article.ViewCount,
		Tags:       article.Tags,
		ImageUrl:   article.ImageUrl,
	}
}

func (a Article) ToEntity(article *Article) *entity.Article {
	return &entity.Article{
		Id:         a.Id,
		Title:      a.Title,
		Summary:    a.Summary,
		ContentUrl: a.ContentUrl,
		Status:     a.Status,
		ViewCount:  a.ViewCount,
		Tags:       a.Tags,
		ImageUrl:   a.ImageUrl,
	}
}
