package domain

// DockerTag ...
type DockerTag struct {
	Name string `json:"name"`
}

// DockerTags ...
type DockerTags []*DockerTag
