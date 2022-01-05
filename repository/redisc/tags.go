package redisc

import (
	"context"
	"encoding/json"
	"go-news-api/domain/entities"
	"go-news-api/domain/tags"
	"time"

	"github.com/go-redis/redis/v8"
)

type tagsCache struct {
	Redis   *redis.Client
	Expires int
}

func NewRedisTagsCache(redis *redis.Client, redisExpires int) tags.ICacheRepository {
	return &tagsCache{
		Redis:   redis,
		Expires: redisExpires,
	}
}

func (c *tagsCache) Set(ctx context.Context, key string, value interface{}) error {

	json, err := json.Marshal(value)
	if err != nil {
		return err
	}

	c.Redis.Set(ctx, key, json, time.Duration(c.Expires)*time.Minute)

	return nil
}

func (c *tagsCache) List(ctx context.Context, key string) ([]entities.TagsDTO, error) {
	val, err := c.Redis.Get(ctx, key).Result()

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
	val, err := c.Redis.Get(ctx, key).Result()

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
	c.Redis.FlushAll(ctx)
}
