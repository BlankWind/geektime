//+build wireinject

package main

import (
	"homework04/internal/biz"
	"homework04/internal/data"
	"homework04/internal/service"

	"github.com/google/wire"
)

func InitUserService(usn string) service.UserService {
	wire.Build(data.NewUserDao, service.NewUserService, biz.NewUserBiz)
	return service.UserService{}
}
