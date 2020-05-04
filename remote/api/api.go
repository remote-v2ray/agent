package api

import (
	"net/url"
	"os"

	"github.com/imroc/req"
)

var endpoint string
var authHeader req.Header

func init() {
	if err := Init(); err != nil {
		panic(err)
	}
}

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

	u.RawQuery = ""
	endpoint = u.String()
	authHeader = req.Header{"v2wss-node": node}

	return

}
