package presenters

import (
	"testing"

	"github.com/kou-pg-0131/docker-tags/src/domain"
	"github.com/stretchr/testify/assert"
)

func TestDockerTagsPresenter_Show(t *testing.T) {
	p := new(DockerTagsPresenter)

	s := p.Show("IMAGE", &domain.DockerTags{
		{Name: "3"},
		{Name: "1"},
		{Name: "4"},
		{Name: "2"},
		{Name: "5"},
	})

	assert.Equal(t, `IMAGE:1
IMAGE:2
IMAGE:3
IMAGE:4
IMAGE:5`, s)
}
