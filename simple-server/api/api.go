package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/remote-v2ray/agent/remote/api"
)

func GetUser(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, api.V2rayNodeUser{
		Email:   "t@tt.t",
		AlterID: 0,
		ID:      "16167dc8-16b6-4e6d-b8bb-65dd68113a81",
	})
}

func HandleStats(c echo.Context) (err error) {
	return c.String(http.StatusOK, "I have received stats")
}
