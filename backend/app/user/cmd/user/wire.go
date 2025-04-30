//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/soraQaQ/blog/app/user/internal/biz"
	"github.com/soraQaQ/blog/app/user/internal/conf"
	"github.com/soraQaQ/blog/app/user/internal/data"
	"github.com/soraQaQ/blog/app/user/internal/server"
	"github.com/soraQaQ/blog/app/user/internal/service"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Registry, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
