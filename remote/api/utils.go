package api

import "github.com/imroc/req"

func checkRes(res *req.Resp) (err error) {
	resp := res.Response()
	if resp.StatusCode != 200 {
		var msg string
		msg, err = res.ToString()
		if err != nil {
			return
		}
		err = newError("status code:", resp.StatusCode, ". body: ", msg)
		return
	}
	return
}
