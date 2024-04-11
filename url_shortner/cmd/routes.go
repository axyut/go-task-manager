package cmd

import (
	"encoding/json"
	"io"
	"net/http"
	"regexp"
)

var Shorten = func(w http.ResponseWriter, r *http.Request) {
	// Shorten URL
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var urlData UrlData
	jsonerr := json.Unmarshal(body, &urlData)
	if jsonerr != nil {
		http.Error(w, "Cannot Decode JSON Data.", http.StatusNotAcceptable)
	}

	url := urlData.URL
	var queryUrl string
	if url == "" {
		queryUrl = r.URL.Query().Get("url")
		if queryUrl == "" {
			http.Error(w, "Transport URL as query or JSON body.", http.StatusBadRequest)
			return
		}
		url = queryUrl
	}

	// check if the url is a valid url
	pattern := `^(http(s)?://)(www\.)?[\w-]+(\.[\w-]+)+(/[\w- ;,./?%&=]*)?$`
	if matched, _ := regexp.MatchString(pattern, url); !matched {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	// base62 encoding of url to get short url
	shortened := EncodeBase64([]byte(url))
	send := "http://" + r.Host + "/" + shortened
	if _, ok := shortenedURLs[shortened]; ok {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(send))
		return
	}

	shortenedURLs[shortened] = url
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(send))
}

var Redirect = func(w http.ResponseWriter, r *http.Request) {
	// Redirect to original URL
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	shortened := r.URL.Path[1:]
	if shortened == "" || shortened == "shorten" {
		http.Error(w, "Post URL to /shorten", http.StatusBadRequest)
		return
	}
	if url, ok := shortenedURLs[shortened]; ok {
		http.Redirect(w, r, url, http.StatusSeeOther)
		return
	}

}
