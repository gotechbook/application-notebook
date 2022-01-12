package routers

import (
	"github.com/gotechbook/application-notebook/api/websocket"
	"github.com/gotechbook/application-notebook/servers/ws"
)

func WebsocketInit() {
	ws.Register("ping", websocket.PingController)
	ws.Register("login", websocket.LoginController)
}
