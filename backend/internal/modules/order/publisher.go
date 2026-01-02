package order

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type EventPublisher interface {
	Publish(ctx context.Context, event *OrderEvent) error
}

type RedisPublisher struct {
	client *redis.Client
}

func NewRedisPublisher(c *redis.Client) *RedisPublisher {
	return &RedisPublisher{client: c}
}

func (p *RedisPublisher) Publish(
	ctx context.Context,
	event *OrderEvent,
) error {

	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return p.client.Publish(ctx, "order-events", data).Err()
}
