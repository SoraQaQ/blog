package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/soraQaQ/blog/app/article/internal/biz/command"
	"github.com/soraQaQ/blog/app/article/internal/biz/query"
	"github.com/soraQaQ/blog/app/article/internal/domain"

	pb "github.com/soraQaQ/blog/api/article/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ArticleService struct {
	pb.UnimplementedArticleServiceServer

	log                   *log.Helper
	createCommand         command.CreateArticleHandler
	updateCommand         command.UpdateArticleHandler
	deleteCommand         command.DeleteArticleHandler
	getArticleQuery       query.GetArticleHandler
	getAllArticlesQuery   query.GetAllArticlesHandler
	getArticlesByTagQuery query.GetArticlesByTagHandler
}

func NewArticleService(logger log.Logger, createCommand command.CreateArticleHandler, updateCommand command.UpdateArticleHandler, deleteCommand command.DeleteArticleHandler, getArticleQuery query.GetArticleHandler, getAllArticlesQuery query.GetAllArticlesHandler, getArticlesByTagQuery query.GetArticlesByTagHandler) *ArticleService {
	return &ArticleService{log: log.NewHelper(logger), createCommand: createCommand, updateCommand: updateCommand, deleteCommand: deleteCommand, getArticleQuery: getArticleQuery, getAllArticlesQuery: getAllArticlesQuery, getArticlesByTagQuery: getArticlesByTagQuery}
}

func (s *ArticleService) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (_ *emptypb.Empty, err error) {
	article := &domain.Article{
		Id:         req.Article.Id,
		Title:      req.Article.Title,
		Summary:    req.Article.Summary,
		ContentUrl: req.Article.ContentUrl,
		Status:     req.Article.Status,
		ViewCount:  req.Article.ViewCount,
		Tags:       req.Article.Tags,
		ImageUrl:   req.Article.ImageUrl,
	}
	err = s.createCommand.Handler(ctx, command.CreateArticle{
		Article: article,
	})
	if err != nil {
		return
	}
	return
}

func (s *ArticleService) GetAllArticle(ctx context.Context, req *emptypb.Empty) (res *pb.GetAllArticleReply, err error) {
	articles, err := s.getAllArticlesQuery.Handler(ctx, struct{}{})
	if err != nil {
		return
	}
	pbArticles := make([]*pb.Article, len(articles))
	for _, article := range articles {
		pbArticles = append(pbArticles, &pb.Article{
			Id:         article.Id,
			Title:      article.Title,
			Summary:    article.Summary,
			ContentUrl: article.ContentUrl,
			Status:     article.Status,
			ViewCount:  article.ViewCount,
			Tags:       article.Tags,
			ImageUrl:   article.ImageUrl,
		})
	}
	res = &pb.GetAllArticleReply{
		Articles: pbArticles,
		Total:    int64(len(articles)),
	}
	return
}

func (s *ArticleService) GetArticleById(ctx context.Context, req *pb.GetArticleByIdRequest) (res *pb.GetArticleByIdReply, err error) {
	article, err := s.getArticleQuery.Handler(ctx, query.GetArticle{
		Id: req.Id,
	})
	if err != nil {
		return
	}
	res = &pb.GetArticleByIdReply{
		Article: &pb.Article{
			Id:         article.Id,
			Title:      article.Title,
			Summary:    article.Summary,
			ContentUrl: article.ContentUrl,
			Status:     article.Status,
			ViewCount:  article.ViewCount,
			Tags:       article.Tags,
			ImageUrl:   article.ImageUrl,
		},
	}
	return
}

func (s *ArticleService) GetArticlesByTag(ctx context.Context, req *pb.GetArticlesByTagRequest) (res *pb.GetArticlesByTagReply, err error) {
	articles, err := s.getArticlesByTagQuery.Handler(ctx, query.GetArticlesByTag{
		Tag: req.Tag,
	})
	if err != nil {
		return
	}
	pbArticles := make([]*pb.Article, len(articles))
	for _, article := range articles {
		pbArticles = append(pbArticles, &pb.Article{
			Id:         article.Id,
			Title:      article.Title,
			Summary:    article.Summary,
			ContentUrl: article.ContentUrl,
			Status:     article.Status,
			ViewCount:  article.ViewCount,
			Tags:       article.Tags,
			ImageUrl:   article.ImageUrl,
		})
	}
	res = &pb.GetArticlesByTagReply{
		Articles: pbArticles,
	}
	return
}

func (s *ArticleService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (_ *emptypb.Empty, err error) {
	err = s.updateCommand.Handler(ctx, command.UpdateArticle{
		Article: &domain.Article{
			Id:         req.Article.Id,
			Title:      req.Article.Title,
			Summary:    req.Article.Summary,
			ContentUrl: req.Article.ContentUrl,
			Status:     req.Article.Status,
			ViewCount:  req.Article.ViewCount,
			Tags:       req.Article.Tags,
			ImageUrl:   req.Article.ImageUrl,
		},
		UpdateFn: func(_ context.Context, oldArticle *domain.Article) (*domain.Article, error) {
			oldArticle.Id = req.Article.Id
			oldArticle.Title = req.Article.Title
			oldArticle.Summary = req.Article.Summary
			oldArticle.ContentUrl = req.Article.ContentUrl
			oldArticle.Status = req.Article.Status
			oldArticle.ViewCount = req.Article.ViewCount
			oldArticle.Tags = req.Article.Tags
			oldArticle.ImageUrl = req.Article.ImageUrl
			return oldArticle, nil
		},
	})
	if err != nil {
		return
	}
	return
}

func (s *ArticleService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (_ *emptypb.Empty, err error) {
	err = s.deleteCommand.Handler(ctx, command.DeleteArticle{
		Id: req.Id,
	})
	if err != nil {
		return
	}
	return
}
