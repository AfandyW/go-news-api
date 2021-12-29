package newsservice

import (
	"context"
	"errors"
	"go-news-api/domain/entities"
	"go-news-api/domain/news"
	"go-news-api/domain/tags"
)

type service struct {
	Repo     news.IRepository
	TagsRepo tags.IRepository
}

func NewService(repo news.IRepository, tagsrepo tags.IRepository) news.IService {
	return &service{
		Repo:     repo,
		TagsRepo: tagsrepo,
	}
}

func (serv *service) ListNews(ctx context.Context) ([]entities.NewsDTO, error) {
	result, err := serv.Repo.List(ctx)

	if err != nil {
		return nil, err
	}

	news := []entities.NewsDTO{}

	for _, v := range result {
		new := v.ToDTO()
		news = append(news, new)
	}

	return news, nil
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

	news, err := serv.Repo.Create(ctx, news)

	if err != nil {
		return nil, err
	}

	newsDto := news.ToDTO()

	return &newsDto, err
}
func (serv *service) UpdateNews(ctx context.Context, id uint64, tags []string, name, status string) (*entities.NewsDTO, error) {
	result, err := serv.Repo.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	if result.ID == 0 {
		return nil, errors.New("News Not Found")
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

	result.Update(rTags, name, status)

	news, err := serv.Repo.Update(ctx, id, result)

	if err != nil {
		return nil, err
	}

	newsDto := news.ToDTO()

	return &newsDto, err
}

func (serv *service) DeleteNews(ctx context.Context, id uint64) error {
	result, err := serv.Repo.GetById(ctx, id)

	if err != nil {
		return err
	}

	if result.ID == 0 {
		return errors.New("News Not Found")
	}

	err = serv.Repo.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
