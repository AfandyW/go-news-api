package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCaseN struct {
	name   string
	status string
	want   string
}

func TestNewNews(t *testing.T) {
	tCases := []testCaseN{
		{
			name:   "",
			status: "draft",
			want:   "name cannot be null",
		},
		{
			name:   "Investment",
			status: "",
			want:   "status cannot be null",
		},
		{
			name:   "Investment",
			status: "draftss",
			want:   `Status can only "draft", "deleted", "publish"`,
		},
	}

	for _, tc := range tCases {
		tNews := News{
			Name:   tc.name,
			Status: tc.status,
		}
		t.Run(tc.want, func(t *testing.T) {
			err := tNews.Validate()
			assert.Equal(t, tc.want, err.Error())
		})
	}
}
