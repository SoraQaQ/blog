package command

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/soraQaQ/blog/app/article/internal/domain"
	"github.com/soraQaQ/blog/pkg/decorator"
)

type CreateArticle struct {
	Article *domain.Article
}

type CreateArticleHandler decorator.CommandHandler[CreateArticle]

type createArticleHandler struct {
	repo domain.ArticleRepo
	log  *log.Helper
}

func NewCreateOrderHandler(repo domain.ArticleRepo, logger log.Logger) CreateArticleHandler {
	return decorator.ApplyCommandDecorators[CreateArticle](
		createArticleHandler{
			repo: repo,
			log:  log.NewHelper(log.With(logger, "module", "command")),
		},
		logger,
	)
}

func (c createArticleHandler) Handler(ctx context.Context, cmd CreateArticle) (err error) {
	err = c.repo.Save(ctx, cmd.Article)
	if err != nil {
		return
	}
	return nil
}
