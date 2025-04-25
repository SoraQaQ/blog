package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/soraQaQ/blog/app/article/internal/biz"
	"github.com/soraQaQ/blog/app/article/internal/service/dto"

	pb "github.com/soraQaQ/blog/api/article/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ArticleService struct {
	pb.UnimplementedArticleServiceServer

	log *log.Helper
	uc  *biz.ArticleUsecase
}

func NewArticleService(uc *biz.ArticleUsecase, logger log.Logger) *ArticleService {
	return &ArticleService{uc: uc, log: log.NewHelper(logger)}
}

func (s *ArticleService) CreateArticle(ctx context.Context, article *pb.Article) (_ *emptypb.Empty, err error) {
	s.log.Infof("CreateArticle %+v", article)
	err = s.uc.CreateArticle(ctx, dto.NewArticleConverter().ProtoToEntity(article))
	if err != nil {
		return
	}
	return
}
func (s *ArticleService) GetAllArticle(ctx context.Context, req *emptypb.Empty) (res *pb.GetAllArticleReply, err error) {
	s.log.Infof("GetAllArticle %+v", req)
	articles, err := s.uc.GetAllArticles(ctx)
	if err != nil {
		return
	}
	res = &pb.GetAllArticleReply{Articles: dto.NewArticleConverter().EntitiesToProtos(articles), Total: int64(len(articles))}
	return

}
func (s *ArticleService) GetArticleById(ctx context.Context, req *pb.GetArticleByIdRequest) (res *pb.GetArticleByIdReply, err error) {
	s.log.Infof("GetArticleById %+v", req)
	article, err := s.uc.Get(ctx, req.GetId())
	if err != nil {
		return
	}
	res = &pb.GetArticleByIdReply{
		Article: dto.NewArticleConverter().EntityToProto(article),
	}
	return
}

func (s *ArticleService) GetArticlesByTag(ctx context.Context, req *pb.GetArticlesByTagRequest) (res *pb.GetArticlesByTagReply, err error) {
	s.log.Infof("GetArticleById %+v", req)
	articles, err := s.uc.GetArticlesByTag(ctx, req.GetTag())
	if err != nil {
		return
	}
	res = &pb.GetArticlesByTagReply{
		Articles: dto.NewArticleConverter().EntitiesToProtos(articles),
	}
	return
}

func (s *ArticleService) UpdateArticle(ctx context.Context, article *pb.Article) (_ *emptypb.Empty, err error) {
	s.log.Infof("UpdateArticle %+v", article)
	err = s.uc.UpdateArticle(ctx, dto.NewArticleConverter().ProtoToEntity(article), func(ctx context.Context, article *biz.Article) (*biz.Article, error) {
		return article, nil
	})
	if err != nil {
		return
	}
	return
}

func (s *ArticleService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (_ *emptypb.Empty, err error) {
	s.log.Infof("DeleteArticle %+v", req)
	err = s.uc.DeleteArticle(ctx, req.GetId())
	if err != nil {
		return
	}
	return
}
