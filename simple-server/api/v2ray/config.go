package v2ray

import (
	"v2ray.com/core"
	"v2ray.com/core/app/proxyman"
	"v2ray.com/core/common/net"
	"v2ray.com/core/common/serial"
	"v2ray.com/core/proxy/vmess/inbound"
	"v2ray.com/core/transport/internet"
	"v2ray.com/core/transport/internet/websocket"
)

// GenConfig expose v2ray config
func GenConfig() *core.Config {

	wsPath := "/ray"
	wsPort := 3005

	usersInbound := &core.InboundHandlerConfig{
		Tag: "v2wss",
		ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
			PortRange: net.SinglePortRange(net.Port(wsPort)),
			StreamSettings: &internet.StreamConfig{
				Protocol: internet.TransportProtocol_WebSocket,
				TransportSettings: []*internet.TransportConfig{
					{
						Protocol: internet.TransportProtocol_WebSocket,
						Settings: serial.ToTypedMessage(&websocket.Config{
							Path: wsPath,
						}),
					},
				},
			},
		}),
		ProxySettings: serial.ToTypedMessage(&inbound.Config{}),
	}

	config := getV2rayConfig()
	config.Inbound = append(config.Inbound, usersInbound)

	return config

}
