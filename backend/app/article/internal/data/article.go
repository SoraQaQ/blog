package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/soraQaQ/blog/app/article/internal/biz"
)

type ArticleRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &ArticleRepo{data: data, log: log.NewHelper(logger)}
}

func (a *ArticleRepo) Save(ctx context.Context, article *biz.Article) (*biz.Article, error) {
	newArticle, err := a.data.db.Save(ctx, article)
	if err != nil {
		return nil, err
	}
	return newArticle, nil
}

func (a *ArticleRepo) Update(ctx context.Context, article *biz.Article, f func(context.Context, *biz.Article) (*biz.Article, error)) error {
	err := a.data.db.Update(ctx, article, f)
	if err != nil {
		return err
	}
	return nil
}

func (a *ArticleRepo) GetAll(ctx context.Context) ([]*biz.Article, error) {
	articles, err := a.data.db.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (a *ArticleRepo) GetArticlesByTag(ctx context.Context, s string) ([]*biz.Article, error) {
	articles, err := a.data.db.GetArticlesByTag(ctx, s)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (a *ArticleRepo) Get(ctx context.Context, i int64) (*biz.Article, error) {
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
