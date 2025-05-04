package query

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/soraQaQ/blog/app/article/internal/domain"
	"github.com/soraQaQ/blog/pkg/decorator"
)

type GetAllArticles struct{}

type GetAllArticlesHandler decorator.QueryHandler[GetAllArticles, []*domain.Article]

type getAllArticlesHandler struct {
	repo domain.ArticleRepo
	log  *log.Helper
}

func NewGetAllArticleHandler(repo domain.ArticleRepo, logger log.Logger) GetAllArticlesHandler {
	return decorator.ApplyQueryDecorator[GetAllArticles, []*domain.Article](
		getAllArticlesHandler{
			repo: repo,
			log:  log.NewHelper(log.With(logger, "module", "query")),
		}, logger,
	)
}

func (g getAllArticlesHandler) Handler(ctx context.Context, _ GetAllArticles) ([]*domain.Article, error) {
	articles, err := g.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return articles, nil
}
