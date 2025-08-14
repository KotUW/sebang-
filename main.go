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
var bangs = Newbangs()

func handleSearch(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query().Get("q")
	// userBangs := false

	// // fmt.Println("Got query params: ", req.RequestURI, query, req.Form)

	// cook, err := req.Cookie("user_bangs")
	// if err == nil {
	// 	userBangs = true
	// 	log.Println("Cookie Value: ", cook.Value)
	// }

	// if userBangs {
	// 	log.Println("Is the bang inside user bangs?")
	// }

	if query == "" { // User enter emty search string
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	query = getSearchUrl(query)
	log.Println("Redirecting to: ", query)

	http.Redirect(w, req, query, http.StatusFound)
}

func main() {
	log.SetPrefix("[Banger Search] ")
	log.Println("Started Server at ::1:8080")

	http.HandleFunc("/search/", handleSearch)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := io.Writer.Write(w, index)
		if err != nil {
			log.Println("[ERROR] Trying to send index page:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println(http.ListenAndServe(":8080", nil))
}

func getSearchUrl(query string) string {
	r, _ := regexp.Compile(`(^| )![a-z]*`)
	match := r.FindString(query)
	if match == "" {
		// Redirect to Bangs.default
		return fmt.Sprintf(bangs.Default, url.PathEscape(query))
	}

	searchUrl := bangs.Query(strings.TrimSpace(strings.ToLower(match)))

	query = strings.Replace(query, match, "", 1)
	res := fmt.Sprintf(searchUrl, url.PathEscape(query))
	return res
}
