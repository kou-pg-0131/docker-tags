package controllers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewDockerTagsControllerFactory(t *testing.T) {
	f := NewDockerTagsControllerFactory()

	assert.NotNil(t, f)
}

func TestDockerTagsControllerFactory_Create(t *testing.T) {
	f := new(DockerTagsControllerFactory)
	c := f.Create()

	assert.IsType(t, &DockerTagsController{}, c)
	if c, ok := c.(*DockerTagsController); ok {
		assert.NotNil(t, c.dockerAPIClient)
		assert.NotNil(t, c.dockerTagsPresenter)
	} else {
		t.Fatal()
	}
}
