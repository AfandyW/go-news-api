package tagsrepo

import (
	"context"

	"gorm.io/gorm"

	"go-news-api/domain/entities"
	"go-news-api/domain/tags"
)

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) tags.IRepository {
	return &repository{
		DB: db,
	}
}

func (repo *repository) List(ctx context.Context) ([]*entities.Tags, error) {
	var tags []*entities.Tags
	err := repo.DB.Debug().Find(&tags).Error

	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (repo *repository) GetById(ctx context.Context, id uint64) (*entities.Tags, error) {
	var tags *entities.Tags
	err := repo.DB.Debug().Where("id = ?", id).Find(&tags, id).Error

	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (repo *repository) GetByName(ctx context.Context, name string) (*entities.Tags, error) {
	var tags *entities.Tags
	err := repo.DB.Debug().Where("name = ?", name).Find(&tags).Error

	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (repo *repository) Create(ctx context.Context, tags *entities.Tags) (*entities.Tags, error) {
	err := repo.DB.Debug().Create(&tags).Error

	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (repo *repository) Update(ctx context.Context, tags *entities.Tags) (*entities.Tags, error) {
	err := repo.DB.Debug().Save(&tags).Error

	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (repo *repository) Delete(ctx context.Context, id uint64) error {
	var tags *entities.Tags
	err := repo.DB.Debug().Where("id = ?", id).Delete(&tags).Error

	if err != nil {
		return err
	}

	return nil
}
