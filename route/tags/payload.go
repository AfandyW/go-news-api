package tags

import "errors"

type CreateNewTags struct {
	Name string `json:"name"`
}

func (v CreateNewTags) Validate() error {
	if v.Name == "" {
		return errors.New("name cannot be null")
	}
	return nil
}
