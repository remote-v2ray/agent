package manager

import (
	"context"

	"v2ray.com/core"
	"v2ray.com/core/features/inbound"
	feature_stats "v2ray.com/core/features/stats"
	"v2ray.com/core/proxy"
)

var lum *LocalUserManager

func Init(c *core.Instance) (err error) {
	sm := c.GetFeature(feature_stats.ManagerType()).(feature_stats.Manager)
	im := c.GetFeature(inbound.ManagerType()).(inbound.Manager)
	var ih inbound.Handler
	if ih, err = im.GetHandler(context.Background(), "v2wss"); err != nil {
		return
	}
	gi, ok := ih.(proxy.GetInbound)
	if !ok {
		return newError("can't get inbound proxy from handler.")
	}
	p := gi.GetInbound()
	um, ok := p.(proxy.UserManager)
	if !ok {
		return newError("proxy is not a UserManager")
	}

	lum = NewLocalUserManager()
	lum.StatsManager = sm
	lum.MemoryUserManager = um

	return
}

func AddUser(utoken string) (err error) {

	user, err := GetUser(utoken)
	if err != nil {
		return
	}
	if err = lum.AddUser(user); err != nil {
		return
	}

	return
}
