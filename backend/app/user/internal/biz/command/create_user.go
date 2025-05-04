package command

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/soraQaQ/blog/app/user/internal/domain"
	"github.com/soraQaQ/blog/pkg/decorator"
)

type CreateUser struct {
	User *domain.User
}

type CreateUserHandler decorator.CommandHandler[CreateUser]

type createUserHandler struct {
	log  *log.Helper
	repo domain.UserRepo
}

func NewCreateUserHandler(logger log.Logger, repo domain.UserRepo) CreateUserHandler {
	return decorator.ApplyCommandDecorators[CreateUser](createUserHandler{
		log:  log.NewHelper(logger),
		repo: repo,
	}, logger)
}

func (c createUserHandler) Handler(ctx context.Context, cmd CreateUser) error {
	err := c.repo.Save(ctx, cmd.User)
	if err != nil {
		return err
	}
	return nil
}
