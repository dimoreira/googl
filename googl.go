package googl

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

type Googl struct {
	Key string
}

type ShortMsg struct {
	Kind    string `json:"kind"`
	Id      string `json:"id"`
	LongUrl string `json:"longUrl"`
}

type LongMsg struct {
	Kind    string `json:"kind"`
	Id      string `json:"id"`
	LongUrl string `json:"longUrl"`
	Status  string `json:"status"`
}

func NewClient(key string) *Googl {
	return &Googl{Key: key}
}

func (c *Googl) Shorten(url string) string {
	request := gorequest.New()
	var response string

	gUrl := "https://www.googleapis.com/urlshortener/v1/url?key=" + c.Key
	if c.Key == "" {
		response = "You need to set the Google Url Shortener API Key"
	} else if url == "" {
		response = "You need to set the url to be shortened"
	} else {
		resp, body, _ := request.Post(gUrl).
			Set("Accept", "application/json").
			Set("Content-Type", "application/json").
			Send(`{"longUrl":"` + url + `"}`).End()
		if resp.Status == "200 OK" {
			fmt.Println(body)
			response = "Done! Ok!"
		} else {
			response = "Some error occurred, please try again later"
		}
	}

	return response
}

func (c *Googl) Expand(shortUrl string) string {
	request := gorequest.New()
	var response string

	gUrl := "https://www.googleapis.com/urlshortener/v1/url?key=" + c.Key + "&shortUrl=" + shortUrl
	if c.Key == "" {
		response = "You need to set the Google Url Shortener API Key"
	} else if shortUrl == "" {
		response = "You need to set the url to be expanded"
	} else {
		resp, body, _ := request.Get(gUrl).
			Set("Accept", "application/json").
			Set("Content-Type", "application/json").End()
		if resp.Status == "200 OK" {
			fmt.Println(body)
			response = "Done! Ok!"
		} else {
			response = "Some error occurred, please try again later"
		}
	}

	return response
}
