package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type source struct {
	ID   string `json:"id"`
	name string `json:"name"`
}

type article struct {
	Source      source `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	PublishedAt string `json:publishedAt`
	Content     string `json:content`
}

type newsResp struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults`
	Articles     []article `json:articles`
}

func main() {
	url := "https://newsapi.org/v2/everything?q=korea"
	cli := http.Client{
		Timeout: time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// set apiKey header
	rawKey, err := ioutil.ReadFile("apiKey")
	if err != nil {
		log.Fatal(err)
	}
	key := strings.TrimSpace(string(rawKey))
	fmt.Println(key)
	req.Header.Set("Authorization", string(key))

	// make request
	res, getErr := cli.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	headlines := newsResp{}
	jsonErr := json.Unmarshal(body, &headlines)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(headlines)
}
