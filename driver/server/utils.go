package main

import (
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	REDIS_ADDR = os.Getenv("REDIS_ADDR")
)

func NewRedisClient() *redis.Client {
	if REDIS_ADDR == "" {
		REDIS_ADDR = "redis-17608.c228.us-central1-1.gce.cloud.redislabs.com:17608"
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     REDIS_ADDR,
		Password: "utq2bf7KKuEj8syRLcmyUeJUExMVvif3",
	})

	return rdb
}
