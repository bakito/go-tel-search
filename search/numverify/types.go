package numverify

import "github.com/bakito/go-tel-search/search"

type Error struct {
	Success bool `json:"success"`
	Error   struct {
		Code int    `json:"code"`
		Type string `json:"type"`
		Info string `json:"info"`
	} `json:"error"`
}

type Record struct {
	Valid               bool   `json:"valid"`
	Number              string `json:"number"`
	LocalFormat         string `json:"local_format"`
	InternationalFormat string `json:"international_format"`
	CountryPrefix       string `json:"country_prefix"`
	CountryCode         string `json:"country_code"`
	CountryName         string `json:"country_name"`
	Location            string `json:"location"`
	Carrier             string `json:"carrier"`
	LineType            string `json:"line_type"`
}

// AsResult convert into a result
func (r *Record) AsResult() search.Result {
	return search.Result{
		Phone:         r.Number,
		CountryCode:   r.CountryCode,
		CountryPrefix: r.CountryPrefix,
		CountryName:   r.CountryName,
	}
}
