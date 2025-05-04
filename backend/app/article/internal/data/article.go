package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/soraQaQ/blog/app/article/internal/domain"
)

type ArticleRepo struct {
	data *Data
	log  *log.Helper
}

func NewArticleRepo(data *Data, logger log.Logger) domain.ArticleRepo {
	return &ArticleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (a *ArticleRepo) Save(ctx context.Context, article *domain.Article) (err error) {
	err = a.data.db.Save(ctx, article)
	if err != nil {
		return
	}
	return
}

func (a *ArticleRepo) Update(ctx context.Context, article *domain.Article, f func(context.Context, *domain.Article) (*domain.Article, error)) error {
	err := a.data.db.Update(ctx, article, f)
	if err != nil {
		return err
	}
	return nil
}

func (a *ArticleRepo) GetAll(ctx context.Context) ([]*domain.Article, error) {
	articles, err := a.data.db.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (a *ArticleRepo) GetArticlesByTag(ctx context.Context, s string) ([]*domain.Article, error) {
	articles, err := a.data.db.GetArticlesByTag(ctx, s)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (a *ArticleRepo) Get(ctx context.Context, i int64) (*domain.Article, error) {
	article, err := a.data.db.Get(ctx, i)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (a *ArticleRepo) Delete(ctx context.Context, i int64) error {
	err := a.data.db.Delete(ctx, i)
	if err != nil {
		return err
	}
	return nil
}
