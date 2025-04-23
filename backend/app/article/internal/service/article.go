package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/soraQaQ/blog/app/article/internal/biz"

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

func (s *ArticleService) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *ArticleService) GetAllArticle(ctx context.Context, req *emptypb.Empty) (*pb.GetAllArticleReply, error) {
	return &pb.GetAllArticleReply{}, nil
}

func (s *ArticleService) GetArticleById(ctx context.Context, req *pb.GetArticleByIdRequest) (*pb.GetArticleByIdReply, error) {
	return &pb.GetArticleByIdReply{}, nil
}

func (s *ArticleService) GetArticleByTags(ctx context.Context, req *pb.GetArticleByTagsRequest) (*pb.GetArticleByTagsReply, error) {
	return &pb.GetArticleByTagsReply{}, nil
}

func (s *ArticleService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleReply, error) {
	return &pb.UpdateArticleReply{}, nil
}
