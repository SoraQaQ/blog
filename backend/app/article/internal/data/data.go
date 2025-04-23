package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/soraQaQ/blog/app/article/internal/conf"
	"github.com/soraQaQ/blog/app/article/internal/data/memory"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *memory.ArticleMemory
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	db := memory.NewArticleMemoryRepo(logger)
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db}, cleanup, nil
}
