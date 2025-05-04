package convertor

import (
	pb "github.com/soraQaQ/blog/api/article/v1"
	userpb "github.com/soraQaQ/blog/api/user/v1"
	"github.com/soraQaQ/blog/pkg/entity"
)

type UserConverter struct{}

func (c *UserConverter) EntitiesToProtos(users []*entity.User) (res []*userpb.User) {
	for _, user := range users {
		res = append(res, c.EntityToProto(user))
	}
	return
}

func (c *UserConverter) EntityToProto(u *entity.User) *userpb.User {
	return &userpb.User{
		Id:       u.Id,
		UserName: u.Username,
		NickName: u.Nickname,
		Password: u.Password,
		Email:    u.Email,
	}
}

func (c *UserConverter) ProtoToEntity(u *userpb.User) *entity.User {

	return &entity.User{
		Id:       u.Id,
		Username: u.UserName,
		Nickname: u.NickName,
		Password: u.Password,
		Email:    u.Email,
	}
}

type ArticleConverter struct{}

func (a *ArticleConverter) EntityToProto(article *entity.Article) *pb.Article {
	return &pb.Article{
		Id:         article.Id,
		Title:      article.Title,
		Summary:    article.Summary,
		ContentUrl: article.ContentUrl,
		Status:     article.Status,
		ViewCount:  article.ViewCount,
		Tags:       article.Tags,
		ImageUrl:   article.ImageUrl,
	}
}

func (a *ArticleConverter) ProtoToEntity(in *pb.Article) *entity.Article {
	return &entity.Article{
		Id:         in.Id,
		Title:      in.Title,
		Summary:    in.Summary,
		ContentUrl: in.ContentUrl,
		Status:     in.Status,
		ViewCount:  in.ViewCount,
		Tags:       in.Tags,
		ImageUrl:   in.ImageUrl,
	}
}

func (a *ArticleConverter) EntitiesToProtos(articles []*entity.Article) (res []*pb.Article) {

	for _, article := range articles {
		res = append(res, a.EntityToProto(article))
	}
	return
}

func (a *ArticleConverter) ProtosToEntities(articles []*pb.Article) (res []*entity.Article) {
	for _, article := range articles {
		res = append(res, a.ProtoToEntity(article))
	}
	return
}
