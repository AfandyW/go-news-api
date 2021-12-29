package config

import (
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	Host    string
	Db      int64
	Expires time.Duration
}

func InitClient(cache RedisCache) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.Host,
		Password: "",
		DB:       int(cache.Db),
	})
}
