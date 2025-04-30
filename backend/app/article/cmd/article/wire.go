//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/soraQaQ/blog/app/article/internal/biz"
	"github.com/soraQaQ/blog/app/article/internal/conf"
	"github.com/soraQaQ/blog/app/article/internal/data"
	"github.com/soraQaQ/blog/app/article/internal/server"
	"github.com/soraQaQ/blog/app/article/internal/service"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger, *tracesdk.TracerProvider, *conf.Registry) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
