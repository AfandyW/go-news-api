package tagsservice

import (
	"context"
	"errors"
	"go-news-api/domain/entities"
	"go-news-api/domain/tags"
)

type service struct {
	Repo tags.IRepository
}

func NewService(repo tags.IRepository) tags.IService {
	return &service{
		Repo: repo,
	}
}

func (serv *service) ListTags(ctx context.Context) (*[]entities.TagsDTO, error) {
	result, err := serv.Repo.List(ctx)

	tags := []entities.TagsDTO{}

	for _, v := range result {
		tag := v.ToDTO()
		tags = append(tags, tag)
	}

	// redis cache
	return &tags, err
}

func (serv *service) CreateNewTags(ctx context.Context, t *entities.Tags) (*entities.TagsDTO, error) {
	result, err := serv.Repo.Create(ctx, t)

	if err != nil {
		return nil, err
	}

	tags := result.ToDTO()

	return &tags, err
}
func (serv *service) UpdateTags(ctx context.Context, id uint64, name string) (*entities.TagsDTO, error) {
	tag, err := serv.Repo.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	if tag.ID == 0 {
		return nil, errors.New("Tags Not Found")
	}

	tag.Update(name)

	result, err := serv.Repo.Update(ctx, tag)
	tags := result.ToDTO()
	return &tags, err
}

func (serv *service) DeleteTags(ctx context.Context, id uint64) error {
	tag, err := serv.Repo.GetById(ctx, id)

	if err != nil {
		return err
	}

	if tag.ID == 0 {
		return errors.New("Tags Not Found")
	}

	err = serv.Repo.Delete(ctx, id)

	return err
}
