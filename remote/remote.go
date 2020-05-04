package remote

import (
	"net/url"

	"github.com/remote-v2ray/agent/remote/api"
	"github.com/remote-v2ray/agent/remote/manager"
	"v2ray.com/core"
)

func init() {
	if err := api.Init(); err != nil {
		panic(err)
	}
}

func Init(c *core.Instance) (err error) {
	return manager.Init(c)
}

func GetNodeConfig() ([]byte, error) {
	return api.GetNodeConfig()
}

func ActiveUser(url *url.URL) (err error) {
	query := url.Query()

	if query["user"] == nil || len(query["user"]) == 0 {
		err = newError("can't user to active")
		return
	}

	user := query["user"][0]

	if err = manager.AddUser(user); err != nil {
		return
	}

	return
}
