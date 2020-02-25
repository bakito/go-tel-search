package searchch

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/bakito/go-tel-search/search"
	"golang.org/x/text/encoding/charmap"
)

const (
	// DefaultURL the default api service url
	DefaultURL = "https://tel.search.ch/api/"
	// TestURL the test api service url
	TestURL = "https://tel.search.ch/examples/api-response.xml"
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

// Search implement search
func (c *client) Search(query ...string) ([]search.Result, error) {

	was := strings.Join(query, "+")

	res := &Feed{}

	url := fmt.Sprintf("%s?key=%s&was=%s", c.url, c.key, was)

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

	if err == nil && resp.StatusCode != 200 {
		err = fmt.Errorf("status code %v", resp.StatusCode)
	}

	if err != nil {
		return nil, err
	}

	if err = parseXML(resp.Body, res); err != nil {
		return nil, err
	}

	if res.ErrorCode != 0 {
		return nil, fmt.Errorf("code: %d, message: %s, reason: %s", res.ErrorCode, res.ErrorMessage, res.ErrorReason)
	}

	var results []search.Result

	for _, e := range res.Entry {
		results = append(results, e.AsResult())
	}

	return results, err
}

func parseXML(xmlDoc io.Reader, target interface{}) error {
	decoder := xml.NewDecoder(xmlDoc)
	decoder.CharsetReader = makeCharsetReader
	return decoder.Decode(target)
}

func makeCharsetReader(charset string, input io.Reader) (io.Reader, error) {
	if charset == "ISO-8859-1" {
		// Windows-1252 is a superset of ISO-8859-1, so should do here
		return charmap.Windows1252.NewDecoder().Reader(input), nil
	}
	return nil, fmt.Errorf("Unknown charset: %s", charset)
}
