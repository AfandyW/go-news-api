package tags

import (
	"context"
	"go-news-api/domain/entities"
	"go-news-api/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	db, _ = test.NewDB()
	repo  = NewRepository(db)
)

func TestCreateTags(t *testing.T) {
	// Arrange
	ctx := context.Background()

	// Action
	tag := entities.Tags{
		Name: "Investment",
	}

	result, err := repo.Create(ctx, &tag)

	// Assert
	t.Run("[Create Tags]should return tags name, and without error", func(t *testing.T) {

		assert.Nil(t, err)
		assert.Equal(t, tag.Name, result.Name)
	})
}

func TestGetByIdAndName(t *testing.T) {
	// Arrange
	ctx := context.Background()

	// Action
	var id uint64 = 1
	tag, err := repo.GetById(ctx, id)

	// Assert
	t.Run("[Get Tags By Id]should return id, and without error", func(t *testing.T) {

		assert.Nil(t, err)
		assert.Equal(t, uint(id), tag.ID)
	})

	//Action
	res, err := repo.GetByName(ctx, tag.Name)

	// Assert
	t.Run("[Get Tags By Name]should return name, and without error", func(t *testing.T) {

		assert.Nil(t, err)
		assert.Equal(t, tag.Name, res.Name)
	})
}

func TestUpdateTags(t *testing.T) {
	// Arrange
	ctx := context.Background()
	tag, _ := repo.GetById(ctx, 1)

	// Action
	tag.Update("Crypto")
	result, err := repo.Update(ctx, tag)

	// Assert
	t.Run("[Update Tags]should return tags name, and without error", func(t *testing.T) {

		assert.Nil(t, err)
		assert.Equal(t, tag.Name, result.Name)
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
