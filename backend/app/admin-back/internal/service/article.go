package service

import (
	"context"
	adminpb "github.com/soraQaQ/blog/api/admin/v1"
	"github.com/soraQaQ/blog/app/admin/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *AdminService) CreateArticle(ctx context.Context, req *adminpb.CreateArticleRequest) (*emptypb.Empty, error) {
	article := &biz.Article{
		Id:         req.Article.Id,
		Title:      req.Article.Title,
		Summary:    req.Article.Summary,
		ContentUrl: req.Article.ContentUrl,
		Status:     req.Article.Status,
		ViewCount:  req.Article.ViewCount,
		Tags:       req.Article.Tags,
		ImageUrl:   req.Article.ImageUrl,
	}
	err := s.articleUseCase.Save(ctx, article)
	if err != nil {
		return &emptypb.Empty{}, nil
	}
	return &emptypb.Empty{}, nil
}

func (s *AdminService) ListArticle(ctx context.Context, req *emptypb.Empty) (*adminpb.ListArticleReply, error) {
	articles, err := s.articleUseCase.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	articleList := make([]*adminpb.Article, 0, len(articles))
	for _, article := range articles {
		articleList = append(articleList, &adminpb.Article{
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
	return &adminpb.ListArticleReply{
		Items: articleList,
		Total: int32(len(articles)),
	}, nil
}

func (s *AdminService) GetArticleById(ctx context.Context, req *adminpb.GetArticleRequest) (*adminpb.GetArticleReply, error) {
	article, err := s.articleUseCase.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	a := &adminpb.Article{
		Id:         article.Id,
		Title:      article.Title,
		Summary:    article.Summary,
		ContentUrl: article.ContentUrl,
		Status:     article.Status,
		ViewCount:  article.ViewCount,
		Tags:       article.Tags,
		ImageUrl:   article.ImageUrl,
	}
	return &adminpb.GetArticleReply{
		Article: a,
	}, nil
}

func (s *AdminService) GetArticlesByTag(ctx context.Context, req *adminpb.GetArticlesByTagRequest) (*adminpb.ListArticleReply, error) {
	articles, err := s.articleUseCase.GetArticlesByTag(ctx, req.Tag)
	if err != nil {
		return nil, err
	}
	articleList := make([]*adminpb.Article, 0, len(articles))
	for _, article := range articles {
		articleList = append(articleList, &adminpb.Article{
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
	return &adminpb.ListArticleReply{
		Items: articleList,
		Total: int32(len(articles)),
	}, nil
}

func (s *AdminService) UpdateArticle(ctx context.Context, req *adminpb.UpdateArticleRequest) (*emptypb.Empty, error) {
	article := &biz.Article{
		Id:         req.Article.Id,
		Title:      req.Article.Title,
		Summary:    req.Article.Summary,
		ContentUrl: req.Article.ContentUrl,
		Status:     req.Article.Status,
		ViewCount:  req.Article.ViewCount,
		Tags:       req.Article.Tags,
		ImageUrl:   req.Article.ImageUrl,
	}
	err := s.articleUseCase.Update(ctx, article)
	if err != nil {
		return &emptypb.Empty{}, nil
	}
	return &emptypb.Empty{}, nil
}

func (s *AdminService) DeleteArticle(ctx context.Context, req *adminpb.DeleteArticleRequest) (*emptypb.Empty, error) {
	err := s.articleUseCase.Delete(ctx, req.Id)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}
