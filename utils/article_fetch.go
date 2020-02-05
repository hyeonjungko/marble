package utils

import (
	"io/ioutil"
	"log"
	"net/http"
)

// FetchArticle fetches the text of article given url
func FetchArticle(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(html)
}
