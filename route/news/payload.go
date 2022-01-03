package news

type (
	//handle create
	CreateNewNews struct {
		Name   string   `json:"name"`
		Status string   `json:"status"`
		Tags   []string `json:"tags"`
	}
)
