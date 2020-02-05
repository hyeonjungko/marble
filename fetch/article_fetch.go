package fetch

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Fetches the text of article given url
func fetchArticle(url string) string {
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
