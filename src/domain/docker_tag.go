package domain

import "sort"

// DockerTag ...
type DockerTag struct {
	Name string `json:"name"`
}

// DockerTags ...
type DockerTags []*DockerTag

// Sort ...
func (ts *DockerTags) Sort() {
	sort.SliceStable(*ts, func(i, j int) bool { return (*ts)[i].Name < (*ts)[j].Name })
}
