package service

import (
	"context"
	"github.com/soraQaQ/blog/app/user/internal/dto"
	"github.com/soraQaQ/blog/pkg/util"

	"github.com/go-kratos/kratos/v2/log"
	userpb "github.com/soraQaQ/blog/api/user/v1"
	"github.com/soraQaQ/blog/app/user/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	userpb.UnimplementedUserServiceServer

	uc  *biz.UserUsecase
	log *log.Helper
}

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserService) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*emptypb.Empty, error) {
	s.log.Infof("CreateUser %+v", req.User)

	hashPassword, err := util.HashPassword(req.User.Password)
	if err != nil {
		s.log.Errorf("CreateUser hash password err: %+v", err)
		return nil, err
	}
	user := &biz.User{
		Id:       req.User.Id,
		Username: req.User.UserName,
		Nickname: req.User.NickName,
		Password: hashPassword,
		Email:    req.User.Email,
	}
	err = s.uc.CreateUser(ctx, user)
	if err != nil {
		s.log.Errorf("CreateUser err: %v", err)
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUsersReply, error) {
	id := req.GetId()
	user, err := s.uc.GetUser(ctx, id)
	if err != nil {
		s.log.Errorf("GetUser err: %v", err)
		return nil, err
	}

	return &userpb.GetUsersReply{
		User: dto.NewUserConverter().EntityToProto(user),
	}, nil
}

func (s *UserService) GetAllUser(ctx context.Context, _ *emptypb.Empty) (*userpb.GetUserAllReply, error) {
	users, err := s.uc.GetAllUsers(ctx)
	if err != nil {
		s.log.Errorf("GetAllUser err: %v", err)
		return nil, err
	}

	return &userpb.GetUserAllReply{
		Users: dto.NewUserConverter().EntitiesToProtos(users),
		Total: int64(len(users)),
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *userpb.UpdateRequest) (*userpb.UpdateReply, error) {
	s.log.Infof("UpdateUser req.User %+v", req.User)
	err := s.uc.UpdateUser(
		ctx,
		dto.NewUserConverter().ProtoToEntity(req.User),
		func(ctx context.Context, user *biz.User) (*biz.User, error) {
			return user, nil
		},
	)
	if err != nil {
		s.log.Errorf("UpdateUser err: %v", err)
		return &userpb.UpdateReply{
			Message: "update unsuccessfully",
			Success: "unsuccessfully",
		}, err
	}
	return &userpb.UpdateReply{
		Message: "update success",
		Success: "success",
	}, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, req *userpb.GetUserByEmailRequest) (*userpb.GetUsersReply, error) {
	s.log.Infof("GetUserByEmail")
	user, err := s.uc.GetUserByEmail(ctx, req.Email)
	if err != nil {
		s.log.Errorf("GetUserByEmail err: %v", err)
		return nil, err
	}
	return &userpb.GetUsersReply{User: dto.NewUserConverter().EntityToProto(user)}, nil
}
