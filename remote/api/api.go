package api

import (
	"net/url"
	"os"
)

var endpoint string

func Init() (err error) {
	endpoint = os.Getenv("APIEndpoint")
	if endpoint == "" {
		err = newError("env `APIEndpoint` is required")
		return
	}
	u, err := url.Parse(endpoint)
	if err != nil {
		return
	}

	node := u.Query().Get("node")
	if node == "" {
		err = newError("query param `node` is required")
		return
	}

	return

}
