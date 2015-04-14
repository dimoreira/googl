package googl

import (
	// "github.com/parnurzeal/gorequest"
)

type Googl struct {
	Key string
}

func NewClient(key string) *Googl {
	return &Googl{Key: key}
}

func (c *Googl) Shorten(url string) string {
	// request := gorequest.New()
	gUrl := url + "?key=" + c.Key
	return gUrl
}

func (c *Googl) Expand(shortUrl string) string {
	return "Expand"
}
