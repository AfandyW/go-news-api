package news

import (
	"context"
	"go-news-api/domain/entities"
	"go-news-api/domain/news"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) news.IRepository {
	return &repository{
		DB: db,
	}
}

func (repo *repository) List(ctx context.Context) ([]*entities.News, error) {
	var news []*entities.News
	err := repo.DB.Debug().Order("created_at desc").Preload("Tags").Find(&news).Error

	if err != nil {
		return nil, err
	}

	return news, nil
}

func (repo *repository) ListByStatus(ctx context.Context, status string) ([]*entities.News, error) {
	var news []*entities.News
	err := repo.DB.Debug().Where("status = ?", status).Order("created_at desc").Preload("Tags").Find(&news).Error

	if err != nil {
		return nil, err
	}

	return news, nil
}

func (repo *repository) GetById(ctx context.Context, id uint64) (*entities.News, error) {
	var news *entities.News
	err := repo.DB.Debug().Where("id = ?", id).Find(&news).Error

	if err != nil {
		return nil, err
	}

	return news, nil
}

func (repo *repository) Create(ctx context.Context, news *entities.News) (*entities.News, error) {
	err := repo.DB.Debug().Create(&news).Error

	if err != nil {
		return nil, err
	}

	return news, nil
}

func (repo *repository) Update(ctx context.Context, id uint64, news *entities.News) (*entities.News, error) {
	err := repo.DB.Debug().Save(&news).Error

	if err != nil {
		return nil, err
	}

	return news, nil
}

func (repo *repository) Delete(ctx context.Context, id uint64) error {
	var news *entities.News
	err := repo.DB.Debug().Where("id = ?", id).Delete(&news).Error

	if err != nil {
		return err
	}

	return nil
}
