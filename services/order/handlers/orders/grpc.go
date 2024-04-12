package handler

import (
	"context"

	"google.golang.org/grpc"

	"github.com/chiatzeheng/pos/services/common/genproto/orders"
	"github.com/chiatzeheng/pos/services/order/types"
)

type OrdersGrpcHandler struct {
	//service injection
	//unimplmented  UnimplementedOrderSssssserviceServer
	ordersService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrdersService(grpc *grpc.Server, ordersService types.OrderService) {
	gRPCHandler := &OrdersGrpcHandler{
		ordersService: ordersService,
	}

	//register the OrderService Server
	orders.RegisterOrderServiceServer(grpc, gRPCHandler)
}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID:    42,
		CustomerID: 2,
		ProductID:  1,
		Quantity:   10,
	}

	err := h.ordersService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}
	return res, nil
}
