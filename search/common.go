package search

import (
	"crypto/tls"
	"net/http"
)

// HTTPClient return a new http client
func HTTPClient(config Config) *http.Client {
	tr := &http.Transport{}
	if config.InsecureSkipVerify {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	return &http.Client{
		Transport: tr,
	}
}
