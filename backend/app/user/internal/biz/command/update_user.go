package command

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/soraQaQ/blog/app/user/internal/domain"
	"github.com/soraQaQ/blog/pkg/decorator"
)

type UpdateUser struct {
	User     *domain.User
	UpdateFn func(ctx context.Context, user *domain.User) (*domain.User, error)
}

type UpdateUserHandler decorator.CommandHandler[UpdateUser]

type updateUserHandler struct {
	log  *log.Helper
	repo domain.UserRepo
}

func NewUpdateUserHandler(logger log.Logger, repo domain.UserRepo) UpdateUserHandler {
	return decorator.ApplyCommandDecorators[UpdateUser](updateUserHandler{
		log:  log.NewHelper(logger),
		repo: repo,
	}, logger)
}

func (u updateUserHandler) Handler(ctx context.Context, cmd UpdateUser) error {
	err := u.repo.Update(ctx, cmd.User, cmd.UpdateFn)
	if err != nil {
		return err
	}
	return nil
}
