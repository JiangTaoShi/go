package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/JiangTaoShi/go/redis/stream"
	"github.com/go-redis/redis/v8"
)

func main() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", "", ""),
		Password: "",
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

	mapPush := make(map[string]interface{})
	mapPush["111"] = "test"
	str, _ := json.Marshal([]string{"zhangsan", "lisi"})
	mapPush["222"] = str
	for i := 0; i < 1; i++ {
		producer.Enqueue(&stream.Message{
			Stream: "CustomerStream6",
			Values: mapPush,
		})
	}
}
