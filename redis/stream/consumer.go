package stream

import (
	"context"
	"github.com/go-redis/redis/v9"
	"log"
)

type ConsumerOptions struct {
	RedisClient *redis.Client
	Stream      string
	GroupName   string
	Consumer    string
	Start       string
}

func NewConsumer(options *ConsumerOptions) (*Consumer, error) {
	return &Consumer{
		options: options,
	}, nil
}

type Consumer struct {
	options *ConsumerOptions
}

//拉取消息
func (c *Consumer) Poll() ([]redis.XStream, error) {
	option := c.options
	ctx := context.Background()
	return option.RedisClient.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    option.GroupName,
		Consumer: option.Consumer,
		Streams:  []string{option.Stream, ">"},
		Count:    1,
		Block:    0,
		NoAck:    false,
	}).Result()
}

//确认消息
func (c *Consumer) Ack(messageId string) (int64, error) {
	option := c.options
	ctx := context.Background()
	return option.RedisClient.XAck(ctx, option.Stream, option.GroupName, messageId).Result()
}

//创建
func (c *Consumer) CreateGroupMkStream() {
	option := c.options
	ctx := context.Background()
	err := option.RedisClient.XGroupCreateMkStream(ctx,
		option.Stream,
		option.GroupName,
		option.Start).Err()
	if err != nil && err.Error() != "BUSYGROUP Consumer Group name already exists" {
		log.Println(err)
	}
}
