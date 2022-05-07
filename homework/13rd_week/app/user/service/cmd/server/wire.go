// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"demo/app/user/service/internal/biz"
	"demo/app/user/service/internal/conf"
	"demo/app/user/service/internal/data"
	"demo/app/user/service/internal/data/cache"
	"demo/app/user/service/internal/data/db"
	"demo/app/user/service/internal/server"
	"demo/app/user/service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, cache.ProviderSet, db.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
