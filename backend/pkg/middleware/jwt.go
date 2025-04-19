package middleware

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/soraQaQ/blog/pkg/auth"
)

var (
	ErrMissingToken = errors.Unauthorized("UNAUTHORIZED", "missing jwt token")
	ErrInvalidToken = errors.Unauthorized("UNAUTHORIZED", "invalid token")
)

func JWTAuth(jwt *auth.JWT, opts ...Option) middleware.Middleware {
	o := &options{
		skipPaths: make(map[string]bool),
	}
	for _, opt := range opts {
		opt(o)
	}

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				operation := tr.Operation()
				if o.skipPaths[operation] {
					return handler(ctx, req)
				}

				auths := strings.SplitN(tr.RequestHeader().Get("Authorization"), " ", 2)
				if len(auths) != 2 || !strings.EqualFold(auths[0], "Bearer") {
					return nil, ErrMissingToken
				}
				log.Infof("get token %s", auths[1])
				claims, err := jwt.ParseToken(auths[1])
				if err != nil {
					return nil, ErrInvalidToken
				}
				ctx = context.WithValue(ctx, "user_id", claims.UserID)
			}
			return handler(ctx, req)
		}
	}
}

type options struct {
	skipPaths map[string]bool
}

type Option func(*options)

func WithSkipPaths(paths []string) Option {
	return func(o *options) {
		for _, path := range paths {
			o.skipPaths[path] = true
		}
	}
}
