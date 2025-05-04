package biz

import (
	"github.com/google/wire"
	"github.com/soraQaQ/blog/app/user/internal/biz/command"
	"github.com/soraQaQ/blog/app/user/internal/biz/query"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	command.NewCreateUserHandler,
	command.NewUpdateUserHandler,
	query.NewGetAllUserHandler,
	query.NewGetUserHandler,
	query.NewGetUserByEmailHandler,
)
