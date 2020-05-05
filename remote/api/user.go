package api

import (
	"github.com/imroc/req"
)

type Hash []byte

type V2rayNodeUser struct {
	Email   string
	AlterID uint64
	ID      string
}

func GetUser(utoken string) (user V2rayNodeUser, err error) {

	params := req.Param{
		"user": utoken,
	}
	res, err := req.Get(endpoint, params, req.Header{"v2wss-action": "GetUser"})
	if err != nil {
		return
	}
	if err = checkRes(res); err != nil {
		return
	}

	if err = res.ToJSON(&user); err != nil {
		return
	}

	if user.Email == "" {
		err = newError("not found user")
		return
	}
	return

}
