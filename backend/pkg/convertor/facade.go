package convertor

import "sync"

var (
	userCv      *UserConverter
	userOnce    sync.Once
	articleCv   *ArticleConverter
	articleOnce sync.Once
)

func NewUserConverter() *UserConverter {
	userOnce.Do(func() {
		userCv = new(UserConverter)
	})
	return userCv
}

func NewArticleConverter() *ArticleConverter {
	articleOnce.Do(func() {
		articleCv = new(ArticleConverter)
	})
	return articleCv
}
