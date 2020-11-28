package infrastructures

import (
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/kou-pg-0131/docker-tags/src/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDockerAPIClient_GetTags_ReturnTagsWhenImageFound(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://registry.hub.docker.com/v1/repositories/IMAGE/tags", nil)
	mhc := new(mockHTTPClient)
	mhc.On("Do", req).Return(&HTTPResponse{Body: strings.NewReader("BODY"), StatusCode: 200}, nil)

	mjd := new(mockJSONDecoder)
	mjd.On("Decode", strings.NewReader("BODY"), mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		ts := args.Get(1).(*domain.DockerTags)
		*ts = domain.DockerTags{{Name: "TAG"}}
	})

	c := &DockerAPIClient{httpClient: mhc, jsonDecoder: mjd}
	ts, err := c.GetTags("IMAGE")

	assert.Nil(t, err)
	assert.Equal(t, &domain.DockerTags{{Name: "TAG"}}, ts)
	mhc.AssertNumberOfCalls(t, "Do", 1)
	mjd.AssertNumberOfCalls(t, "Decode", 1)
}

func TestDockerAPIClient_GetTags_ReturnErrorWhenHTTPRequestFailed(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://registry.hub.docker.com/v1/repositories/IMAGE/tags", nil)
	mhc := new(mockHTTPClient)
	mhc.On("Do", req).Return((*HTTPResponse)(nil), errors.New("SOMETHING_WRONG"))

	c := &DockerAPIClient{httpClient: mhc}
	ts, err := c.GetTags("IMAGE")

	assert.EqualError(t, err, "SOMETHING_WRONG")
	assert.Nil(t, ts)
	mhc.AssertNumberOfCalls(t, "Do", 1)
}

func TestDockerAPIClient_GetTags_ReturnErrorWhenHTTPResponseNotSuccess(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://registry.hub.docker.com/v1/repositories/IMAGE/tags", nil)
	mhc := new(mockHTTPClient)
	mhc.On("Do", req).Return(&HTTPResponse{Body: strings.NewReader("BODY"), StatusCode: 500}, nil)

	c := &DockerAPIClient{httpClient: mhc}
	ts, err := c.GetTags("IMAGE")

	assert.EqualError(t, err, "BODY")
	assert.Nil(t, ts)
	mhc.AssertNumberOfCalls(t, "Do", 1)
}

func TestDockerAPIClient_GetTags_ReturnErrorWhenJSONDecodeFailed(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://registry.hub.docker.com/v1/repositories/IMAGE/tags", nil)
	mhc := new(mockHTTPClient)
	mhc.On("Do", req).Return(&HTTPResponse{Body: strings.NewReader("BODY"), StatusCode: 200}, nil)

	mjd := new(mockJSONDecoder)
	mjd.On("Decode", strings.NewReader("BODY"), mock.Anything).Return(errors.New("SOMETHING_WRONG"))

	c := &DockerAPIClient{httpClient: mhc, jsonDecoder: mjd}
	ts, err := c.GetTags("IMAGE")

	assert.EqualError(t, err, "SOMETHING_WRONG")
	assert.Nil(t, ts)
	mhc.AssertNumberOfCalls(t, "Do", 1)
	mjd.AssertNumberOfCalls(t, "Decode", 1)
}
