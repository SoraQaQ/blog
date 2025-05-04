package decorator

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type CommandHandler[C any] interface {
	Handler(ctx context.Context, cmd C) error
}

func ApplyCommandDecorators[C any](handler CommandHandler[C], logger log.Logger) CommandHandler[C] {
	return &commandLoggingDecorator[C]{
		log:  log.NewHelper(logger),
		base: handler,
	}
}
