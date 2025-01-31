package main

import (
	"context"
	"fmt"
	"log"

	grpcstreaming "github.com/rohitladhar/gogRPC/golangRedis"
)

func main() {
	ctx := context.Background()
	redisClient := grpcstreaming.NewRedisClient(ctx)
	channelName := fmt.Sprintf("notification/%s", "123")
	cmd := redisClient.Publish(ctx, channelName, "Hello World")
	if err := cmd.Err(); err != nil {
		log.Fatalf("Error publishing message to Redis: %v", err)
	}
	fmt.Printf("Message published to channel %s\n", channelName)
}
