package data

import (
	"context"
	"github.com/soraQaQ/blog/app/user/internal/domain"

	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) domain.UserRepo {
	return &UserRepo{data: data, log: log.NewHelper(logger)}
}

func (u *UserRepo) GetUserByEmail(ctx context.Context, email string) (user *domain.User, err error) {
	user, err = u.data.db.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepo) Save(ctx context.Context, user *domain.User) (err error) {
	err = u.data.db.Save(ctx, user)
	if err != nil {
		return
	}
	return
}

func (u *UserRepo) Get(ctx context.Context, id uint64) (user *domain.User, err error) {
	user, err = u.data.db.Get(ctx, id)
	if err != nil {
		return
	}
	return
}

func (u *UserRepo) Update(ctx context.Context, user *domain.User, fn func(context.Context, *domain.User) (*domain.User, error)) (err error) {
	err = u.data.db.Update(ctx, user, fn)
	if err != nil {
		return
	}
	return nil
}

func (u *UserRepo) GetAll(ctx context.Context) ([]*domain.User, error) {
	return u.data.db.GetAll(ctx)
}
