package config

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var Rdb *redis.Client

func RedisConnect() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})
	pong, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(pong)
}
