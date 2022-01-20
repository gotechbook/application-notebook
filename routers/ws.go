package routers

import (
	"github.com/gotechbook/application-notebook/api/websocket/log"
	"github.com/gotechbook/application-notebook/api/websocket/pixel"
	"github.com/gotechbook/application-notebook/api/websocket/system"
	"github.com/gotechbook/application-notebook/servers/ws"
)

func WebsocketInit() {
	// system
	ws.Register("ping", system.PingController)
	ws.Register("login", system.LoginController)
	ws.Register("heartbeat", system.HeartbeatController)

	// pixel
	ws.Register("pixel.extract", pixel.ExtractController)
	ws.Register("pixel.synthesis", pixel.SynthesisController)
	ws.Register("pixel.drag", pixel.DragController)

	// log
	ws.Register("log.block", log.LoggerBlockController)
}
