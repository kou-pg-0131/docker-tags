package gateways

import "github.com/kou-pg-0131/docker-tags/src/domain"

// IDockerAPIClient ...
type IDockerAPIClient interface {
	GetTags(img string) (*domain.DockerTags, error)
}
