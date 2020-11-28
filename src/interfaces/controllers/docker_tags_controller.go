package controllers

import (
	"fmt"

	"github.com/kou-pg-0131/docker-tags/src/interfaces/gateways"
	"github.com/kou-pg-0131/docker-tags/src/interfaces/presenters"
)

// IDockerTagsController ...
type IDockerTagsController interface {
	ShowAll(img string) error
}

// DockerTagsController ...
type DockerTagsController struct {
	dockerAPIClient     gateways.IDockerAPIClient
	dockerTagsPresenter presenters.IDockerTagsPresenter
}

// NewDockerTagsController ...
func NewDockerTagsController(p presenters.IDockerTagsPresenter, c gateways.IDockerAPIClient) *DockerTagsController {
	return &DockerTagsController{
		dockerTagsPresenter: p,
		dockerAPIClient:     c,
	}
}

// ShowAll ...
func (c *DockerTagsController) ShowAll(img string) error {
	ts, err := c.dockerAPIClient.GetTags(img)
	if err != nil {
		return err
	}

	fmt.Println(c.dockerTagsPresenter.ShowAll(img, ts))
	return nil
}
