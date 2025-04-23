package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
)

type Article struct {
	Id         int64
	Title      string
	Summary    string
	ContentUrl string
	status     int64
	ViewCount  int64
	Tags       string
	ImageUrl   []string
}

type ArticleRepo interface {
	Save(context.Context, *Article) (*Article, error)
	Get(context.Context, int64) (*Article, error)
	Update(context.Context, *Article, func(context.Context, *Article) (*Article, error)) error
	GetAll(context.Context) ([]*Article, error)
	GetArticlesByTag(context.Context, string) ([]*Article, error)
	Delete(context.Context, int64) error
}

type ArticleUsecase struct {
	repo ArticleRepo
	log  *log.Helper
}

func NewArticleUsecase(repo ArticleRepo, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ArticleUsecase) GetAllArticles(ctx context.Context) ([]*Article, error) {
	articles, err := uc.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (uc *ArticleUsecase) GetArticlesByTag(ctx context.Context, tag string) ([]*Article, error) {
	if tag == "" {
		return nil, fmt.Errorf("empty tag")
	}
	articles, err := uc.repo.GetArticlesByTag(ctx, tag)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (uc *ArticleUsecase) Get(ctx context.Context, id int64) (*Article, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid id")
	}
	article, err := uc.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (uc *ArticleUsecase) DeleteArticle(ctx context.Context, id int64) error {
	if id <= 0 {
		return fmt.Errorf("invalid id")
	}

	article, err := uc.repo.Get(ctx, id)
	if err != nil {
		return err
	}

	if article == nil {
		return fmt.Errorf("article not found")
	}

	err = uc.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (uc *ArticleUsecase) CreateArticle(ctx context.Context, article *Article) (*Article, error) {
	if err := validateArticle(article); err != nil {
		return nil, err
	}
	newArticle, err := uc.repo.Save(ctx, article)
	if err != nil {
		return nil, err
	}
	return newArticle, nil
}

func validateArticle(article *Article) error {
	if article == nil {
		return fmt.Errorf("article is nil")
	}
	if article.Id <= 0 {
		return fmt.Errorf("invalid id, id must be greater than zero")
	}
	if article.Title == "" {
		return fmt.Errorf("invalid title, title is empty")
	}
	if article.Summary == "" {
		return fmt.Errorf("invalid summary, summary is empty")
	}
	if article.ContentUrl == "" {
		return fmt.Errorf("invalid content_url, content_url is empty")
	}
	return nil
}
