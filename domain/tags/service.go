package tags

import (
	"context"
	"go-news-api/domain/entities"
)

type IService interface {
	ListTags(ctx context.Context) (*[]entities.TagsDTO, error)
	CreateNewTags(ctx context.Context, tags *entities.Tags) (*entities.TagsDTO, error)
	UpdateTags(ctx context.Context, id uint64, name string) (*entities.TagsDTO, error)
	DeleteTags(ctx context.Context, id uint64) error
}