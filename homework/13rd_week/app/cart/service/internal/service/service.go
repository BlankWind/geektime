package service

import (
	v1 "demo/api/cart/v1"
	"demo/app/cart/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewCartService)

type CartService struct {
	v1.UnimplementedCartServer

	cc  *biz.CartUseCase
	log *log.Helper
}

func NewCartService(cc *biz.CartUseCase, logger log.Logger) *CartService {
	return &CartService{
		cc:  cc,
		log: log.NewHelper(log.With(logger, "module", "service/cart"))}
}
