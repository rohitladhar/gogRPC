package service

import (
	"context"

	"github.com/rohitladhar/gogRPC/kgg/services/common/genproto/orders"
)

type OrderService struct{

}
var ordersDb = make([]*orders.Order,0)

func NewOrderService () *OrderService{
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context,order *orders.Order) error{
	ordersDb = append(ordersDb, order)
	return nil
}

func(s *OrderService) GetOrders(ctx context.Context) []*orders.Order{
	return ordersDb
}