package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	articlepb "github.com/soraQaQ/blog/api/article/v1"
	"github.com/soraQaQ/blog/app/admin/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ArticleRepo struct {
	data *Data
	log  *log.Helper
}

func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &ArticleRepo{data: data, log: log.NewHelper(log.With(logger, "module", "data/article"))}
}

func (a *ArticleRepo) Save(ctx context.Context, article *biz.Article) error {
	_, err := a.data.ac.CreateArticle(ctx, &articlepb.CreateArticleRequest{
		Article: &articlepb.Article{
			Id:         article.Id,
			Title:      article.Title,
			Summary:    article.Summary,
			ContentUrl: article.ContentUrl,
			Status:     article.Status,
			ViewCount:  article.ViewCount,
			Tags:       article.Tags,
			ImageUrl:   article.ImageUrl,
		},
	})

	if err != nil {
		return err
	}
	return nil
}

func (a *ArticleRepo) Get(ctx context.Context, i int64) (*biz.Article, error) {
	reply, err := a.data.ac.GetArticleById(ctx, &articlepb.GetArticleByIdRequest{
		Id: i,
	})
	if err != nil {
		return nil, err
	}
	article := &biz.Article{
		Id:         reply.Article.Id,
		Title:      reply.Article.Title,
		Summary:    reply.Article.Summary,
		ContentUrl: reply.Article.ContentUrl,
		Status:     reply.Article.Status,
		ViewCount:  reply.Article.ViewCount,
		Tags:       reply.Article.Tags,
		ImageUrl:   reply.Article.ImageUrl,
	}
	return article, nil
}

func (a *ArticleRepo) Update(ctx context.Context, article *biz.Article) error {
	_, err := a.data.ac.UpdateArticle(ctx, &articlepb.UpdateArticleRequest{
		Article: &articlepb.Article{
			Id:         article.Id,
			Title:      article.Title,
			Summary:    article.Summary,
			ContentUrl: article.ContentUrl,
			Status:     article.Status,
			ViewCount:  article.ViewCount,
			Tags:       article.Tags,
			ImageUrl:   article.ImageUrl,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (a *ArticleRepo) GetAll(ctx context.Context) ([]*biz.Article, error) {
	reply, err := a.data.ac.GetAllArticle(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	articles := make([]*biz.Article, 0, len(reply.Articles))
	for _, article := range reply.Articles {
		articles = append(articles, &biz.Article{
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
	return articles, nil
}

func (a *ArticleRepo) GetArticlesByTag(ctx context.Context, tag string) ([]*biz.Article, error) {
	reply, err := a.data.ac.GetArticlesByTag(ctx, &articlepb.GetArticlesByTagRequest{
		Tag: tag,
	})
	if err != nil {
		return nil, err
	}
	articles := make([]*biz.Article, 0, len(reply.Articles))
	for _, article := range reply.Articles {
		articles = append(articles, &biz.Article{
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
	return articles, nil
}

func (a *ArticleRepo) Delete(ctx context.Context, i int64) error {
	_, err := a.data.ac.DeleteArticle(ctx, &articlepb.DeleteArticleRequest{
		Id: i,
	})
	if err != nil {
		return err
	}
	return nil
}
