package query

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/soraQaQ/blog/app/user/internal/domain"
	"github.com/soraQaQ/blog/pkg/decorator"
)

type GetAllUser struct{}

type GetAllUserHandler decorator.QueryHandler[GetAllUser, []*domain.User]

type getAllUsersHandler struct {
	log  *log.Helper
	repo domain.UserRepo
}

func NewGetAllUserHandler(logger log.Logger, repo domain.UserRepo) GetAllUserHandler {
	return decorator.ApplyQueryDecorator[GetAllUser, []*domain.User](
		getAllUsersHandler{
			log:  log.NewHelper(logger),
			repo: repo,
		}, logger)
}

func (c getAllUsersHandler) Handler(ctx context.Context, _ GetAllUser) ([]*domain.User, error) {
	users, err := c.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
