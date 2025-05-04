package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
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

type ArticleRepo interface {
	Save(context.Context, *Article) error
	Get(context.Context, int64) (*Article, error)
	Update(context.Context, *Article) error
	GetAll(context.Context) ([]*Article, error)
	GetArticlesByTag(context.Context, string) ([]*Article, error)
	Delete(context.Context, int64) error
}

type ArticleUsecase struct {
	repo ArticleRepo
	log  *log.Helper
}

func NewArticleUsecase(repo ArticleRepo, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/article"))}
}

func (uc *ArticleUsecase) Save(ctx context.Context, article *Article) error {
	err := uc.repo.Save(ctx, article)
	if err != nil {
		return err
	}
	return nil
}

func (uc *ArticleUsecase) Get(ctx context.Context, id int64) (*Article, error) {
	article, err := uc.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (uc *ArticleUsecase) Update(ctx context.Context, article *Article) error {
	err := uc.repo.Update(ctx, article)
	if err != nil {
		return err
	}
	return nil
}

func (uc *ArticleUsecase) GetAll(ctx context.Context) ([]*Article, error) {
	articles, err := uc.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (uc *ArticleUsecase) Delete(ctx context.Context, id int64) error {
	err := uc.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (uc *ArticleUsecase) GetArticlesByTag(ctx context.Context, tags string) ([]*Article, error) {
	articles, err := uc.repo.GetArticlesByTag(ctx, tags)
	if err != nil {
		return nil, err
	}
	return articles, nil
}
