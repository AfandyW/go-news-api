package redisc

import (
	"context"
	"encoding/json"
	"fmt"
	"go-news-api/config"
	"go-news-api/domain/entities"
	"go-news-api/domain/news"
	"time"

	"github.com/go-redis/redis/v8"
)

type newsCache struct {
	Host    string
	Db      int64
	Expires int
	Port    string
}

func NewRedisCach(redis config.RedisCache) news.ICacheRepository {
	return &newsCache{
		Host:    redis.Host,
		Db:      redis.Db,
		Expires: redis.Expires,
		Port:    redis.Port,
	}
}

func (c *newsCache) getClientN() *redis.Client {
	address := fmt.Sprintf(`%s:%s`, c.Host, c.Port)

	return redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       int(c.Db),
	})
}

func (c *newsCache) Set(ctx context.Context, key string, value interface{}) error {
	client := c.getClientN()

	json, err := json.Marshal(value)
	if err != nil {
		return err
	}

	client.Set(ctx, key, json, time.Duration(c.Expires)*time.Minute)

	return nil
}

func (c *newsCache) List(ctx context.Context, key string) ([]entities.NewsDTO, error) {
	client := c.getClientN()

	val, err := client.Get(ctx, key).Result()

	if err != nil {
		return nil, err
	}

	news := []entities.NewsDTO{}

	err = json.Unmarshal([]byte(val), &news)
	if err != nil {
		return nil, err
	}

	return news, nil
}

func (c *newsCache) Get(ctx context.Context, key string) (*entities.News, error) {
	client := c.getClientN()

	val, err := client.Get(ctx, key).Result()

	if err != nil {
		return nil, err
	}

	news := entities.News{}

	err = json.Unmarshal([]byte(val), &news)
	if err != nil {
		return nil, err
	}

	return &news, nil
}

func (c *newsCache) FlushAll(ctx context.Context) {
	client := c.getClientN()

	client.FlushAll(ctx)
}
