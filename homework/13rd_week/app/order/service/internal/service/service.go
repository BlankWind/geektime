package service

import (
	v1 "demo/api/order/v1"
	"demo/app/order/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewOrderService)

type OrderService struct {
	v1.UnimplementedOrderServer

	oc  *biz.OrderUseCase
	log *log.Helper
}

func NewOrderService(oc *biz.OrderUseCase, logger log.Logger) *OrderService {
	return &OrderService{

		oc:  oc,
		log: log.NewHelper(log.With(logger, "module", "service/order"))}
}
