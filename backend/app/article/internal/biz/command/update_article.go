package command

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/soraQaQ/blog/app/article/internal/domain"
	"github.com/soraQaQ/blog/pkg/decorator"
)

type UpdateArticle struct {
	Article  *domain.Article
	UpdateFn func(context.Context, *domain.Article) (*domain.Article, error)
}

type UpdateArticleHandler decorator.CommandHandler[UpdateArticle]

type updateArticleHandler struct {
	log  *log.Helper
	repo domain.ArticleRepo
}

func NewUpdateArticleHandler(logger log.Logger, repo domain.ArticleRepo) UpdateArticleHandler {
	return decorator.ApplyCommandDecorators[UpdateArticle](
		updateArticleHandler{
			log:  log.NewHelper(log.With(logger, "module", "command")),
			repo: repo,
		}, logger)
}

func (g updateArticleHandler) Handler(ctx context.Context, cmd UpdateArticle) error {
	err := g.repo.Update(ctx, cmd.Article, cmd.UpdateFn)
	if err != nil {
		return err
	}
	return nil
}
