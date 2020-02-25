package search

// Client search client interface
type Client interface {
	Search(query ...string) ([]Result, error)
}
