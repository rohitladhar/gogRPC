package notificationservice

import (
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rohitladhar/gogRPC/golangRedis/notificaionservice/notificationproto"
	"google.golang.org/grpc"
)

var _ notificationproto.NotificationServiceServer = (*Handler)(nil)

type Handler struct {
	notificationproto.UnimplementedNotificationServiceServer
	redisClient *redis.Client
}

func NewHandler(redisClient *redis.Client) *Handler {
	return &Handler{
		redisClient: redisClient,
	}
}

func (h *Handler) GetNotification(req *notificationproto.NotificationRequest,stream grpc.ServerStreamingServer[notificationproto.Notification]) error {
	pubsub := h.redisClient.Subscribe(stream.Context(), fmt.Sprintf("notification/ %s", req.GetUserId()))
	for {
		select{
		case <-stream.Context().Done():
			return stream.Context().Err()
		case msg := <-pubsub.Channel():
			if err := stream.Send(&notificationproto.Notification{
				UserId: req.GetUserId(),
				Content: fmt.Sprintf("New Notification %s: %s", time.Now().String(),msg.Payload),
				CreatedAt: time.Now().UnixMilli(),
			});
			err != nil {
				return fmt.Errorf("Could Not Send notification: %w", err)
			}
		}
	}
}