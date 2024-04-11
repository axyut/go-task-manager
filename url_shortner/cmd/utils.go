package cmd

import (
	"encoding/base64"
	"log"
)

var shortenedURLs = make(map[string]string)

type UrlData struct {
	URL string `json:"url"`
}

func EncodeBase64(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

func DecodeBase64(input string) string {

	data, err := base64.StdEncoding.DecodeString(input)

	if err != nil {
		log.Default().Println("Error decoding base64")
	}
	return (string(data))
}
