package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/soraQaQ/blog/app/user/internal/errors"
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
	Update(context.Context, *User, func(context.Context, *User) (*User, error)) error
	GetAll(context.Context) ([]*User, error)
	GetUserByEmail(context.Context, string) (*User, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) CreateUser(ctx context.Context, user *User) (err error) {
	if err = validateUser(user); err != nil {
		uc.log.WithContext(ctx).Errorf("userUsecase.CreateUser: %v", err)
		return
	}
	err = uc.repo.Save(ctx, user)
	if err != nil {
		return
	}
	return nil
}

func (uc *UserUsecase) GetUser(ctx context.Context, id uint64) (users *User, err error) {
	users, err = uc.repo.Get(ctx, id)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("userUsecase.GetUser: %v", err)
		return nil, fmt.Errorf("userUsecase.GetUser: %w", err)
	}
	return users, nil
}

func (uc *UserUsecase) GetAllUsers(ctx context.Context) (users []*User, err error) {
	users, err = uc.repo.GetAll(ctx)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("userUsecase.GetAllUsers: %v", err)
		return nil, fmt.Errorf("userUsecase.GetAllUsers: %w", err)
	}
	return users, nil
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, user *User, updateFn func(context.Context, *User) (*User, error)) (err error) {
	err = uc.repo.Update(ctx, user, updateFn)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("userUsecase.UpdateUser: %v", err)
		return fmt.Errorf("userUsecase.UpdateUser: %w", err)
	}
	return nil
}

func (uc *UserUsecase) GetUserByEmail(ctx context.Context, email string) (user *User, err error) {
	user, err = uc.repo.GetUserByEmail(ctx, email)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("userUsecase.GetUserByEmail: %v", err)
		return nil, fmt.Errorf("userUsecase.GetUserByEmail: %w", err)
	}
	return user, nil
}

func validateUser(user *User) error {
	if user.Username == "" {
		return errors.ErrInvalidUsername
	}
	if user.Password == "" {
		return errors.ErrInvalidPassword
	}
	if user.Email == "" {
		return errors.ErrInvalidEmail
	}
	return nil
}
