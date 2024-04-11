package main

import (
	"log"
	"net/http"

	"github.com/axyut/url_shortner/cmd"
)

func main() {

	http.HandleFunc("/", cmd.Redirect)
	http.HandleFunc("/shorten", cmd.Shorten)

	log.Default().Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
