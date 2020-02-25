package numverify

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/bakito/go-tel-search/search"
)

const (
	// DefaultURL the default api service url
	DefaultURL = "http://apilayer.net/api/validate/"
)

// New create a new client with the given key
func New(key string) search.Client {
	return NewFor(search.Config{
		Key: key,
	})
}

// NewFor create a new client for the given config
func NewFor(config search.Config) search.Client {
	url := config.URL
	if url == "" {
		url = DefaultURL
	}

	tr := &http.Transport{}
	if config.InsecureSkipVerify {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	return &client{
		key: config.Key,
		url: url,
		hc:  search.HTTPClient(config),
	}
}

// client implements Client
var _ search.Client = &client{}

type client struct {
	key string
	url string
	hc  *http.Client
}

func (c *client) Search(query ...string) ([]search.Result, error) {

	if len(query) != 1 {
		return nil, fmt.Errorf("only one phone number supported as query")
	}
	// it can be safely placed inside a URL query
	safePhone := url.QueryEscape(query[0])

	url := fmt.Sprintf("%s?access_key=%s&number=%s", c.url, c.key, safePhone)

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var e Error

		if err := json.NewDecoder(resp.Body).Decode(&e); err != nil {
			log.Println(err)
		}
		return nil, fmt.Errorf("code: %d, type: %s, info: %s", e.Error.Code, e.Error.Type, e.Error.Info)
	}

	var record Record

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	return []search.Result{record.AsResult()}, nil
}
