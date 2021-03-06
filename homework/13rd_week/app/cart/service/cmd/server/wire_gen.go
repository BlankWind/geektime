// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"demo/app/cart/service/internal/biz"
	"demo/app/cart/service/internal/conf"
	"demo/app/cart/service/internal/data"
	"demo/app/cart/service/internal/server"
	"demo/app/cart/service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger, tracerProvider *trace.TracerProvider) (*kratos.App, func(), error) {
	database := data.NewMongo(confData)
	dataData, cleanup, err := data.NewData(database, logger)
	if err != nil {
		return nil, nil, err
	}
	cartRepo := data.NewCartRepo(dataData, logger)
	cartUseCase := biz.NewCartUseCase(cartRepo, logger)
	cartService := service.NewCartService(cartUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, logger, tracerProvider, cartService)
	httpServer := server.NewHTTPServer(confServer, cartService, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, grpcServer, httpServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
