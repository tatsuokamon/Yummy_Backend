package ytbdtc

import "net/http"

type Client struct {
	Key string
	client *http.Client
}

func (c *Client)Search(sp SearchParam) {
}
