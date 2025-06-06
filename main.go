package main

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

//go:embed "public/index.html"
var index []byte
var bangs = init_bangs()

func handleSearch(w http.ResponseWriter, req *http.Request) {
	// req.ParseForm()
	// query := req.Form.Get("q")
	query := req.URL.Query().Get("q")

	// fmt.Println("Got query params: ", req.RequestURI, query, req.Form)

	if query == "" {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	query = getSearchUrl(query)
	log.Println("Redirecting to: ", query)

	http.Redirect(w, req, query, http.StatusFound)
}

func main() {
	log.Println("Started Server at ::1:8080")

	http.HandleFunc("/search/", handleSearch)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, string(index))
	})

	log.Println(http.ListenAndServe(":8080", nil))
}

func getSearchUrl(query string) string {
	geturl := func(query string) (string, string) {
		r, _ := regexp.Compile(` ![a-z]*`)
		match := r.FindString(query)
		if match == "" {
			return "https://google.com/search?q=%s", query
		}

		r_url := bangs[strings.TrimSpace(strings.ToLower(match))]
		query = strings.ReplaceAll(query, match, "")
		if r_url == "" {
			return "https://google.com/search?q=%s", query
		}
		return r_url, query
	}
	search_url, query := geturl(query)
	res := fmt.Sprintf(search_url, url.PathEscape(query))
	return res
}
