package redislib

import (
	"github.com/go-redis/redis"
	"github.com/gotechbook/application-notebook/config"
	"github.com/gotechbook/application-notebook/logger"
)

var (
	client *redis.Client
	err    error
)

func ExampleNewClient() {

	client = redis.NewClient(&redis.Options{
		Addr:         config.C.Redis.Addr,
		Password:     config.C.Redis.Password,
		DB:           config.C.Redis.Db,
		PoolSize:     config.C.Redis.PoolSize,
		MinIdleConns: config.C.Redis.MinIdleConns,
	})

	if _, err = client.Ping().Result(); err != nil {
		logger.Error("redis init error", err, nil)
		return
	}
	logger.Info("redis init success", nil)
}

func GetClient() (c *redis.Client) {
	return client
}
