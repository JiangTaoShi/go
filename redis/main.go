package main

import (
	"fmt"
	"github.com/JiangTaoShi/go/redis/stream"
)

func main() {
	Setup()
	producer, err := stream.NewProducer(&stream.ProducerOptions{
		RedisClient: RedisClient,
	})
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 10; i++ {
		producer.Enqueue(&stream.Message{
			Stream: "test-stream-01",
			Values: map[string]interface{}{
				"oid": i,
			},
		})
	}
}
