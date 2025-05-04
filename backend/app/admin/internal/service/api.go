package service

import (
	"context"
	"fmt"
	"github.com/soraQaQ/blog/app/admin/internal/biz"
	"github.com/soraQaQ/blog/app/admin/internal/conf"
	"time"

	"github.com/soraQaQ/blog/pkg/util"

	adminpb "github.com/soraQaQ/blog/api/admin/v1"
	"github.com/soraQaQ/blog/pkg/auth"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/go-kratos/kratos/v2/log"
)

type AdminService struct {
	adminpb.UnimplementedAdminServiceServer
	userUseCase    *biz.UserUseCase
	articleUseCase *biz.ArticleUsecase
	log            *log.Helper
	ra             *conf.Auth
}

func NewAdminService(logger log.Logger, ra *conf.Auth, us *biz.UserUseCase, as *biz.ArticleUsecase) *AdminService {
	return &AdminService{log: log.NewHelper(logger), ra: ra, userUseCase: us, articleUseCase: as}
}

func (s *AdminService) ListUser(ctx context.Context, _ *emptypb.Empty) (*adminpb.ListUserReply, error) {
	users, err := s.userUseCase.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var replyUsers []*adminpb.User
	for _, user := range users {
		replyUsers = append(replyUsers, &adminpb.User{
			Id:       &user.Id,
			UserName: &user.Username,
			Password: &user.Password,
			NickName: &user.Nickname,
			Email:    &user.Email,
			Token:    nil,
		})
	}

	return &adminpb.ListUserReply{
		Items: replyUsers,
		Total: int32(len(users)),
	}, nil
}

func (s *AdminService) Login(ctx context.Context, req *adminpb.LoginReq) (*adminpb.LoginReply, error) {
	user, err := s.userUseCase.GetUserByEmail(ctx, req.Email)
	if err != nil {
		s.log.Errorf("s.userClient.GetUser err: %+v", err)
		return nil, fmt.Errorf("s.userClient.GetUserByEmail err: %+v", err)
	}

	if !util.CheckPasswordHash(req.GetPassword(), user.Password) {
		return nil, fmt.Errorf("password error")
	}

	id := user.Id

	jwt := auth.NewJWT(s.ra.ApiKey)

	token, err := jwt.GenerateToken(id)

	if err != nil {
		s.log.Errorf("jwt.GenerateToken err: %+v", err)
		return nil, err
	}

	return &adminpb.LoginReply{
		Id:    id,
		Token: token,
	}, nil
}

func (s *AdminService) Logout(ctx context.Context, req *adminpb.LogoutReq) (*adminpb.LogoutReply, error) {
	return &adminpb.LogoutReply{}, nil
}

func (s *AdminService) Register(ctx context.Context, req *adminpb.RegisterReq) (_ *emptypb.Empty, err error) {
	hashPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		s.log.Errorf("util.HashPassword err: %+v", err)
	}
	user := &biz.User{
		Id:       uint64(time.Now().Unix()),
		Username: req.Username,
		Password: hashPassword,
		Nickname: req.Nickname,
		Email:    req.Email,
	}

	err = s.userUseCase.Save(ctx, user)

	if err != nil {
		s.log.Errorf("s.userClient.CreateUser err: %+v", err)
		return nil, fmt.Errorf("s.userClient.CreateUser err: %+v", err)
	}

	return &emptypb.Empty{}, nil
}
