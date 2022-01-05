package config

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	Host    string
	Db      int64
	Expires int
	Port    string
}

func InitRedis(c RedisCache) *redis.Client {
	address := fmt.Sprintf(`%s:%s`, c.Host, c.Port)

	return redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       int(c.Db),
	})
}
