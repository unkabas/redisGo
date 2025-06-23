package config

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func RedisConnect() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})
	var err error
	if err != nil {
		fmt.Println("Something wrong with ", ctx, "or", rdb)
	}

}
