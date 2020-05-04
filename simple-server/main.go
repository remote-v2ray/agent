package main

import (
	"net/http"
	"simple-server/api"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Any("/api", func(c echo.Context) error {
		h := c.Request().Header
		switch h.Get("v2wss-action") {
		case "GetUser":
			return api.GetUser(c)
		case "GetPbConfig":
			return api.GenNodeConfig(c)
		case "PushStats":
			return api.HandleStats(c)
		default:
			return c.String(http.StatusNotFound, "no action for you")
		}
	})
	e.Logger.Fatal(e.Start(":1323"))
}
