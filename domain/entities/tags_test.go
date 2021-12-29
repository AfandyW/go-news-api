package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name string
	want string
}

func TestNewTags(t *testing.T) {
	tCases := testCase{
		name: "",
		want: "name cannot be null",
	}
	tTags := Tags{
		Name: tCases.name,
	}

	t.Run(tCases.want, func(t *testing.T) {
		err := tTags.Validate()
		assert.Equal(t, tCases.want, err.Error())
	})

	t.Run("Should validate success, with no error", func(t *testing.T) {
		name := "Investment"
		tTags := Tags{
			Name: name,
		}

		err := tTags.Validate()
		assert.Equal(t, nil, err)
	})
}
