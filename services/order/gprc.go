package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	handler "github.com/chiatzeheng/pos/services/order/handlers/orders"
	"github.com/chiatzeheng/pos/services/order/service"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {

	//Send a new TCP connection
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	//register gprc services
	orderService := service.NewOrderService()
	handler.NewGrpcOrdersService(grpcServer, orderService)

	gRPCServer := grpc.NewServer()

	return gRPCServer.Serve(lis)
}
