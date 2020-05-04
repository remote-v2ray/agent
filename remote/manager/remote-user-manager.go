package manager

import (
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/remote-v2ray/agent/remote/api"
)

// remote user cache
var ruc = func() (c *cache.Cache) {
	c = cache.New(5*time.Minute, time.Minute)
	c.OnEvicted(func(_ string, _user interface{}) {
		user := _user.(api.V2rayNodeUser)
		newError("Delete Cache User: ", user.Email).AtDebug().WriteToLog()
	})
	return
}()

func GetUser(utoken string) (user api.V2rayNodeUser, err error) {

	if user, found := ruc.Get(utoken); found {
		return user.(api.V2rayNodeUser), nil
	}

	if user, err = api.GetUser(utoken); err != nil {
		return
	}

	if user.Email == "" {
		err = newError("not found user")
		return
	}

	newError("Set Cache User: ", user.Email).AtDebug().WriteToLog()
	ruc.Set(utoken, user, cache.DefaultExpiration)

	return
}
