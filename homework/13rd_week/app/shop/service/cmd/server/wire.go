// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"demo/app/shop/service/internal/biz"
	"demo/app/shop/service/internal/conf"
	"demo/app/shop/service/internal/data"
	"demo/app/shop/service/internal/server"
	"demo/app/shop/service/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, *conf.Auth, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
