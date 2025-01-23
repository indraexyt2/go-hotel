package helpers

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func SetupRedis() {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"127.0.0.1:6371",
			"127.0.0.1:6372",
			"127.0.0.1:6373",
			"127.0.0.1:6374",
			"127.0.0.1:6375",
			"127.0.0.1:6376",
		},
		ClusterSlots: func(ctx context.Context) ([]redis.ClusterSlot, error) {
			slots := []redis.ClusterSlot{
				{
					Start: 0,
					End:   5460,
					Nodes: []redis.ClusterNode{
						{Addr: "127.0.0.1:6371"},
						{Addr: "127.0.0.1:6375"},
					},
				},
				{
					Start: 5461,
					End:   10922,
					Nodes: []redis.ClusterNode{
						{Addr: "127.0.0.1:6372"},
						{Addr: "127.0.0.1:6376"},
					},
				},
				{
					Start: 10923,
					End:   16383,
					Nodes: []redis.ClusterNode{
						{Addr: "127.0.0.1:6373"},
						{Addr: "127.0.0.1:6374"},
					},
				},
			}
			return slots, nil
		},
	})

	result, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		Logger.Error("Failed to connect to redis: ", err)
		return
	}

	Logger.Info("Redis connected: PING!! |", result)
}
