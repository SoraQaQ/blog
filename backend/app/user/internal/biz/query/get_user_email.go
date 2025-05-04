package query

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/soraQaQ/blog/app/user/internal/domain"
	"github.com/soraQaQ/blog/pkg/decorator"
)

type GetUserByEmail struct {
	Email string
}

type GetUsersByEmailHandler decorator.QueryHandler[GetUserByEmail, *domain.User]

type getUserByEmailHandler struct {
	log  *log.Helper
	repo domain.UserRepo
}

func NewGetUserByEmailHandler(logger log.Logger, repo domain.UserRepo) GetUsersByEmailHandler {
	return decorator.ApplyQueryDecorator[GetUserByEmail, *domain.User](
		getUserByEmailHandler{
			log:  log.NewHelper(logger),
			repo: repo,
		}, logger)
}

func (g getUserByEmailHandler) Handler(ctx context.Context, query GetUserByEmail) (*domain.User, error) {
	user, err := g.repo.GetUserByEmail(ctx, query.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
