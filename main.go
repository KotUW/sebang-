package main

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

//go:embed "public/index.html"
var index []byte

func handleSearch(w http.ResponseWriter, req *http.Request) {
	// req.ParseForm()
	// query := req.Form.Get("q")
	query := req.URL.Query().Get("q")

	// fmt.Println("Got query params: ", req.RequestURI, query, req.Form)

	if query == "" {
		io.WriteString(w, string(index))
		return
	}

	query = getSearchUrl(query)
	log.Println("Redirecting to: ", query)

	http.Redirect(w, req, query, http.StatusFound)
}

func main() {
	fmt.Println("Started Server at ::1:8080")

	http.HandleFunc("/search", handleSearch)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	io.WriteString(w, string(index))
	// })

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getSearchUrl(query string) string {
	bangs := strings.Split(query, " ")
	var r_url string

	for i, word := range bangs {
		if strings.HasPrefix(word, "!") {
			switch word {
			case "!wiki":
				r_url = "https://en.wikipedia.org/w/index.php?search=%s"
			case "!cwen": // wikipedia - cite this page.
				r_url = "https://en.wikipedia.org/wiki/Special:CiteThisPage?page=%s"
			case "!yt":
				r_url = "https://www.youtube.com/results?search_query=%s"
			case "!tinyurl":
				r_url = "https://tinyurl.com/create.php?source=indexpage&url=%s&submit=Make+TinyURL%21&alias="
			case "!ddg":
				r_url = "https://duckduckgo.com/?q=%s"
			case "!go":
				r_url = "http://golang.org/search?q=%s"
			case "!pip":
				r_url = "https://pypi.python.org/pypi?:action=search&term=%s&submit=search"
			case "!gh":
				r_url = "https://github.com/search?q=%s&type=repositories"
			}
			bangs[i] = ""
		} else {
			r_url = "https://google.com/search?q=%s"
		}

	}

	res := fmt.Sprintf(r_url, url.PathEscape(strings.Join(bangs, " ")))
	return res
}
