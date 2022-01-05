package news

import "errors"

var (
	draft   = "draft"
	deleted = "deleted"
	publish = "publish"
)

type CreateNewNews struct {
	Name   string   `json:"name"`
	Status string   `json:"status"`
	Tags   []string `json:"tags"`
}

func (v CreateNewNews) Validate() error {
	if v.Name == "" {
		return errors.New("name cannot be null")
	}

	if v.Status == "" {
		return errors.New("status cannot be null")
	}

	if v.Status != draft && v.Status != deleted && v.Status != publish {
		return errors.New(`Status can only "draft", "deleted", "publish"`)
	}

	if v.Tags == nil {
		return errors.New("tags cannot be null")
	}
	return nil
}
