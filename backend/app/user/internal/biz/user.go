package biz

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/soraQaQ/blog/pkg/util"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/soraQaQ/blog/pkg/errors"
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
	uc.log.Debugf("create user: %v", user)
	if err = validateUser(user); err != nil {
		uc.log.WithContext(ctx).Errorf("userUsecase.CreateUser: %v", err)
		return
	}
	hashPassword, err := util.HashPassword(user.Password)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("userUsecase.CreateUser: %v", err)
		return
	}

	uc.log.Debugf("hash password: %v", hashPassword)
	uc.log.Debugf("user: %v", user)
	newUser := &User{
		Id:       uint64(time.Now().UnixNano()),
		Username: user.Username,
		Nickname: user.Nickname,
		Password: hashPassword,
		Email:    user.Email,
	}

	err = uc.repo.Save(ctx, newUser)
	uc.log.WithContext(ctx).Debugf("memory save success: %v", err)
	if err != nil {
		return
	}
	return
}

func (uc *UserUsecase) GetUser(ctx context.Context, id uint64) (users *User, err error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid user id")
	}
	users, err = uc.repo.Get(ctx, id)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("userUsecase.GetUser: %v", err)
		return nil, fmt.Errorf("userUsecase.GetUser: %w", err)
	}
	return
}

func (uc *UserUsecase) GetAllUsers(ctx context.Context) (users []*User, err error) {
	users, err = uc.repo.GetAll(ctx)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("userUsecase.GetAllUsers: %v", err)
		return nil, fmt.Errorf("userUsecase.GetAllUsers: %w", err)
	}
	return
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, user *User, updateFn func(context.Context, *User) (*User, error)) (err error) {
	//if err = validateUser(user); err != nil {
	//	return err
	//}

	oldUser, err := uc.GetUser(ctx, user.Id)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("userUsecase.UpdateUser: %v", err)
		return errors.ErrUserNotFound
	}

	user = &User{
		Id:       user.Id,
		Username: user.Username,
		Nickname: user.Nickname,
		Password: oldUser.Password,
		Email:    oldUser.Email,
	}

	err = uc.repo.Update(ctx, user, updateFn)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("userUsecase.UpdateUser: %v", err)
		return fmt.Errorf("userUsecase.UpdateUser: %w", err)
	}
	return
}

func (uc *UserUsecase) GetUserByEmail(ctx context.Context, email string) (user *User, err error) {
	if err = validateEmail(email); err != nil {
		return nil, err
	}
	user, err = uc.repo.GetUserByEmail(ctx, email)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("userUsecase.GetUserByEmail: %v", err)
		return nil, fmt.Errorf("userUsecase.GetUserByEmail: %w", err)
	}
	return
}

func validateUser(user *User) error {
	if user.Username == "" {
		return errors.ErrInvalidUsername
	}
	if user.Password == "" {
		return errors.ErrInvalidPassword
	}
	if err := validateEmail(user.Email); err != nil {
		return err
	}
	if user.Id <= 0 {
		return errors.ErrorArticleInvalid
	}
	return nil
}

func validateEmail(email string) error {
	if email == "" {
		return errors.ErrInvalidEmail
	}

	// 使用正则表达式验证邮箱格式
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(pattern, email)
	if err != nil {
		return fmt.Errorf("验证邮箱格式时发生错误: %w", err)
	}

	if !matched {
		return errors.ErrInvalidEmail
	}

	return nil
}
