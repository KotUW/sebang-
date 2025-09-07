package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

//go:embed "public/index.html"
var index []byte

// Global variables, so they are cached.
var bangs = NewBangs()
var bangRegex = regexp.MustCompile(`(^| )![a-z]*`)

func handleSearch(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query().Get("q")

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
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if _, err := w.Write(index); err != nil {
			log.Printf("[ERROR] Failed to send index page: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	// Configure server with better defaults
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      nil,
	}

	log.Fatal(server.ListenAndServe())
}

func getSearchUrl(query string) string {
	match := bangRegex.FindString(query)
	if match == "" {
		// Redirect to Bangs.default
		return fmt.Sprintf(bangs.Default, url.PathEscape(query))
	}

	searchUrl := bangs.Query(strings.TrimSpace(strings.ToLower(match)))

	query = strings.Replace(query, match, "", 1)
	res := fmt.Sprintf(searchUrl, url.PathEscape(query))
	return res
}
