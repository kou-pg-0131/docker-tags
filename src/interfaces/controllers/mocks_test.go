package controllers

import (
	"github.com/kou-pg-0131/docker-tags/src/domain"
	"github.com/stretchr/testify/mock"
)

type mockDockerAPIClient struct {
	mock.Mock
}

func (m *mockDockerAPIClient) GetTags(img string) (*domain.DockerTags, error) {
	args := m.Called(img)
	return args.Get(0).(*domain.DockerTags), args.Error(1)
}

type mockDockerTagsPresenter struct {
	mock.Mock
}

func (m *mockDockerTagsPresenter) ShowAll(img string, ts *domain.DockerTags) string {
	args := m.Called(img, ts)
	return args.String(0)
}
