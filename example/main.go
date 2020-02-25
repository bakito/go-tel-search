package main

import (
	"fmt"

	"github.com/bakito/go-tel-search/search"
	"github.com/bakito/go-tel-search/search/numverify"
	"github.com/bakito/go-tel-search/search/searchch"
)

func main() {
	cl := searchch.NewFor(search.Config{
		Key: "xxx",
		URL: searchch.TestURL,
	})

	res, err := cl.Search("0111111111")
	if err != nil {
		panic(err)
	}

	for _, e := range res {
		fmt.Println(e)
	}

	cl = numverify.NewFor(search.Config{
		Key: "xxx",
	})

	res, err = cl.Search("00111111111")
	if err != nil {
		panic(err)
	}

	for _, e := range res {
		fmt.Println(e)
	}
}
