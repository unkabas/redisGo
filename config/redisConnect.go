package config

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var ctx *context.Context
var rdb *redis.Client

func RedisConnect() {

	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})
	err := rdb.HSet(ctx, "city", "", "").Err()
	if err != nil {
		panic(err)
	}

}
