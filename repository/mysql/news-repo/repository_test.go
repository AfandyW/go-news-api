package newsrepo

import (
	"context"
	"go-news-api/domain/entities"
	tagsrepo "go-news-api/repository/mysql/tags-repo"
	"go-news-api/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	db, _    = test.NewDB()
	repo     = NewRepository(db)
	tagsRepo = tagsrepo.NewRepository(db)
)

func TestCreateNews(t *testing.T) {
	// Arrange
	ctx := context.Background()
	test.SeedtableTags(tagsRepo)

	var tags = []string{
		"Investment", "Crypto",
	}

	// Action
	tAgs := []*entities.Tags{}

	for _, v := range tags {
		tag, _ := tagsRepo.GetByName(ctx, v)
		tAgs = append(tAgs, tag)
	}

	newsName := "Investasi Sekarang"
	draft := "draft"

	news := entities.News{
		Name:   newsName,
		Status: draft,
		Tags:   tAgs,
	}

	result, err := repo.Create(ctx, &news)

	// Assert
	t.Run("[Create News]should return news property", func(t *testing.T) {

		assert.Nil(t, err)
		assert.Equal(t, tags[0], news.Tags[0].Name)
		assert.Equal(t, tags[1], news.Tags[1].Name)
		assert.Equal(t, newsName, news.Name)
		assert.Equal(t, draft, news.Status)
		assert.Equal(t, result, &news)
	})
}

func TestGetById(t *testing.T) {
	// Arrange
	ctx := context.Background()

	// Action
	var id uint64 = 1
	news, err := repo.GetById(ctx, id)

	// Assert
	t.Run("[Get News By Id]should return id, and without error", func(t *testing.T) {

		assert.Nil(t, err)
		assert.Equal(t, uint(id), news.ID)
	})
}

func TestUpdateTags(t *testing.T) {
	// Arrange
	ctx := context.Background()

	var tags = []string{
		"Business", "Crypto",
	}

	// Action
	tAgs := []*entities.Tags{}

	for _, v := range tags {
		tag, _ := tagsRepo.GetByName(ctx, v)
		tAgs = append(tAgs, tag)
	}

	newsName := "Investasi Sekarang"
	draft := "publish"

	news, _ := repo.GetById(ctx, 1)
	news.Update(tAgs, newsName, draft)

	result, err := repo.Update(ctx, 1, news)

	// Assert
	t.Run("[Update News]should return news property", func(t *testing.T) {

		assert.Nil(t, err)
		assert.Equal(t, tags[0], news.Tags[0].Name)
		assert.Equal(t, tags[1], news.Tags[1].Name)
		assert.Equal(t, newsName, news.Name)
		assert.Equal(t, draft, news.Status)
		assert.Equal(t, result, news)
	})
}

func TestDeleteTags(t *testing.T) {
	// Arrange
	ctx := context.Background()
	tag, _ := repo.GetById(ctx, 1)

	// Action
	err := repo.Delete(ctx, uint64(tag.ID))

	// Assert
	t.Run("[Delete Tags]should without error", func(t *testing.T) {

		assert.Nil(t, err)
	})
}
