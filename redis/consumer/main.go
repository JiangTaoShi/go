package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/JiangTaoShi/go/redis/stream"
	"github.com/go-redis/redis/v8"
	"log"
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

	b := make([]byte, 16)
	_, err = rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	fmt.Println("start success" + uuid)

	consumer, err := stream.NewConsumer(&stream.ConsumerOptions{
		RedisClient: RedisClient,
		Stream:      "CustomerStream6",
		GroupName:   "CustomerStreamGroup",
		Consumer:    uuid,
		Start:       "0",
	})
	if err != nil {
		fmt.Println(err)
	}
	consumer.CreateGroupMkStream()
	for {
		ctx := context.Background()
		entities, err := consumer.Poll(ctx)
		if err != nil {
			fmt.Println(err)
		}
		//TODO
		if entities != nil {
			for i := 0; i < len(entities[0].Messages); i++ {
				messageId := entities[0].Messages[i].ID
				values := entities[0].Messages[i].Values
				fmt.Println(values)
				data := values["222"]
				fmt.Println(fmt.Sprintf("%+v", data))
				//ACK
				consumer.Ack(ctx, messageId)
			}
		}
	}
}
