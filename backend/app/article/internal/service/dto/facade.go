package dto

import "sync"

var (
	userCv   *ArticleConverter
	userOnce sync.Once
)

func NewArticleConverter() *ArticleConverter {
	userOnce.Do(func() {
		userCv = new(ArticleConverter)
	})
	return userCv
}
