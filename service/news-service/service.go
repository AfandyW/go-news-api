package newsservice

import (
	"context"
	"errors"
	"fmt"
	"go-news-api/domain/entities"
	"go-news-api/domain/news"
	"go-news-api/domain/tags"
)

type service struct {
	Repo     news.IRepository
	TagsRepo tags.IRepository
	Cache    news.ICacheRepository
}

func NewService(repo news.IRepository, tagsrepo tags.IRepository, cache news.ICacheRepository) news.IService {
	return &service{
		Repo:     repo,
		TagsRepo: tagsrepo,
		Cache:    cache,
	}
}

func (serv *service) ListNews(ctx context.Context) ([]entities.NewsDTO, error) {
	keys := "ListNews"
	news, _ := serv.Cache.List(ctx, keys)

	if news == nil {
		result, err := serv.Repo.List(ctx)

		if err != nil {
			return nil, err
		}

		newsDto := []entities.NewsDTO{}

		for _, v := range result {
			new := v.ToDTO()
			newsDto = append(newsDto, new)
		}
		serv.Cache.Set(ctx, keys, newsDto)
		news = newsDto
	}

	return news, nil
}

func (serv *service) ListNewsByStatus(ctx context.Context, status string) ([]entities.NewsDTO, error) {
	keys := "ListNews" + status
	news, _ := serv.Cache.List(ctx, keys)

	if news == nil {
		result, err := serv.Repo.ListByStatus(ctx, status)

		if err != nil {
			return nil, err
		}

		newsDto := []entities.NewsDTO{}

		for _, v := range result {
			new := v.ToDTO()
			newsDto = append(newsDto, new)
		}
		serv.Cache.Set(ctx, keys, newsDto)
		news = newsDto
	}

	return news, nil
}

func (serv *service) ListNewsByTopic(ctx context.Context, topic string) ([]entities.TagsDTONews, error) {
	keys := "listNews" + topic
	tags, _ := serv.Cache.ListTopic(ctx, keys)

	if tags == nil {
		result, err := serv.TagsRepo.ListByTopic(ctx, topic)

		if err != nil {
			return nil, err
		}

		tagsDto := []entities.TagsDTONews{}

		for _, v := range result {
			tag := v.ToDTOWithNews()
			tagsDto = append(tagsDto, tag)
		}

		tags = tagsDto

		err = serv.Cache.Set(ctx, "listTags", tagsDto)
	}

	return tags, nil
}

func (serv *service) CreateNewNews(ctx context.Context, tags []string, name, status string) (*entities.NewsDTO, error) {
	rTags := []*entities.Tags{}

	for _, v := range tags {
		tag, err := serv.TagsRepo.GetByName(ctx, v)

		if err != nil {
			return nil, err
		}

		if tag.ID == 0 {
			return nil, errors.New("Tags Not Found")
		}

		rTags = append(rTags, tag)
	}

	news := &entities.News{
		Name:   name,
		Status: status,
		Tags:   rTags,
	}

	err := news.Validate()
	if err != nil {
		return nil, err
	}

	news, err = serv.Repo.Create(ctx, news)

	if err != nil {
		return nil, err
	}

	newsDto := news.ToDTO()
	serv.Cache.FlushAll(ctx)

	return &newsDto, err
}
func (serv *service) UpdateNews(ctx context.Context, id uint64, tags []string, name, status string) (*entities.NewsDTO, error) {
	keys := fmt.Sprintf("n%d", id)
	news, _ := serv.Cache.Get(ctx, keys)

	if news == nil {
		result, err := serv.Repo.GetById(ctx, id)

		if err != nil {
			return nil, err
		}

		if result.ID == 0 {
			return nil, errors.New("News Not Found")
		}

		serv.Cache.Set(ctx, keys, result)
		news = result
	}

	rTags := []*entities.Tags{}

	for _, v := range tags {
		tag, err := serv.TagsRepo.GetByName(ctx, v)
		if err != nil {
			return nil, err
		}

		if tag.ID == 0 {
			return nil, errors.New("Tags Not Found")
		}

		rTags = append(rTags, tag)
	}

	news.Update(rTags, name, status)

	news, err := serv.Repo.Update(ctx, id, news)

	if err != nil {
		return nil, err
	}

	newsDto := news.ToDTO()
	serv.Cache.FlushAll(ctx)

	return &newsDto, err
}

func (serv *service) DeleteNews(ctx context.Context, id uint64) error {
	keys := fmt.Sprintf("n%d", id)
	news, _ := serv.Cache.Get(ctx, keys)

	if news == nil {
		result, err := serv.Repo.GetById(ctx, id)

		if err != nil {
			return err
		}

		if result.ID == 0 {
			return errors.New("News Not Found")
		}
	}

	err := serv.Repo.Delete(ctx, id)

	if err != nil {
		return err
	}

	serv.Cache.FlushAll(ctx)

	return nil
}
