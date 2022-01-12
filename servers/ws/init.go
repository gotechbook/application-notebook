package ws

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"github.com/gotechbook/application-notebook/config"
	"github.com/gotechbook/application-notebook/helper"
	"github.com/gotechbook/application-notebook/logger"
)

const (
	defaultAppId = 101
)

var (
	ClientManager = NewManager()
	appIds        = []uint32{defaultAppId, 102, 103}
	ServerIp      string
	ServerPort    string
)

func GetAppIds() []uint32 {
	return appIds
}

func GetServer() *Server {
	return NewServer(ServerIp, ServerPort)
}

func IsLocal(s *Server) (isLocal bool) {
	if s.Ip == ServerIp && s.Port == ServerPort {
		isLocal = true
	}
	return
}

func InAppIds(appId uint32) (inAppId bool) {
	for _, v := range appIds {
		if v == appId {
			inAppId = true
			return
		}
	}
	return
}

func GetDefaultAppId() uint32 {
	return defaultAppId
}

func StartWebSocket() {
	ServerIp = helper.GetServerIp()
	webSocketPort := strconv.FormatInt(int64(config.C.Server.WebSocketPort), 10)
	rpcPort := config.C.Server.RpcPort
	ServerPort = strconv.FormatInt(int64(rpcPort), 10)
	http.HandleFunc("/wss", wsPage)
	// 添加处理程序
	go ClientManager.start()
	logger.Info("WebSocket 启动程序成功", fmt.Sprintf("%s:%s", ServerIp, webSocketPort))
	http.ListenAndServe(":"+webSocketPort, nil)
}

func wsPage(w http.ResponseWriter, req *http.Request) {
	// 升级协议
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		fmt.Println("升级协议", "ua:", r.Header["User-Agent"], "referer:", r.Header["Referer"])
		return true
	}}).Upgrade(w, req, nil)

	if err != nil {
		http.NotFound(w, req)
		return
	}
	fmt.Println("webSocket 建立连接:", conn.RemoteAddr().String())
	currentTime := uint64(time.Now().Unix())
	client := NewClient(conn.RemoteAddr().String(), conn, currentTime)
	go client.read()
	go client.write()

	// 用户连接事件
	ClientManager.Register <- client
}
