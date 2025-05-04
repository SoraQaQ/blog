package query

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/soraQaQ/blog/app/article/internal/domain"
	"github.com/soraQaQ/blog/pkg/decorator"
)

type GetArticle struct {
	Id int64
}

type GetArticleHandler decorator.QueryHandler[GetArticle, *domain.Article]

type getArticleHandler struct {
	repo domain.ArticleRepo
	log  *log.Helper
}

func NewGetArticleHandler(repo domain.ArticleRepo, logger log.Logger) GetArticleHandler {
	return decorator.ApplyQueryDecorator[GetArticle, *domain.Article](
		getArticleHandler{
			repo: repo,
			log:  log.NewHelper(log.With(logger, "module", "query")),
		}, logger,
	)
}

func (g getArticleHandler) Handler(ctx context.Context, query GetArticle) (*domain.Article, error) {
	article, err := g.repo.Get(ctx, query.Id)
	if err != nil {
		return nil, err
	}
	return article, nil
}
