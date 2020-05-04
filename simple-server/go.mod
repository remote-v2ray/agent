module simple-server

go 1.14

require (
	github.com/golang/protobuf v1.4.0
	github.com/labstack/echo/v4 v4.1.16
	github.com/remote-v2ray/agent v0.0.0-00010101000000-000000000000
	v2ray.com/core v4.22.0+incompatible
)

// v2ray v4.23.1
replace v2ray.com/core => github.com/v2ray/v2ray-core v4.23.1+incompatible

replace github.com/remote-v2ray/agent => ../
