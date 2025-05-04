package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	userpb "github.com/soraQaQ/blog/api/user/v1"
	"github.com/soraQaQ/blog/app/admin/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{data: data, log: log.NewHelper(log.With(logger, "moudle", "data/user"))}
}

func (u *UserRepo) Save(ctx context.Context, user *biz.User) error {
	_, err := u.data.uc.CreateUser(ctx, &userpb.CreateUserRequest{
		User: &userpb.User{
			Id:       user.Id,
			UserName: user.Username,
			Password: user.Password,
			NickName: user.Nickname,
			Email:    user.Email,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) Get(ctx context.Context, id uint64) (*biz.User, error) {
	reply, err := u.data.uc.GetUser(ctx, &userpb.GetUserRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:       reply.User.Id,
		Username: reply.User.UserName,
		Nickname: reply.User.NickName,
		Password: reply.User.Password,
		Email:    reply.User.Email,
	}, nil
}

func (u *UserRepo) Update(ctx context.Context, user *biz.User) error {
	_, err := u.data.uc.UpdateUser(ctx, &userpb.UpdateUserRequest{
		Id:       user.Id,
		UserName: user.Username,
		NickName: user.Nickname,
	})
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) GetAll(ctx context.Context) ([]*biz.User, error) {
	reply, err := u.data.uc.GetAllUser(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	var users []*biz.User
	for _, user := range reply.Users {
		users = append(users, &biz.User{
			Id:       user.Id,
			Username: user.UserName,
			Nickname: user.NickName,
			Password: user.Password,
			Email:    user.Email,
		})
	}
	return users, nil
}

func (u *UserRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	reply, err := u.data.uc.GetUserByEmail(ctx, &userpb.GetUserByEmailRequest{
		Email: email,
	})
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:       reply.User.Id,
		Username: reply.User.UserName,
		Nickname: reply.User.NickName,
		Password: reply.User.Password,
		Email:    reply.User.Email,
	}, nil
}
