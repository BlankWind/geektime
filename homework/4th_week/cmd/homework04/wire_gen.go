// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"homework04/internal/biz"
	"homework04/internal/data"
	"homework04/internal/service"
)

// Injectors from wire.go:

func InitUserService(usn string) service.UserService {
	user := data.NewUserDao(usn)
	userBiz := biz.NewUserBiz(user)
	userService := service.NewUserService(userBiz)
	return userService
}
