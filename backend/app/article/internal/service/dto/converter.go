package dto

import (
	pb "github.com/soraQaQ/blog/api/article/v1"
	"github.com/soraQaQ/blog/app/article/internal/biz"
)

type ArticleConverter struct{}

func (a *ArticleConverter) EntityToProto(article *biz.Article) *pb.Article {
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

func (a *ArticleConverter) ProtoToEntity(in *pb.Article) *biz.Article {
	return &biz.Article{
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

func (a *ArticleConverter) EntitiesToProtos(articles []*biz.Article) (res []*pb.Article) {

	for _, article := range articles {
		res = append(res, a.EntityToProto(article))
	}
	return
}

func (a *ArticleConverter) ProtosToEntities(articles []*pb.Article) (res []*biz.Article) {
	for _, article := range articles {
		res = append(res, a.ProtoToEntity(article))
	}
	return
}
