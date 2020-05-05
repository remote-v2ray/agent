package api

import (
	"github.com/imroc/req"
)

func GetNodeConfig() (config []byte, err error) {

	res, err := req.Get(endpoint, req.Header{"v2wss-action": "GetPbConfig"})
	if err != nil {
		return
	}
	if err = checkRes(res); err != nil {
		return
	}

	config, err = res.ToBytes()
	if err != nil {
		return
	}

	return
}
