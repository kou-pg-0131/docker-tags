package infrastructures

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"

	"github.com/kou-pg-0131/docker-tags/src/domain"
)

// DockerAPIClient ...
type DockerAPIClient struct {
	httpClient  IHTTPClient
	jsonDecoder IJSONDecoder
}

// NewDockerAPIClient ...
func NewDockerAPIClient() *DockerAPIClient {
	return &DockerAPIClient{
		httpClient:  NewHTTPClient(),
		jsonDecoder: NewJSONDecoder(),
	}
}

// GetTags ...
func (c *DockerAPIClient) GetTags(image string) (*domain.DockerTags, error) {
	u := fmt.Sprintf("https://registry.hub.docker.com/v1/repositories/%s/tags", image)
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		return nil, errors.New(buf.String())
	}

	ts := new(domain.DockerTags)
	if err := c.jsonDecoder.Decode(resp.Body, ts); err != nil {
		return nil, err
	}

	return ts, nil
}
