package searchch_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/bakito/go-tel-search/search"
	"github.com/bakito/go-tel-search/search/searchch"

	. "gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func Test_Search_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		content, err := os.ReadFile("../../testdata/searchch/api-response.xml")
		Assert(t, is.Nil(err))
		_, err = w.Write(content)
		Assert(t, is.Nil(err))
	}))
	defer ts.Close()

	cl := searchch.NewFor(search.Config{
		Key: "xxx",
		URL: ts.URL,
	})

	res, err := cl.Search("0111111111")
	Assert(t, is.Nil(err))
	Assert(t, is.Len(res, 2))

	Assert(t, is.Equal("Meier", res[0].Name))
	Assert(t, is.Equal("John", res[0].Firstname))

	Assert(t, is.Equal("John Meier IT Consulting", res[1].Name))
	Assert(t, is.Equal("", res[1].Firstname))
}

func Test_Search_Error(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		content, err := os.ReadFile("../../testdata/searchch/api-error.xml")
		Assert(t, is.Nil(err))
		w.WriteHeader(http.StatusForbidden)
		_, err = w.Write(content)
		Assert(t, is.Nil(err))
	}))
	defer ts.Close()

	cl := searchch.NewFor(search.Config{
		Key: "xxx",
		URL: ts.URL,
	})

	res, err := cl.Search("0111111111")
	Assert(t, err != nil)

	Assert(t, is.Len(res, 0))
}
