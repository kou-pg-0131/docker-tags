package controllers

import (
	"errors"
	"testing"

	"github.com/kou-pg-0131/docker-tags/src/domain"
	"github.com/stretchr/testify/assert"
)

func TestDockerTagsController_ShowAll_ReturnNilWhenSuccess(t *testing.T) {
	mac := new(mockDockerAPIClient)
	mac.On("GetTags", "IMAGE").Return(&domain.DockerTags{{Name: "TAG"}}, nil)

	mtp := new(mockDockerTagsPresenter)
	mtp.On("ShowAll", "IMAGE", &domain.DockerTags{{Name: "TAG"}}).Return("OUTPUT")

	c := &DockerTagsController{dockerAPIClient: mac, dockerTagsPresenter: mtp}
	err := c.ShowAll("IMAGE")

	assert.Nil(t, err)
	mac.AssertNumberOfCalls(t, "GetTags", 1)
	mtp.AssertNumberOfCalls(t, "ShowAll", 1)
}

func TestDockerTagsController_ShowAll_ReturnErrorWhenGetTagsFailed(t *testing.T) {
	mac := new(mockDockerAPIClient)
	mac.On("GetTags", "IMAGE").Return((*domain.DockerTags)(nil), errors.New("SOMETHING_WRONG"))

	c := &DockerTagsController{dockerAPIClient: mac}
	err := c.ShowAll("IMAGE")

	assert.EqualError(t, err, "SOMETHING_WRONG")
	mac.AssertNumberOfCalls(t, "GetTags", 1)
}
