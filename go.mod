module github.com/remote-v2ray/agent

go 1.13

require (
	github.com/golang/protobuf v1.4.0
	github.com/gorilla/websocket v1.4.2
	github.com/imroc/req v0.3.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	golang.org/x/sys v0.0.0-20200501145240-bc7a7d42d5c3 // indirect
	v2ray.com/core v4.22.0+incompatible
)

// remote v2ray v4.22.1
replace v2ray.com/core => github.com/remote-v2ray/v2ray-core v1.24.5-0.20200203055314-86d5bd866b3d
