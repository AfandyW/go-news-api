package tagsservice

import (
	"context"
	"errors"
	"fmt"
	"go-news-api/domain/entities"
	"go-news-api/domain/tags"
)

type service struct {
	Repo  tags.IRepository
	Cache tags.ICacheRepository
}

func NewService(repo tags.IRepository, cache tags.ICacheRepository) tags.IService {
	return &service{
		Repo:  repo,
		Cache: cache,
	}
}

func (serv *service) ListTags(ctx context.Context) (*[]entities.TagsDTO, error) {

	tags, _ := serv.Cache.List(ctx, "listTags")

	if tags == nil {
		fmt.Println("ambil dari db")
		result, err := serv.Repo.List(ctx)

		if err != nil {
			return nil, err
		}

		tagsDto := []entities.TagsDTO{}

		for _, v := range result {
			tag := v.ToDTO()
			tagsDto = append(tagsDto, tag)
		}

		tags = tagsDto

		err = serv.Cache.Set(ctx, "listTags", tagsDto)
	}

	return &tags, nil
}

func (serv *service) CreateNewTags(ctx context.Context, t *entities.Tags) (*entities.TagsDTO, error) {
	result, err := serv.Repo.Create(ctx, t)

	if err != nil {
		return nil, err
	}

	tags := result.ToDTO()

	serv.Cache.FlushAll(ctx)
	return &tags, err
}
func (serv *service) UpdateTags(ctx context.Context, id uint64, name string) (*entities.TagsDTO, error) {
	keys := fmt.Sprintf("t%d", id)
	tag, _ := serv.Cache.Get(ctx, keys)

	if tag == nil {

		result, err := serv.Repo.GetById(ctx, id)

		if err != nil {
			return nil, err
		}

		if result.ID == 0 {
			return nil, errors.New("Tags Not Found")
		}

		serv.Cache.Set(ctx, keys, result)
		tag = result
	}

	tag.Update(name)

	result, err := serv.Repo.Update(ctx, tag)
	tags := result.ToDTO()

	serv.Cache.FlushAll(ctx)
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

	serv.Cache.FlushAll(ctx)

	return err
}
