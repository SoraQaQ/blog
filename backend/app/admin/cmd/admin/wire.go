//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/soraQaQ/blog/app/admin/internal/biz"
	"github.com/soraQaQ/blog/app/admin/internal/conf"
	"github.com/soraQaQ/blog/app/admin/internal/data"
	"github.com/soraQaQ/blog/app/admin/internal/server"
	"github.com/soraQaQ/blog/app/admin/internal/service"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

// wireApp init kratos application.
func initApp(*conf.Server, *conf.Data, *conf.Auth, log.Logger, *tracesdk.TracerProvider, *conf.Registry) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, data.ProviderSet, biz.ProviderSet, newApp))
}
