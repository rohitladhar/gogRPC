package main

import (
	"context"
	"log"
	"net"

	pb "github.com/rohitladhar/gogRPC/coffeeshop_proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCoffeeShopServer
}

func (s *server) GetMenu(menuRequest *pb.MenuRequest, srv pb.CoffeeShop_GetMenuServer) error {
	items := []*pb.Item {
		{
			Id:   "1",
			Name: "Black Coffee",
		},
		{
			Id:"2",
			Name:"Americano",
		},
		{
			Id:"3",
			Name:"Vanilla Latte",
		},
	}

	for i := range items {
		srv.Send(&pb.Menu{
			Items: items[0: i+1],
		})
	}
	return nil
}
func (s *server) PlaceOrder(context context.Context, order *pb.Order) (*pb.Receipt, error) {
	return &pb.Receipt{
		Id:"ABC123",
	}, nil
}
func (s *server) GetOrderStatus(context context.Context,receipt *pb.Receipt) (*pb.OrderStatus, error) {
	return &pb.OrderStatus{
		OrderId: receipt.Id,
		Status: "In Progress",
	},nil
}

func main(){
	lis, err := net.Listen("tcp",":9001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", lis)
	}

	grpcServer := grpc.NewServer();
	pb.RegisterCoffeeShopServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil{
		log.Fatalf("Faied to serve %s", err)
	}
}