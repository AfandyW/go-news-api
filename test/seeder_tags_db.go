package test

import (
	"context"
	"go-news-api/domain/entities"
	"go-news-api/domain/tags"
)

func SeedtableTags(repo tags.IRepository) {
	ctx := context.Background()
	var tags = []string{
		"Investment", "Crypto", "Business",
	}
	for _, v := range tags {
		tag := entities.Tags{
			Name: v,
		}
		repo.Create(ctx, &tag)
	}
}
