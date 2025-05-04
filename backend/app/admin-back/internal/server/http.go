package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwt2 "github.com/golang-jwt/jwt/v5"
	v1 "github.com/soraQaQ/blog/api/admin/v1"
	"github.com/soraQaQ/blog/app/admin/internal/conf"
	"github.com/soraQaQ/blog/app/admin/internal/service"
	myMiddleware "github.com/soraQaQ/blog/pkg/middleware"
	"github.com/soraQaQ/blog/pkg/response"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

func TraceIDMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			// 获取 TraceID
			spanCtx := trace.SpanContextFromContext(ctx)
			if spanCtx.IsValid() {
				traceID := spanCtx.TraceID().String()

				// 将 TraceID 设置到响应头中
				if tr, ok := transport.FromServerContext(ctx); ok {
					if ht, ok := tr.(http.Transporter); ok {
						ht.ReplyHeader().Set("X-Trace-ID", traceID)
						ht.Request().Context()
					}

				}
			}

			// 继续处理请求
			return handler(ctx, req)
		}
	}
}

func NewMiddleware(ac *conf.Auth, logger log.Logger, tp *tracesdk.TracerProvider) http.ServerOption {
	whiteList := make(map[string]struct{})
	whiteList["/admin.v1.AdminService/Register"] = struct{}{}
	whiteList["/admin.v1.AdminService/Login"] = struct{}{}
	return http.Middleware(
		recovery.Recovery(),
		tracing.Server(tracing.WithTracerProvider(tp)),
		logging.Server(logger),
		TraceIDMiddleware(),
		selector.Server(
			jwt.Server(func(token *jwt2.Token) (interface{}, error) {
				return []byte(ac.ApiKey), nil
			},
				jwt.WithSigningMethod(jwt2.SigningMethodHS256),
			),
		).Match(myMiddleware.NewWhiteListMatcher(whiteList)).Build(),
	)
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, ac *conf.Auth, admin *service.AdminService, logger log.Logger, tp *tracesdk.TracerProvider) *http.Server {
	var opts = []http.ServerOption{
		NewMiddleware(ac, logger, tp),
		http.ResponseEncoder(response.EncoderResponse()),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterAdminServiceHTTPServer(srv, admin)
	return srv
}
