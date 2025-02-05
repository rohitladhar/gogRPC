package handler

import (
	"context"

	"github.com/rohitladhar/gogRPC/kgg/services/common/genproto/orders"
	"github.com/rohitladhar/gogRPC/kgg/services/orders/types"
	"google.golang.org/grpc"
)

type OrdersGrpcHandler struct {
	ordersService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrdersService(grpc *grpc.Server, ordersService types.OrderService) {
	gRPCHandler := &OrdersGrpcHandler{
		ordersService: ordersService,
	}
	orders.RegisterOrderServiceServer(grpc,gRPCHandler)
}

func (h *OrdersGrpcHandler) GetOrders(ctx context.Context,req *orders.GetOrdersRequest) (*orders.GetOrderResponse,error){
	os := h.ordersService.GetOrders(ctx)
	res := &orders.GetOrderResponse{
		Orders:os,
	}
	return res, nil
}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context,req *orders.CreateOrderRequest) (*orders.CreateOrderResponse,error){
	order := &orders.Order{
		OrderId: 33,
		CustomerId: 33,
		ProductID: 1,
		Quantity: 3,
	}

	err := h.ordersService.CreateOrder(ctx, order)
	if err!= nil {
		return nil,err
	}
	res := &orders.CreateOrderResponse{
		Status:"success",
	}
	return res , nil;
}