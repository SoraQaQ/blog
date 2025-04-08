package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/soraQaQ/blog/app/user/internal/conf"
	"github.com/soraQaQ/blog/app/user/internal/data/memory"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	userMemory *memory.UserMemoryRepo
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	userMemory := memory.NewUserMemoryRepo(logger)
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{userMemory: userMemory}, cleanup, nil
}
