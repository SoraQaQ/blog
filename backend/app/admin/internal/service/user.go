package service

import (
	"context"
	adminpb "github.com/soraQaQ/blog/api/admin/v1"
	"github.com/soraQaQ/blog/app/admin/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *AdminService) UpdateUser(ctx context.Context, req *adminpb.UpdateUserRequest) (*emptypb.Empty, error) {
	user := &biz.User{
		Id:       req.Id,
		Username: req.UserName,
		Nickname: req.NickName,
		Password: req.Password,
	}
	err := s.userUseCase.Update(ctx, user)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil

}
