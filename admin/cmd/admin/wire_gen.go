// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"admin/internal/conf"
	"admin/internal/server"
	"admin/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func initApp(confServer *conf.Server, data *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	adminService := service.NewAdminService(logger)
	httpServer := server.NewHTTPServer(confServer, adminService, logger)
	app := newApp(logger, httpServer)
	return app, func() {
	}, nil
}
