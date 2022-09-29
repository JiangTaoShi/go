package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
)

var RedisClient *redis.Client

func Setup() {
	RedisClient = redis.NewClient(&redis.Options{
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
}
