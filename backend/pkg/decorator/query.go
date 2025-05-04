package decorator

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type QueryHandler[Q, R any] interface {
	Handler(ctx context.Context, query Q) (R, error)
}

func ApplyQueryDecorator[Q, R any](handler QueryHandler[Q, R], logger log.Logger) QueryHandler[Q, R] {
	return &queryLoggingDecorator[Q, R]{
		log:  log.NewHelper(logger),
		base: handler,
	}
}
