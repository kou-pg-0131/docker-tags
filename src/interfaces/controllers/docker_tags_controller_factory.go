package controllers

import (
	"github.com/kou-pg-0131/docker-tags/src/infrastructures"
	"github.com/kou-pg-0131/docker-tags/src/interfaces/presenters"
)

// DockerTagsControllerFactory ...
type DockerTagsControllerFactory struct{}

// NewDockerTagsControllerFactory ...
func NewDockerTagsControllerFactory() *DockerTagsControllerFactory {
	return new(DockerTagsControllerFactory)
}

// Create ...
func (f *DockerTagsControllerFactory) Create() IDockerTagsController {
	return NewDockerTagsController(
		presenters.NewDockerTagsPresenter(),
		infrastructures.NewDockerAPIClient(),
	)
}
