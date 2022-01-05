package redisc

import (
	"context"
	"encoding/json"
	"go-news-api/domain/entities"
	"go-news-api/domain/news"
	"time"

	"github.com/go-redis/redis/v8"
)

type newsCache struct {
	Redis   *redis.Client
	Expires int
}

func NewRedisNewsCache(redis *redis.Client, redisExpires int) news.ICacheRepository {
	return &newsCache{
		Redis:   redis,
		Expires: redisExpires,
	}
}

func (c *newsCache) Set(ctx context.Context, key string, value interface{}) error {
	json, err := json.Marshal(value)
	if err != nil {
		return err
	}

	c.Redis.Set(ctx, key, json, time.Duration(c.Expires)*time.Minute)

	return nil
}

func (c *newsCache) List(ctx context.Context, key string) ([]entities.NewsDTO, error) {
	val, err := c.Redis.Get(ctx, key).Result()

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

func (c *newsCache) ListTopic(ctx context.Context, key string) ([]entities.TagsDTONews, error) {
	val, err := c.Redis.Get(ctx, key).Result()

	if err != nil {
		return nil, err
	}

	news := []entities.TagsDTONews{}

	err = json.Unmarshal([]byte(val), &news)
	if err != nil {
		return nil, err
	}

	return news, nil
}

func (c *newsCache) Get(ctx context.Context, key string) (*entities.News, error) {
	val, err := c.Redis.Get(ctx, key).Result()

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
	c.Redis.FlushAll(ctx)
}
