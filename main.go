package main

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

//go:embed "public/index.html"
var index []byte

func handleSearch(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	query := req.Form.Get("q")

	// fmt.Println("Got query params: ", req.RequestURI, query, req.Form)

	if query == "" {
		http.Error(w, "Empty Query", http.StatusBadRequest)
		return
	}

	http.Redirect(w, req, getSearchUrl(query), http.StatusFound)
}

func main() {
	fmt.Println("Started Server at ::1:8080")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, string(index))
	})

	http.HandleFunc("/search", handleSearch)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getSearchUrl(query string) string {
	res := fmt.Sprintf("https://google.com/search?q=%s", url.PathEscape(query))
	return res
}
