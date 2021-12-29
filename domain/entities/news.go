package entities

import (
	"errors"

	"gorm.io/gorm"
)

var (
	draft   = "draft"
	deleted = "deleted"
	publish = "publish"
)

type News struct {
	gorm.Model
	Name   string  `gorm:"not null;unique" json:"name"`
	Status string  `gorm:"not null" json:"status"`
	Tags   []*Tags `gorm:"many2many:news_tags" json:"tags"`
}

func (news *News) Validate() error {
	if news.Name == "" {
		return errors.New("name cannot be null")
	}

	if news.Status == "" {
		return errors.New("status cannot be null")
	}

	if news.Status != draft && news.Status != deleted && news.Status != publish {
		return errors.New(`Status can only "draft", "deleted", "publish"`)
	}

	return nil
}

func (news *News) Update(tags []*Tags, name, status string) error {
	if news.Name != "" {
		news.Name = name
	}

	if news.Status != "" {
		news.Status = status
	}

	if tags != nil {
		news.Tags = tags
	}

	return nil
}

type NewsDTO struct {
	Id     uint
	Name   string
	Status string
	Tags   []TagsDTO
}

func (news *News) ToDTO() NewsDTO {
	tagsDTO := []TagsDTO{}
	for _, v := range news.Tags {
		tagsDTO = append(tagsDTO, v.ToDTO())
	}

	return NewsDTO{
		Id:     news.Model.ID,
		Name:   news.Name,
		Status: news.Status,
		Tags:   tagsDTO,
	}
}

type NewsDTOTags struct {
	Id     uint
	Name   string
	Status string
}

func (news *News) ToDTOTags() NewsDTOTags {
	return NewsDTOTags{
		Id:     news.Model.ID,
		Name:   news.Name,
		Status: news.Status,
	}
}
