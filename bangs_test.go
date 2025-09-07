package main

import (
	"log"
	"testing"
)

func TestInitBang(t *testing.T) {
	bangs := NewBangs()

	err := bangs.Add("!go", "http://golang.org/search?q=%s")
	if err == nil {
		t.Fail()
	}
}

func TestGetSearchUrl(t *testing.T) {
	res := getSearchUrl("Hello, world !yt")

	if res != "https://www.youtube.com/results?search_query=Hello%2C%20world" {
		log.Println(res)
		t.Fail()
	}
}
