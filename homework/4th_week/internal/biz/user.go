package biz

import (
	"fmt"
	"homework04/internal/data"
)

type UserBiz struct {
	dao data.User
}

func NewUserBiz(d data.User) UserBiz {
	return UserBiz{dao: d}
}

func (ub *UserBiz) GetUser() string {
	return fmt.Sprintf("name:%s", ub.dao.GetUserName())
}
