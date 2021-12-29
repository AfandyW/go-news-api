package redisc

import (
	"context"
	"encoding/json"
	"fmt"
	"go-news-api/config"
	"go-news-api/domain/entities"
	"go-news-api/domain/tags"
	"time"

	"github.com/go-redis/redis/v8"
)

type tagsCache struct {
	Host    string
	Db      int64
	Expires int
	Port    string
}

func NewRedisCache(redis config.RedisCache) tags.ICacheRepository {
	return &tagsCache{
		Host:    redis.Host,
		Db:      redis.Db,
		Expires: redis.Expires,
		Port:    redis.Port,
	}
}

func (c *tagsCache) getClient() *redis.Client {
	address := fmt.Sprintf(`%s:%s`, c.Host, c.Port)

	return redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       int(c.Db),
	})
}

func (c *tagsCache) Set(ctx context.Context, key string, value interface{}) error {
	client := c.getClient()

	json, err := json.Marshal(value)
	if err != nil {
		return err
	}

	client.Set(ctx, key, json, time.Duration(c.Expires)*time.Minute)

	return nil
}

func (c *tagsCache) List(ctx context.Context, key string) ([]entities.TagsDTO, error) {
	client := c.getClient()

	val, err := client.Get(ctx, key).Result()

	if err != nil {
		return nil, err
	}

	tags := []entities.TagsDTO{}

	err = json.Unmarshal([]byte(val), &tags)
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (c *tagsCache) Get(ctx context.Context, key string) (*entities.Tags, error) {
	client := c.getClient()

	val, err := client.Get(ctx, key).Result()

	if err != nil {
		return nil, err
	}

	tags := entities.Tags{}

	err = json.Unmarshal([]byte(val), &tags)
	if err != nil {
		return nil, err
	}

	return &tags, nil
}

func (c *tagsCache) FlushAll(ctx context.Context) {
	client := c.getClient()

	client.FlushAll(ctx)
}
