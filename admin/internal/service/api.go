package service

import (
	v1 "admin/api/admin/v1"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/go-kratos/kratos/v2/log"
)

type AdminService struct {
	v1.UnimplementedAdminServiceServer

	log *log.Helper
}

func NewAdminService(logger log.Logger) *AdminService {

	return &AdminService{log: log.NewHelper(logger)}
}

func (s *AdminService) ListUser(ctx context.Context, _ *emptypb.Empty) (*v1.ListUserReply, error) {
	return &v1.ListUserReply{}, nil
}

func (s *AdminService) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginReply, error) {
	return &v1.LoginReply{}, nil
}

func (s *AdminService) Logout(ctx context.Context, req *v1.LogoutReq) (*v1.LogoutReply, error) {
	return &v1.LogoutReply{}, nil
}

func (s *AdminService) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterReply, error) {
	return &v1.RegisterReply{}, nil
}
