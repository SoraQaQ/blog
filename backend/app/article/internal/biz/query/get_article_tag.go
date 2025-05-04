package query

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/soraQaQ/blog/app/article/internal/domain"
	"github.com/soraQaQ/blog/pkg/decorator"
)

type GetArticlesByTag struct {
	Tag string
}

type GetArticlesByTagHandler decorator.QueryHandler[GetArticlesByTag, []*domain.Article]

type getArticlesByTagHandler struct {
	repo domain.ArticleRepo
	log  *log.Helper
}

func NewGetArticlesByTagHandler(repo domain.ArticleRepo, logger log.Logger) GetArticlesByTagHandler {
	return decorator.ApplyQueryDecorator[GetArticlesByTag, []*domain.Article](
		getArticlesByTagHandler{
			repo: repo,
			log:  log.NewHelper(log.With(logger, "module", "query")),
		}, logger,
	)
}

func (g getArticlesByTagHandler) Handler(ctx context.Context, query GetArticlesByTag) ([]*domain.Article, error) {
	articles, err := g.repo.GetArticlesByTag(ctx, query.Tag)
	if err != nil {
		return nil, err
	}
	return articles, nil
}
