package biz

import (
	"github.com/google/wire"
	"github.com/soraQaQ/blog/app/article/internal/biz/command"
	"github.com/soraQaQ/blog/app/article/internal/biz/query"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(command.NewCreateOrderHandler,
	command.NewUpdateArticleHandler,
	command.NewDeleteArticleHandler,
	query.NewGetArticleHandler,
	query.NewGetArticlesByTagHandler,
	query.NewGetAllArticleHandler,
)
