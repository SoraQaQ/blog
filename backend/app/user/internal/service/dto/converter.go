package dto

import (
	userpb "github.com/soraQaQ/blog/api/user/v1"
	"github.com/soraQaQ/blog/app/user/internal/biz"
)

type UserConverter struct{}

func (c *UserConverter) EntitiesToProtos(users []*biz.User) (res []*userpb.User) {
	for _, user := range users {
		res = append(res, c.EntityToProto(user))
	}
	return
}

func (c *UserConverter) EntityToProto(u *biz.User) *userpb.User {
	return &userpb.User{
		Id:       u.Id,
		UserName: u.Username,
		NickName: u.Nickname,
		Password: u.Password,
		Email:    u.Email,
	}
}

func (c *UserConverter) ProtoToEntity(u *userpb.User) *biz.User {

	return &biz.User{
		Id:       u.Id,
		Username: u.UserName,
		Nickname: u.NickName,
		Password: u.Password,
		Email:    u.Email,
	}
}
