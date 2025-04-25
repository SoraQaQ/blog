package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	userpb "github.com/soraQaQ/blog/api/user/v1"
	"github.com/soraQaQ/blog/app/user/internal/biz"
	"github.com/soraQaQ/blog/app/user/internal/service/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *UserService) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserReply, error) {
	id := req.GetId()
	user, err := s.uc.GetUser(ctx, id)
	if err != nil {
		s.log.Errorf("GetUser err: %v", err)
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return &userpb.GetUserReply{User: dto.NewUserConverter().EntityToProto(user)}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*emptypb.Empty, error) {
	s.log.Infof("UpdateUser req.User %d", req.GetId())
	err := s.uc.UpdateUser(
		ctx,
		dto.NewUserConverter().ProtoToEntity(&userpb.User{
			Id:       req.Id,
			UserName: req.UserName,
			NickName: req.NickName,
		}),
		func(ctx context.Context, user *biz.User) (*biz.User, error) {
			return user, nil
		},
	)
	if err != nil {
		s.log.Errorf("UpdateUser err: %v", err)
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (s *UserService) CreateUser(ctx context.Context, user *userpb.CreateUserRequest) (*emptypb.Empty, error) {
	s.log.Infof("CreateUser user: %+v", user)

	err := s.uc.CreateUser(ctx, dto.NewUserConverter().ProtoToEntity(user.User))

	if err != nil {
		s.log.Errorf("CreateUser 失败: %v", err)
		return nil, status.Error(codes.NotFound, err.Error())
	}
	
	return &emptypb.Empty{}, nil
}

func (s *UserService) GetAllUser(ctx context.Context, _ *emptypb.Empty) (res *userpb.GetUserAllReply, err error) {
	users, err := s.uc.GetAllUsers(ctx)
	if err != nil {
		s.log.Errorf("GetAllUser err: %v", err)
		return nil, status.Error(codes.NotFound, err.Error())
	}

	res = &userpb.GetUserAllReply{
		Users: dto.NewUserConverter().EntitiesToProtos(users),
		Total: int64(len(users)),
	}
	return
}

func (s *UserService) GetUserByEmail(ctx context.Context, req *userpb.GetUserByEmailRequest) (res *userpb.GetUserByEmailReply, err error) {
	s.log.Infof("GetUserByEmail")
	user, err := s.uc.GetUserByEmail(ctx, req.Email)
	if err != nil {
		s.log.Errorf("GetUserByEmail err: %v", err)
		return nil, status.Error(codes.NotFound, err.Error())
	}
	res = &userpb.GetUserByEmailReply{
		User: dto.NewUserConverter().EntityToProto(user),
	}
	return
}
