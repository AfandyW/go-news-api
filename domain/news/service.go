package news

import (
	"context"
	"go-news-api/domain/entities"
)

type IService interface {
	ListNews(ctx context.Context) ([]entities.NewsDTO, error)
	CreateNewNews(ctx context.Context, tags []string, name, status string) (*entities.NewsDTO, error)
	UpdateNews(ctx context.Context, id uint64, tags []string, name, status string) (*entities.NewsDTO, error)
	DeleteNews(ctx context.Context, id uint64) error
}
