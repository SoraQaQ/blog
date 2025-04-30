package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/soraQaQ/blog/app/user/internal/conf"
	"github.com/soraQaQ/blog/app/user/internal/data/memory"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo, memory.NewUserMemoryRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *memory.UserMemoryRepo
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *memory.UserMemoryRepo) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db}, cleanup, nil
}
