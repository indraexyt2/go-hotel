package helpers

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
	"strings"
)

var RedisClient *redis.ClusterClient

func SetupRedis() {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: strings.Split(os.Getenv("REDIS_HOST"), ","),
	})

	result, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		Logger.Error("Failed to connect to redis: ", err)
		return
	}

	RedisClient = rdb

	Logger.Info("Redis connected: PING!! |", result)
}
