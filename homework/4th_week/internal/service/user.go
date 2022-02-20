package service

import (
	"fmt"
	"homework04/internal/biz"
)

type UserService struct {
	biz biz.UserBiz
}

func NewUserService(b biz.UserBiz) UserService {
	return UserService{biz: b}
}

func (us *UserService) Start() string {
	return fmt.Sprintf("Server Start! ->%s", us.biz.GetUser())
}
