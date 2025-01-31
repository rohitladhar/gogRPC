package main

import (
	"context"
	"encoding/json"
	"fmt"

	"io"
	"log"

	notificationservice "github.com/rohitladhar/gogRPC/golangRedis/notificaionservice"
	"github.com/rohitladhar/gogRPC/golangRedis/notificaionservice/notificationproto"
)

func main(){
	client, err := notificationservice.NewClients()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	stream, err := client.GetNotification(ctx, &notificationproto.NotificationRequest{
		UserId: "123",
	})
	
	if err != nil {
		log.Fatal(err)
	}
	log.Println(stream.Recv())
	for {
		notication, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to read notifications: %v",err)
		}
		
		b,err := json.MarshalIndent(notication,"","\t")
		if err != nil{
			log.Fatalf("failed to execute:%v",err)
		}
		fmt.Println(string(b))
	}
}