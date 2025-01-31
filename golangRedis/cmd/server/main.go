package main

import (
	"context"
	"log"
	"log/slog"
	"net"

	grpcstreaming "github.com/rohitladhar/gogRPC/golangRedis"
	notificationservice "github.com/rohitladhar/gogRPC/golangRedis/notificaionservice"
	"github.com/rohitladhar/gogRPC/golangRedis/notificaionservice/notificationproto"
	"google.golang.org/grpc"
)
func main(){
	lis, err := net.Listen("tcp", notificationservice.Address)
	if err != nil {
		log.Fatalf("Failed to listen :%v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	ctx := context.Background()
	redisClient := grpcstreaming.NewRedisClient(ctx)
	handler := notificationservice.NewHandler(redisClient)
	notificationproto.RegisterNotificationServiceServer(grpcServer,handler)
	slog.Info("Listening on "+ notificationservice.Address)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}