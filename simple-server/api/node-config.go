package api

import (
	"net/http"
	"simple-server/api/v2ray"

	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo/v4"
)

func GenNodeConfig(c echo.Context) (err error) {
	config := v2ray.GenConfig()
	pbconfig, err := proto.Marshal(config)
	return c.JSONBlob(http.StatusOK, pbconfig)
}
