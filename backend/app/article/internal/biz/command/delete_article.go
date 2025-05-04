package command

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/soraQaQ/blog/app/article/internal/domain"
	"github.com/soraQaQ/blog/pkg/decorator"
)

type DeleteArticle struct {
	Id int64
}

type DeleteArticleHandler decorator.CommandHandler[DeleteArticle]

type deleteArticleHandler struct {
	repo domain.ArticleRepo
	log  *log.Helper
}

func NewDeleteArticleHandler(repo domain.ArticleRepo, logger log.Logger) DeleteArticleHandler {
	return decorator.CommandHandler[DeleteArticle](deleteArticleHandler{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "command")),
	})
}

func (d deleteArticleHandler) Handler(ctx context.Context, cmd DeleteArticle) error {
	err := d.repo.Delete(ctx, cmd.Id)
	if err != nil {
		return err
	}
	return nil
}
