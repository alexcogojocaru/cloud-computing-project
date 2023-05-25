package service

import (
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	REDIS_ADDR = os.Getenv("REDIS_ADDR")
)

func NewRedisClient() *redis.Client {
	if REDIS_ADDR == "" {
		REDIS_ADDR = "localhost:6379"
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     REDIS_ADDR,
		Password: "",
	})

	return rdb
}