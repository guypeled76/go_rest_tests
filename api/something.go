package api

import "fmt"

// Something is a result that should be returned as json
type Something struct {
	Project string `json:"project"`
	Version string `json:"version"`
}

// GetSomething gets something :-)
func GetSomething(project string, version string) (*Something, error) {

	// Simulate error
	if project == "" {
		return nil, fmt.Errorf("Project should not be empty")
	}

	// Simulate no data state
	if version == "" {
		return nil, nil
	}

	return &Something{
		Project: project,
		Version: version,
	}, nil
}
