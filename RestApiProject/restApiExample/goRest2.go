package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Posts []struct {
	Userid int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	params := url.Values{}
	params.Add("userId", "1")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts?" + params.Encode())
	if err != nil {
		log.Printf("Request Failed: %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// Log the request body
	bodyString := string(body)
	log.Print(bodyString)
	// Unmarshal result
	posts := Posts{}
	err = json.Unmarshal(body, &posts)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return
	}
	log.Printf("The title of the first post is %s", posts[0].Title)
}