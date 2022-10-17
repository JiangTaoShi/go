package stream

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type ProducerOptions struct {
	RedisClient *redis.Client
}

func NewProducer(options *ProducerOptions) (*Producer, error) {
	return &Producer{
		options: options,
	}, nil
}

type Producer struct {
	options *ProducerOptions
}

func (p *Producer) Enqueue(msg *Message) error {
	ctx := context.Background()
	args := &redis.XAddArgs{
		ID:     msg.ID,
		Stream: msg.Stream,
		Values: msg.Values,
	}
	id, err := p.options.RedisClient.XAdd(ctx, args).Result()
	if err != nil {
		fmt.Println(err)
		return err
	}
	msg.ID = id
	return nil
}
