package api

import (
	"github.com/imroc/req"
)

// [][ email, uplink, downlink ]
func PushStats(stats [][]interface{}) (err error) {

	body := req.BodyJSON(map[string]interface{}{
		"stats": stats,
	})
	res, err := req.Post(endpoint, body, authHeader, req.Header{"v2wss-action": "PushStats"})
	if err != nil {
		return
	}
	if err = checkRes(res); err != nil {
		return
	}

	return

}
