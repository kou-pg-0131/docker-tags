package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDockerTag_Sort(t *testing.T) {
	ts := &DockerTags{
		{Name: "3"},
		{Name: "1"},
		{Name: "4"},
		{Name: "2"},
		{Name: "5"},
	}

	ts.Sort()
	assert.Equal(t, &DockerTags{
		{Name: "1"},
		{Name: "2"},
		{Name: "3"},
		{Name: "4"},
		{Name: "5"},
	}, ts)
}
