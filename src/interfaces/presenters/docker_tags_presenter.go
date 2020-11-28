package presenters

import (
	"fmt"
	"strings"

	"github.com/kou-pg-0131/docker-tags/src/domain"
)

// IDockerTagsPresenter ...
type IDockerTagsPresenter interface {
	ShowAll(img string, ts *domain.DockerTags) string
}

// DockerTagsPresenter ...
type DockerTagsPresenter struct{}

// NewDockerTagsPresenter ...
func NewDockerTagsPresenter() *DockerTagsPresenter {
	return new(DockerTagsPresenter)
}

// ShowAll ...
func (p *DockerTagsPresenter) ShowAll(img string, ts *domain.DockerTags) string {
	ts.Sort()

	rows := []string{}

	for _, t := range *ts {
		rows = append(rows, fmt.Sprintf("%s:%s", img, t.Name))
	}

	return strings.Join(rows, "\n")
}
