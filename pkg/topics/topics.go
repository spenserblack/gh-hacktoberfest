// Package topics provides a simple interface for creating and managing topics.
package topics

// Hacktoberfest is the topic for hacktoberfest
const Hacktoberfest = "hacktoberfest"

// Topics is a list of topic names for the REST API.
type Topics struct {
	Names []string `json:"names"`
}

// Set is a set of topic names.
type Set map[string]struct{}

// Set returns a Set from a list of topic names.
func (t *Topics) Set() Set {
	set := make(Set)
	for _, name := range t.Names {
		set[name] = struct{}{}
	}
	return set
}

// Add adds a topic to the list.
func (set Set) Add(name string) {
	set[name] = struct{}{}
}

// Topics converts the set to a list.
func (set Set) Topics() Topics {
	names := make([]string, 0, len(set))
	for name := range set {
		names = append(names, name)
	}
	return Topics{Names: names}
}
