package memory

import (
	"context"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/soraQaQ/blog/app/article/internal/biz"
	"github.com/soraQaQ/blog/pkg/errors"
)

type ArticleMemory struct {
	store []*biz.Article
	log   *log.Helper
	lock  *sync.RWMutex
}

func NewArticleMemoryRepo(logger log.Logger) *ArticleMemory {
	s := make([]*biz.Article, 0)
	return &ArticleMemory{
		store: s,
		log:   log.NewHelper(logger),
		lock:  &sync.RWMutex{},
	}
}

func (a *ArticleMemory) Save(ctx context.Context, article *biz.Article) (*biz.Article, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	newArticle := &biz.Article{
		Id:         article.Id,
		Title:      article.Title,
		Summary:    article.Summary,
		ContentUrl: article.ContentUrl,
		ViewCount:  article.ViewCount,
		Tags:       article.Tags,
		ImageUrl:   article.ImageUrl,
	}
	a.store = append(a.store, newArticle)
	a.log.Infof("create article %+v", newArticle)
	return newArticle, nil
}

func (a *ArticleMemory) Get(ctx context.Context, id int64) (*biz.Article, error) {
	a.lock.RLock()
	defer a.lock.RUnlock()
	for _, article := range a.store {
		if article.Id == id {
			return article, nil
		}
	}
	return nil, errors.ErrorArticleNotFound
}

func (a *ArticleMemory) GetAll(ctx context.Context) ([]*biz.Article, error) {
	a.lock.RLock()
	defer a.lock.RUnlock()
	return a.store, nil
}

func (a *ArticleMemory) GetArticlesByTag(ctx context.Context, s string) ([]*biz.Article, error) {
	a.lock.RLock()
	defer a.lock.RUnlock()
	articles := make([]*biz.Article, 0)
	for _, article := range a.store {
		if article.Tags == s {
			articles = append(articles, article)
		}
	}
	return articles, nil
}

func (a *ArticleMemory) Delete(ctx context.Context, id int64) error {
	a.lock.Lock()
	defer a.lock.Unlock()
	found := false
	for i, article := range a.store {
		if article.Id == id {
			found = true
			a.store = append(a.store[:i], a.store[i+1:]...)
		}
	}
	if !found {
		return errors.ErrorArticleNotFound
	}
	return nil

}

func (a *ArticleMemory) Update(ctx context.Context, article *biz.Article, updateFn func(context.Context, *biz.Article) (*biz.Article, error)) error {
	a.lock.Lock()
	defer a.lock.Unlock()
	found := false
	for i, ar := range a.store {
		if ar.Id == article.Id {
			found = true
			updateArticle, err := updateFn(ctx, article)
			if err != nil {
				return err
			}
			a.store[i] = updateArticle
		}
	}
	if !found {
		return errors.ErrorArticleNotFound
	}
	return nil
}
