package news

import (
	"context"
	"go-news-api/domain/entities"
)

type IRepository interface {
	List(ctx context.Context) ([]*entities.News, error)
	GetById(ctx context.Context, id uint64) (*entities.News, error)
	Create(ctx context.Context, news *entities.News) (*entities.News, error)
	Update(ctx context.Context, id uint64, news *entities.News) (*entities.News, error)
	Delete(ctx context.Context, id uint64) error
}

type ICacheRepository interface {
	Set(ctx context.Context, key string, value interface{}) error
	List(ctx context.Context, key string) ([]entities.NewsDTO, error)
	Get(ctx context.Context, key string) (*entities.News, error)
	FlushAll(ctx context.Context)
}
