package search

// Result tel search result
type Result struct {
	Name          string      `json:"name"`
	Firstname     string      `json:"firstname"`
	Street        string      `json:"street"`
	StreetNo      string      `json:"streetno"`
	Zip           string      `json:"zip"`
	City          string      `json:"city"`
	State         string      `json:"canton"`
	Phone         string      `json:"phone"`
	Plain         interface{} `json:"plain"`
	CountryPrefix string      `json:"country_prefix"`
	CountryCode   string      `json:"country_code"`
	CountryName   string      `json:"country_name"`
}

// Config client config
type Config struct {
	// Key the search.ch API key
	Key string
	// URL the service URL
	URL string
	// InsecureSkipVerify ignore tls
	InsecureSkipVerify bool
}
