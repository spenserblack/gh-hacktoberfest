// Package label provides Go structures for GitHub issue labels.
package label

// Label represents a GitHub issue label.
type Label struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
}
