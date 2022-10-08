package stream

import (
	"context"
	"github.com/go-redis/redis/v8"
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
func (c *Consumer) Poll(context context.Context) ([]redis.XStream, error) {
	option := c.options
	return option.RedisClient.XReadGroup(context, &redis.XReadGroupArgs{
		Group:    option.GroupName,
		Consumer: option.Consumer,
		Streams:  []string{option.Stream, ">"},
		Count:    1,
		Block:    0,
		NoAck:    false,
	}).Result()
}

//确认消息
func (c *Consumer) Ack(context context.Context, messageId string) (int64, error) {
	option := c.options
	return option.RedisClient.XAck(context, option.Stream, option.GroupName, messageId).Result()
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
