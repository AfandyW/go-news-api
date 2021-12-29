package entities

import (
	"errors"

	"gorm.io/gorm"
)

type Tags struct {
	gorm.Model
	Name string  `gorm:"not null;unique" json:"name"`
	News []*News `gorm:"many2many:news_tags" json:"news"`
}

func (tags *Tags) Validate() error {
	if tags.Name == "" {
		return errors.New("name cannot be null")
	}

	return nil
}

func (tags *Tags) Update(name string) error {
	if name != "" {
		tags.Name = name
	}

	return nil
}

type TagsDTO struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func (t *Tags) ToDTO() TagsDTO {
	return TagsDTO{
		Id:   t.Model.ID,
		Name: t.Name,
	}
}

type TagsDTONews struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	News []NewsDTOTags
}

func (t *Tags) ToDTOWithNews() TagsDTONews {
	NewsDTO := []NewsDTOTags{}
	for _, v := range t.News {
		NewsDTO = append(NewsDTO, v.ToDTOTags())
	}

	return TagsDTONews{
		Id:   t.Model.ID,
		Name: t.Name,
		News: NewsDTO,
	}
}
