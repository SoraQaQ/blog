package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id       uint64
	Username string
	Nickname string
	Password string
	Email    string
}

type UserRepo interface {
	Save(context.Context, *User) error
	Get(context.Context, uint64) (*User, error)
	Update(context.Context, *User) error
	GetAll(context.Context) ([]*User, error)
	GetUserByEmail(context.Context, string) (*User, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/user"))}
}

func (u *UserUseCase) Save(ctx context.Context, user *User) error {
	return u.repo.Save(ctx, user)
}

func (u *UserUseCase) Get(ctx context.Context, id uint64) (*User, error) {
	return u.repo.Get(ctx, id)
}

func (u *UserUseCase) Update(ctx context.Context, user *User) error {
	return u.repo.Update(ctx, user)
}

func (u *UserUseCase) GetAll(ctx context.Context) ([]*User, error) {
	return u.repo.GetAll(ctx)
}

func (u *UserUseCase) GetUserByEmail(ctx context.Context, s string) (*User, error) {
	return u.repo.GetUserByEmail(ctx, s)
}
