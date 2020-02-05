package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
	//"github.com/hyeonjungko/marble/utils"
)

type source struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type article struct {
	Source      source `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}

type newsResp struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []article `json:"articles"`
}

func main() {
	url := "https://newsapi.org/v2/everything?q=google"
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
	//fmt.Println(key)
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

	// if news.status != "ok", print err and exit

	// DEV ONLY
	json, err := json.MarshalIndent(headlines, "", "    ")
	fmt.Println(string(json))

	expurl := headlines.Articles[0].URL
	fmt.Println(expurl)
	//exphtml := utils.FetchArticle(expurl)
	//fmt.Print(exphtml)
}
