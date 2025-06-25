package database

import (
	"fmt"
	"strconv"

	"go-echo-clean-architecture/internal/models/config"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	rdb redis.Client
}

func NewRedisClient(config config.RedisConfig) redis.Client {
	rdb := *redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, strconv.Itoa(config.Port)),
		Password: config.Password,
		DB:       config.DbIndex,
	})

	return rdb
}
