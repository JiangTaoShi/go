package main

import (
	"context"
	"fmt"
	"github.com/JiangTaoShi/go/redis/stream"
	"github.com/go-redis/redis/v8"
)

func main() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", "39.99.174.39", "7379"),
		Password: "pay.Media@2020#redis",
		DB:       2,
		PoolSize: 100,
	})
	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
	}

	producer, err := stream.NewProducer(&stream.ProducerOptions{
		RedisClient: RedisClient,
	})
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 30; i++ {
		producer.Enqueue(&stream.Message{
			Stream: "test-stream-01",
			Values: map[string]interface{}{
				"oid": i,
			},
		})
	}
}
