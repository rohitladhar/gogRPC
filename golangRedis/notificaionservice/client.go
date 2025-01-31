package notificationservice

import (
	"fmt"
	"log"

	"github.com/rohitladhar/gogRPC/golangRedis/notificaionservice/notificationproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClients() (notificationproto.NotificationServiceClient, error) {
	// Create a background context (you can replace this with your own if needed)
	

	// Define gRPC dial options
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	// Establish the gRPC connection using DialContext
	conn, err := grpc.NewClient(Address, opts...)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to %v: %v", Address, err)
	}

	log.Println(notificationproto.NewNotificationServiceClient(conn))
	// Return the notification service client created using the connection
	return notificationproto.NewNotificationServiceClient(conn), nil
}
