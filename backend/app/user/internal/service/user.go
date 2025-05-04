package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	userpb "github.com/soraQaQ/blog/api/user/v1"
	"github.com/soraQaQ/blog/app/user/internal/biz/command"
	"github.com/soraQaQ/blog/app/user/internal/biz/query"
	"github.com/soraQaQ/blog/app/user/internal/domain"
	"github.com/soraQaQ/blog/pkg/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	userpb.UnimplementedUserServiceServer

	createCommand        command.CreateUserHandler
	updateCommand        command.UpdateUserHandler
	getAllUserQuery      query.GetAllUserHandler
	getUserQuery         query.GetUserHandler
	getUsersByEmailQuery query.GetUsersByEmailHandler
	log                  *log.Helper
}

func NewUserService(
	createCommand command.CreateUserHandler,
	updateCommand command.UpdateUserHandler,
	getUserQuery query.GetUserHandler,
	getAllUserQuery query.GetAllUserHandler,
	getUsersByEmailQuery query.GetUsersByEmailHandler,
	logger log.Logger,
) *UserService {
	return &UserService{
		createCommand:        createCommand,
		updateCommand:        updateCommand,
		getAllUserQuery:      getAllUserQuery,
		getUserQuery:         getUserQuery,
		getUsersByEmailQuery: getUsersByEmailQuery,
		log:                  log.NewHelper(logger),
	}
}

func (s *UserService) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserReply, error) {
	user, err := s.getUserQuery.Handler(ctx, query.GetUser{
		req.Id,
	})
	if err != nil {
		s.log.Errorf("GetUser err: %v", err)
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return &userpb.GetUserReply{User: &userpb.User{
		Id:       user.Id,
		UserName: user.Username,
		Password: user.Password,
		NickName: user.Nickname,
		Email:    user.Email,
	}}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*emptypb.Empty, error) {
	s.log.Infof("UpdateUser req.User %d", req.GetId())
	err := s.updateCommand.Handler(ctx, command.UpdateUser{
		User: &domain.User{
			Id:       req.Id,
			Username: req.UserName,
			Nickname: req.NickName,
			Password: req.Password,
		},
		UpdateFn: func(ctx context.Context, oldUser *domain.User) (*domain.User, error) {
			hashPassword, err := util.HashPassword(req.Password)
			if err != nil {
				return nil, err
			}
			oldUser.Password = hashPassword
			oldUser.Nickname = req.NickName
			oldUser.Username = req.UserName
			return oldUser, nil
		},
	})
	if err != nil {
		s.log.Errorf("UpdateUser err: %v", err)
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*emptypb.Empty, error) {
	err := s.createCommand.Handler(ctx, command.CreateUser{
		User: &domain.User{
			Id:       req.User.Id,
			Username: req.User.UserName,
			Nickname: req.User.NickName,
			Password: req.User.Password,
			Email:    req.User.Email,
		},
	})
	if err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}

func (s *UserService) GetAllUser(ctx context.Context, _ *emptypb.Empty) (res *userpb.GetUserAllReply, err error) {
	users, err := s.getAllUserQuery.Handler(ctx, query.GetAllUser{})
	if err != nil {
		return
	}
	var pbUsers []*userpb.User
	for _, user := range users {
		pbUsers = append(pbUsers, &userpb.User{
			Id:       user.Id,
			UserName: user.Username,
			Password: user.Password,
			NickName: user.Nickname,
			Email:    user.Email,
		})
	}
	res = &userpb.GetUserAllReply{
		Users: pbUsers,
		Total: int64(len(pbUsers)),
	}
	return
}

func (s *UserService) GetUserByEmail(ctx context.Context, req *userpb.GetUserByEmailRequest) (res *userpb.GetUserByEmailReply, err error) {
	user, err := s.getUsersByEmailQuery.Handler(ctx, query.GetUserByEmail{
		Email: req.Email,
	})
	if err != nil {
		return
	}
	res = &userpb.GetUserByEmailReply{
		User: &userpb.User{
			Id:       user.Id,
			UserName: user.Username,
			Password: user.Password,
			NickName: user.Nickname,
			Email:    user.Email,
		},
	}
	return
}
