package types

import (
	"context"

	"github.com/rohitladhar/gogRPC/kgg/services/common/genproto/orders"
)

type OrderService interface{
	CreateOrder(context.Context, *orders.Order ) error
}