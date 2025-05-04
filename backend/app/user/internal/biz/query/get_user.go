package query

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/soraQaQ/blog/app/user/internal/domain"
	"github.com/soraQaQ/blog/pkg/decorator"
)

type GetUser struct {
	Id uint64
}

type GetUserHandler decorator.QueryHandler[GetUser, *domain.User]

type getUserHandler struct {
	log  *log.Helper
	repo domain.UserRepo
}

func NewGetUserHandler(logger log.Logger, repo domain.UserRepo) GetUserHandler {
	return decorator.ApplyQueryDecorator[GetUser, *domain.User](getUserHandler{
		log:  log.NewHelper(logger),
		repo: repo,
	}, logger)
}

func (g getUserHandler) Handler(ctx context.Context, query GetUser) (*domain.User, error) {
	user, err := g.repo.Get(ctx, query.Id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
