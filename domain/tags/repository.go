package tags

import (
	"context"
	"go-news-api/domain/entities"
)

type IRepository interface {
	List(ctx context.Context) ([]*entities.Tags, error)
	GetById(ctx context.Context, id uint64) (*entities.Tags, error)
	GetByName(ctx context.Context, name string) (*entities.Tags, error)
	Create(ctx context.Context, tags *entities.Tags) (*entities.Tags, error)
	Update(ctx context.Context, tags *entities.Tags) (*entities.Tags, error)
	Delete(ctx context.Context, id uint64) error
}
